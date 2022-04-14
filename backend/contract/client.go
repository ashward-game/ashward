package contract

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	"go.uber.org/zap"
	"math/big"
	orbitContext "orbit_nft/contract/context"
	"orbit_nft/contract/event"
	"orbit_nft/contract/event/whitelist"
	"orbit_nft/contract/rpc"
	"orbit_nft/db/model"
	"orbit_nft/db/repository"
	"orbit_nft/db/service"
	"orbit_nft/logger"
	"orbit_nft/util"
	"sync"
	"time"
)

const (
	DefaultClientSetSize = 3

	waitConsumerDuration = 20 * time.Second
)

type Client struct {
	// Mutex is used to guarantee that only a single transaction at a time can be executed.
	// For example, calling to `openBox` should be done sequentially so that
	// the nonce of each transaction is correctly used.
	// Fix #142.
	mu          sync.Mutex
	ethc        *ethclient.Client
	addressFile string
	chain       string // eth or bsc or something else
	chainId     string // mainnet or testnet

	*config

	signer *bind.TransactOpts

	subscriberWg sync.WaitGroup
}

func newClient(addressFile, chain, chainId string, opts ...clientOpt) (*Client, error) {
	client, err := rpc.Dial(chain, chainId)
	cfg := new(config)
	for _, fn := range opts {
		fn(cfg)
	}

	return &Client{
		ethc:        client,
		addressFile: addressFile,
		chain:       chain,
		chainId:     chainId,
		config:      cfg,
	}, err
}

func NewBscClient(addressFile, chainId string, opts ...clientOpt) (*Client, error) {
	return newClient(addressFile, rpc.BSC, chainId, opts...)
}

func newAuthenticatedClient(addressFile, chain, network string, secrets *clientSecrets, opts ...clientOpt) (*Client, error) {
	client, err := newClient(addressFile, chain, network, opts...)
	if err != nil {
		return nil, err
	}
	if err := client.authenticate(secrets); err != nil {
		return nil, err
	}
	return client, nil
}

func NewBscAuthenticatedClient(addressFile, chainId string, secrets *clientSecrets, opts ...clientOpt) (*Client, error) {
	return newAuthenticatedClient(addressFile, rpc.BSC, chainId, secrets, opts...)
}

func (cli *Client) authenticate(secrets *clientSecrets) error {
	chainID, _ := cli.ethc.ChainID(context.Background())
	wallet, err := hdwallet.NewFromMnemonic(secrets.Mnemonic)
	if err != nil {
		return err
	}
	path := hdwallet.MustParseDerivationPath(secrets.DerivationPathOwner)
	account, err := wallet.Derive(path, true)
	if err != nil {
		return err
	}
	privateKey, err := wallet.PrivateKey(account)
	if err != nil {
		return err
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return err
	}
	cli.signer = auth
	return nil
}

func (cli *Client) AddressFile() string {
	return cli.addressFile
}

func (cli *Client) Transact(method func(opts *bind.TransactOpts) (*types.Transaction, error)) (*types.Transaction, error) {
	cli.mu.Lock()
	defer cli.mu.Unlock()
	return method(cli.signer)
}

func (cli *Client) Client() *ethclient.Client {
	return cli.ethc
}

func (cli *Client) Close() {
	cli.ethc.Close()
	// wait until all subscribers are finished
	util.WaitUntil(cli.subscriberWg.Wait, waitConsumerDuration)
}

func (cli *Client) SubscribeAll(ctx context.Context, addresses ...string) {
	// setting up whitelist
	err := whitelist.Setup(cli.addressFile)
	if err != nil {
		cli.logger.Fatal(err.Error(), logContext("error while setting up whitelist contracts"))
	}

	// setting up context
	ctx = context.WithValue(ctx, orbitContext.KeyAddressFile, cli.addressFile)
	ctx = context.WithValue(ctx, orbitContext.KeyChainId, cli.chainId)
	ctx = context.WithValue(ctx, orbitContext.KeyDB, cli.sqlDB)

	subscribeMap := make(map[common.Address]event.ParserHandler)
	for k := range event.ParserHandlers {
		if len(addresses) > 0 && util.StrSliceContains(addresses, k) {
			// ignore
			continue
		}
		address, err := util.GetContractAddress(cli.addressFile, k)
		if err != nil {
			cli.logger.Fatal("Cannot get address for handler", zap.Error(err), zap.String("contract", k))
		}
		handler := event.ParserHandlers[k]
		if handler == nil {
			cli.logger.Fatal("contract is not supported.", zap.String("contract", k))
		}
		subscribeMap[common.HexToAddress(address)] = handler
	}

	if len(subscribeMap) > 0 {
		cli.subscriberWg.Add(1)
		go func() {
			defer cli.subscriberWg.Done()
			cli.logger.Info("subscribing to event logs")
			cli.subscribe(ctx, subscribeMap)
		}()
	} else {
		cli.logger.Info("there is nothing to subscribe")
	}
}

