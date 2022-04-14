package contract

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"go.uber.org/atomic"
	orbitContext "orbit_nft/contract/context"
	"orbit_nft/contract/event"
	"orbit_nft/logger"
	"time"
)

type EventFallbackHandler func(log types.Log, err error) error

type eventDispatchWorker struct {
	addr            common.Address
	fromBlockNumber uint64
	fromLogIndex    uint
	source          chan *RichEventLog
	exit            chan bool
	handler         event.ParserHandler
	fallback        EventFallbackHandler
	isRunning       *atomic.Bool
	timeout         time.Duration
}

type eventDispatchWorkerOption func(worker *eventDispatchWorker)

func EventWorkerWithFallback(f EventFallbackHandler) eventDispatchWorkerOption {
	return func(w *eventDispatchWorker) {
		w.fallback = f
	}
}

func EventWorkerWithTimeout(d time.Duration) eventDispatchWorkerOption {
	return func(w *eventDispatchWorker) {
		w.timeout = d
	}
}

func newEventDispatchWorker(address common.Address, fromBlock uint64, fromLogIndex uint, f event.ParserHandler, opts ...eventDispatchWorkerOption) *eventDispatchWorker {
	w := &eventDispatchWorker{
		addr:            address,
		fromBlockNumber: fromBlock,
		fromLogIndex:    fromLogIndex,
		source:          make(chan *RichEventLog, DefaultPollBatchSize),
		exit:            make(chan bool),
		handler:         f,
		isRunning:       atomic.NewBool(false),
		timeout:         DefaultEventWorkerTimeout,
	}

	for _, opt := range opts {
		opt(w)
	}

	return w
}

func (w *eventDispatchWorker) Start(ctx context.Context) {
	var lastErr error = nil
	w.isRunning.Store(true)
	for {
		select {
		case <-ctx.Done():
			return
		case richLog := <-w.source:
			if !w.isRunning.Load() {
				logger.Infow("Worker are pausing", "address", w.addr.Hex(), "lastError", lastErr)
				continue
			}
			// pre validate scan range
			if richLog.EndBlock < w.fromBlockNumber {
				continue
			}
			if len(richLog.Logs) == 0 && w.fallback != nil {
				// still fallback with a log with only blockNumber, logIndex and address for tracking process
				_ = w.fallback(types.Log{
					Address:     w.addr,
					BlockNumber: richLog.EndBlock,
					Index:       0,
				}, nil)
			}
			// FIX: there are delays between provider's nodes when listening and making query to blockchain
			// SOLUTION: sleep for 10 seconds before processing the log
			delay := time.Since(richLog.ScanTime)
			if delay < HandlerWaitDuration {
				time.Sleep(HandlerWaitDuration - delay)
			}
			func() {
				cctx, f := context.WithCancel(context.WithValue(ctx, orbitContext.KeyEthEndpoint, richLog.Network))
				defer f()
				for _, l := range richLog.Logs {
					if w.addr != l.Address {
						continue
					}
					// validate scan range
					if l.BlockNumber < w.fromBlockNumber {
						continue
					} else if l.BlockNumber == w.fromBlockNumber {
						if l.Index <= w.fromLogIndex {
							continue
						}
					}
					lastErr = w.handler.ParseHandle(cctx, &l)
					if w.fallback != nil {
						if err := w.fallback(l, lastErr); err != nil {
							w.isRunning.Store(false)
						}
					}
				}
			}()
		}
	}
}

func (w *eventDispatchWorker) IsRunning() bool {
	return w.isRunning.Load()
}

func (w *eventDispatchWorker) Resume() {
	w.isRunning.Store(true)
}

func (w *eventDispatchWorker) Pause() {
	w.isRunning.Store(false)
}

func (w *eventDispatchWorker) Source() chan *RichEventLog {
	return w.source
}

func (w *eventDispatchWorker) Name() string {
	return w.addr.Hex()
}

func (w *eventDispatchWorker) Address() common.Address {
	return w.addr
}

func (w *eventDispatchWorker) Timeout() time.Duration {
	return w.timeout
}
