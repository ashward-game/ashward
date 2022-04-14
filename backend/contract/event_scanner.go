package contract

import (
	"context"
	"errors"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"math/rand"
	"orbit_nft/logger"
	"strings"
	"time"
)

const (
	DefaultEventWorkerTimeout   = 15 * time.Second
	DefaultPollBatchSize        = 5000
	MinPollBatchSize            = 100
	DefaultWaitNewBlockDuration = 5 * time.Second
	MaxFetchLogTryTime          = 3

	HandlerWaitDuration    = 10 * time.Second
	EventLogChanBufferSize = 500
)

type RichEventLog struct {
	Network    string
	ScanTime   time.Time
	StartBlock uint64
	EndBlock   uint64
	Logs       []types.Log
}

type eventScanner struct {
	ethcs     []*ethclient.Client
	ethcNames []string

	defaultBatch int64
	autoTuning   bool
	minBatchSize int64
}

func newEventScanner(ethc map[string]*ethclient.Client, defaultBatch int64, autoTuning bool) *eventScanner {
	var ethcs []*ethclient.Client
	var ethcName []string
	for k, v := range ethc {
		ethcName = append(ethcName, k)
		ethcs = append(ethcs, v)
	}
	return &eventScanner{
		ethcs:        ethcs,
		ethcNames:    ethcName,
		defaultBatch: defaultBatch,
		autoTuning:   autoTuning,
		minBatchSize: MinPollBatchSize,
	}
}

func (c *eventScanner) getClient() (*ethclient.Client, string) {
	var idx = rand.Intn(len(c.ethcs))
	return c.ethcs[idx], c.ethcNames[idx]
}

func (c *eventScanner) getLatestBlock(ctx context.Context) (*big.Int, error) {
	ethc, _ := c.getClient()
	header, err := ethc.HeaderByNumber(ctx, nil)
	if err != nil {
		return nil, err
	}

	return new(big.Int).Set(header.Number), nil
}

// Poll crawl event logs from fromBlock (use current block if not specified),
// until reach last block or wait if toBlock is not specified
func (c *eventScanner) Poll(ctx context.Context, addresses []common.Address, fromBlock, toBlock *big.Int, sub chan *RichEventLog) {
	var curFromBlock *big.Int
	var curToBlock *big.Int
	var tmpLatestBlock *big.Int // used in case of fetching and wait
	var curBatchSize = c.defaultBatch
	var err error

	if fromBlock == nil {
		// assign from block as current block
		fromBlock, err = c.getLatestBlock(ctx)
		if err != nil {
			logger.Errorw("Cannot get current block, stop crawl right now!", "error", err)
			panic(errors.New("cannot get current block"))
		}
	}
	curFromBlock = new(big.Int).Set(fromBlock)

	if toBlock == nil {
		tmpLatestBlock, err = c.getLatestBlock(ctx)
		if err != nil {
			logger.Errorw("Cannot get current block, stop crawl right now!", "error", err)
			panic(errors.New("cannot get current block"))
		}
	}
	curToBlock = new(big.Int)

	var retryCounter = 0
	for {
		select {
		case <-ctx.Done():
			logger.Infow("Stop crawl event log", "contracts", addresses)
			return
		default:
			if toBlock != nil {
				if curFromBlock.Cmp(toBlock) >= 0 {
					logger.Infow("Finish poll event log",
						"addresses", addresses,
						"fromBlock", fromBlock.String(),
						"toBlock", toBlock.String())
					return
				}

				// calc next toBlock
				if curToBlock.Add(curFromBlock, big.NewInt(curBatchSize)).Cmp(toBlock) >= 0 {
					curToBlock = new(big.Int).Set(toBlock)
				}
			} else {
				if curToBlock.Add(curFromBlock, big.NewInt(curBatchSize)).Cmp(tmpLatestBlock) >= 0 {
					// renew tmpLatestBlock
					newLatestBlock, err := c.getLatestBlock(ctx)
					if err != nil {
						logger.Errorw("Cannot get current block, stop crawl right now!", "error", err)
						return
					}
					if tmpLatestBlock.Cmp(newLatestBlock) >= 0 {
						// wait for new block
						logger.Infof("Sleep a while to wait new block")
						time.Sleep(DefaultWaitNewBlockDuration)
						continue
					}

					tmpLatestBlock = newLatestBlock
					curToBlock = curToBlock.Set(tmpLatestBlock)
				}
			}

			// start fetching logs
			evtLog, err := c.tryCrawl(ctx, ethereum.FilterQuery{
				FromBlock: curFromBlock,
				ToBlock:   curToBlock,
				Addresses: addresses,
			})
			if err != nil {
				// handle for each kind of error (network, timeout, exceed rate limit, ...)
				logger.Errorw("Cannot crawl event log, retry again", "error", err)

				if strings.Contains(err.Error(), "too many requests") {
					time.Sleep(DefaultWaitNewBlockDuration / 2)
				} else if strings.Contains(err.Error(), "exceed maximum block range") {
					if c.autoTuning {
						curBatchSize = curBatchSize / 2
					}
				} else {
					retryCounter++
					// tuning batch size if failure too much
					if retryCounter >= MaxFetchLogTryTime && c.autoTuning {
						// simple policy, reduce by divide by 2
						if curBatchSize/2 >= c.minBatchSize {
							curBatchSize = curBatchSize / 2
						}
						retryCounter = 0
					}
				}

				continue
			}

			// reset counter
			sub <- evtLog
			retryCounter = 0
			curFromBlock.Add(curToBlock, big.NewInt(1))
		}
	}
}

func (c *eventScanner) tryCrawl(ctx context.Context, query ethereum.FilterQuery) (*RichEventLog, error) {
	start := time.Now()
	var logs []types.Log
	var err error
	defer func() {
		logger.Debugw("Fetch logs",
			"startBlock", query.FromBlock.Int64(),
			"toBlock", query.ToBlock.Int64(),
			"processTime", time.Since(start),
			"logSize", len(logs))
	}()
	var ethc, name = c.getClient()
	logs, err = ethc.FilterLogs(ctx, query)
	if err != nil {
		return nil, err
	}

	var startBlock, endBlock uint64 = 0, 0
	for _, l := range logs {
		if startBlock == 0 {
			startBlock = l.BlockNumber
		} else if startBlock > l.BlockNumber {
			startBlock = l.BlockNumber
		}
		if endBlock == 0 {
			endBlock = l.BlockNumber
		} else if endBlock < l.BlockNumber {
			endBlock = l.BlockNumber
		}
	}
	if startBlock == 0 {
		startBlock = query.FromBlock.Uint64()
	}
	if endBlock == 0 {
		endBlock = query.ToBlock.Uint64()
	}

	return &RichEventLog{
		Logs:       logs,
		StartBlock: startBlock,
		EndBlock:   endBlock,
		ScanTime:   time.Now(),
		Network:    name,
	}, nil
}