// subscribe a single contract
func (cli *Client) subscribe(ctx context.Context, subscribeMap map[common.Address]event.ParserHandler) {
	clientSet, err := cli.clientFactory("c-", DefaultClientSetSize)
	if err != nil {
		logger.Fatalw("Cannot establish connection to rpc endpoint(s)", "error", err)
	}
	crawler := newEventScanner(clientSet, DefaultPollBatchSize, true)
	var addresses []common.Address
	for k := range subscribeMap {
		addresses = append(addresses, k)
	}
	// fetch latest block index from metadata logs
	metadataRepo := repository.NewMetadataRepository(cli.sqlDB)
	metadataService := service.NewMetadataService(metadataRepo)
	var metaNames []string
	for _, v := range addresses {
		metaNames = append(metaNames, fmt.Sprintf("%s:%s", model.MetadataCurrentBlock, v.Hex()))
	}
	metadata, err := metadataService.FindMetadata(metaNames)
	if err != nil {
		logger.Fatalw("Cannot got metadata from db", "error", err)
	}
	var oldestBlockNum uint64 = 0
	var metaLog = make(map[common.Address]struct {
		Log types.Log
		ID  uint
	})
	for _, m := range metadata {
		ethLog := types.Log{}
		err := json.Unmarshal([]byte(m.Value), &ethLog)
		if err != nil {
			logger.Warnw("malformed data, unable to parse latest event log from metadata", "name", m.Name, "value", m.Value, "error", err)
			continue
		}
		if oldestBlockNum == 0 {
			oldestBlockNum = ethLog.BlockNumber
		} else if oldestBlockNum > ethLog.BlockNumber {
			oldestBlockNum = ethLog.BlockNumber
		}
		metaLog[ethLog.Address] = struct {
			Log types.Log
			ID  uint
		}{Log: ethLog, ID: m.ID}
	}

	sub := make(chan *RichEventLog, EventLogChanBufferSize)
	wg := sync.WaitGroup{}
	var eventDispatchWorkers []*eventDispatchWorker
	for k, v := range subscribeMap {
		eventDispatchWorkers = append(eventDispatchWorkers,
			newEventDispatchWorker(k, metaLog[k].Log.BlockNumber, metaLog[k].Log.Index, v,
				EventWorkerWithFallback(func(l types.Log, err error) error {
					if err != nil {
						d, _ := l.MarshalJSON()
						logger.Errorw("error when parse handle",
							"contract", k.Hex(), "log", string(d), "error", err)
						return err
					} else {
						// store metadata
						data, _ := json.Marshal(l)
						if err := metadataService.AddOrUpdateCurrentBlock(&model.Metadata{
							Model: model.Model{ID: metaLog[k].ID},
							Name:  fmt.Sprintf("%s:%s", model.MetadataCurrentBlock, k.Hex()),
							Value: string(data),
						}); err != nil {
							logger.Errorw("error when store metadata, ignore to continue consume", "error", err, "meta", metaLog[k].ID, "value", string(data))
						}
					}
					return nil
				}),
			),
		)
	}
	for idx := range eventDispatchWorkers {
		wg.Add(1)
		w := eventDispatchWorkers[idx]
		go func() {
			defer wg.Done()
			w.Start(ctx)
		}()
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for richLog := range sub {
			for _, w := range eventDispatchWorkers {
				if w.IsRunning() {
					select {
					case w.Source() <- richLog:
					case <-time.After(w.Timeout()):
						logger.Warnw("event dispatcher worker consumes too slow, temporarily stop dispatch event to it", "name", w.Name())
						w.Pause()
					}
				}
			}
		}
	}()

	var fromBlock *big.Int = nil
	if oldestBlockNum != 0 {
		fromBlock = big.NewInt(int64(oldestBlockNum))
	}
	crawler.Poll(ctx, addresses, fromBlock, nil, sub)
	logger.Infof("Stopping polling")
	close(sub)

	// wait til all sub log is processed
	wg.Wait()
	logger.Infof("Finish subscribe")

	//contractAddress := common.HexToAddress(address)
	//query := ethereum.FilterQuery{
	//	Addresses: []common.Address{
	//		contractAddress,
	//	},
	//}
	//
	//logs := make(chan types.Log)
	//sub, err := cli.ethc.SubscribeFilterLogs(context.Background(), query, logs)
	//if err != nil {
	//	cli.logger.Fatal(err.Error(), zap.String("contract", name))
	//}
	//defer sub.Unsubscribe()
	//
	//var wg sync.WaitGroup
	//for {
	//	select {
	//	case <-ctx.Done():
	//		wg.Wait()
	//		return
	//	case err := <-sub.Err():
	//		cli.logger.Error(err.Error(), logContext("error while listening to events"))
	//	case vLog := <-logs:
	//		wg.Add(1)
	//		go func(vLog types.Log) {
	//			defer wg.Done()
	//			// FIX: there are delays between provider's nodes when listening and making query to blockchain
	//			// SOLUTION: sleep for 10 seconds before processing the log
	//			time.Sleep(10 * time.Second)
	//
	//			err := handler.ParseHandle(ctx, &vLog)
	//			if err != nil {
	//				cli.logger.Error(err.Error(), logContext("error while handling events"))
	//			}
	//		}(vLog)
	//	}
	//}
}

func (c *Client) clientFactory(prefix string, numClient int) (map[string]*ethclient.Client, error) {
	m := make(map[string]*ethclient.Client)
	for i := 0; i < numClient; i++ {
		ethc, err := rpc.Dial(c.chain, c.chainId)
		if err != nil {
			return nil, err
		}

		m[fmt.Sprintf("%s%d", prefix, i)] = ethc
	}
	return m, nil
}
