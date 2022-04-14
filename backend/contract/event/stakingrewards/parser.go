package stakingrewards

import (
	"context"
	"errors"
	"math/big"
	"orbit_nft/contract/abi/stakingrewards"
	orbitContext "orbit_nft/contract/context"
	"orbit_nft/contract/event"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

const LogExitedName string = "Exited"
const LogExitedSig string = "Exited(address)"

type LogExited struct {
	User common.Address
}

func ParseLogExited(contractAbi *abi.ABI, vLog *types.Log) (*LogExited, error) {
	var evt LogExited
	if len(vLog.Data) > 0 {
		err := contractAbi.UnpackIntoInterface(&evt, LogExitedName, vLog.Data)
		if err != nil {
			return nil, err
		}
	}

	evt.User = common.HexToAddress(vLog.Topics[1].Hex())

	return &evt, nil
}

const LogOwnershipTransferredName string = "OwnershipTransferred"
const LogOwnershipTransferredSig string = "OwnershipTransferred(address,address)"

type LogOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
}

func ParseLogOwnershipTransferred(contractAbi *abi.ABI, vLog *types.Log) (*LogOwnershipTransferred, error) {
	var evt LogOwnershipTransferred
	if len(vLog.Data) > 0 {
		err := contractAbi.UnpackIntoInterface(&evt, LogOwnershipTransferredName, vLog.Data)
		if err != nil {
			return nil, err
		}
	}

	evt.PreviousOwner = common.HexToAddress(vLog.Topics[1].Hex())

	evt.NewOwner = common.HexToAddress(vLog.Topics[2].Hex())

	return &evt, nil
}

const LogPausedName string = "Paused"
const LogPausedSig string = "Paused(address)"

type LogPaused struct {
	Account common.Address
}

func ParseLogPaused(contractAbi *abi.ABI, vLog *types.Log) (*LogPaused, error) {
	var evt LogPaused
	if len(vLog.Data) > 0 {
		err := contractAbi.UnpackIntoInterface(&evt, LogPausedName, vLog.Data)
		if err != nil {
			return nil, err
		}
	}

	return &evt, nil
}

const LogRewardsPaidName string = "RewardsPaid"
const LogRewardsPaidSig string = "RewardsPaid(address,uint256)"

type LogRewardsPaid struct {
	User   common.Address
	Amount *big.Int
}

func ParseLogRewardsPaid(contractAbi *abi.ABI, vLog *types.Log) (*LogRewardsPaid, error) {
	var evt LogRewardsPaid
	if len(vLog.Data) > 0 {
		err := contractAbi.UnpackIntoInterface(&evt, LogRewardsPaidName, vLog.Data)
		if err != nil {
			return nil, err
		}
	}

	evt.User = common.HexToAddress(vLog.Topics[1].Hex())

	return &evt, nil
}

const LogStakedName string = "Staked"
const LogStakedSig string = "Staked(address,uint256)"

type LogStaked struct {
	User   common.Address
	Amount *big.Int
}

func ParseLogStaked(contractAbi *abi.ABI, vLog *types.Log) (*LogStaked, error) {
	var evt LogStaked
	if len(vLog.Data) > 0 {
		err := contractAbi.UnpackIntoInterface(&evt, LogStakedName, vLog.Data)
		if err != nil {
			return nil, err
		}
	}

	evt.User = common.HexToAddress(vLog.Topics[1].Hex())

	return &evt, nil
}

const LogStoppedName string = "Stopped"
const LogStoppedSig string = "Stopped()"

type LogStopped struct {
}

func ParseLogStopped(contractAbi *abi.ABI, vLog *types.Log) (*LogStopped, error) {
	var evt LogStopped
	if len(vLog.Data) > 0 {
		err := contractAbi.UnpackIntoInterface(&evt, LogStoppedName, vLog.Data)
		if err != nil {
			return nil, err
		}
	}

	return &evt, nil
}

const LogTokensApprovedName string = "TokensApproved"
const LogTokensApprovedSig string = "TokensApproved(address,uint256,bytes)"

type LogTokensApproved struct {
	Sender common.Address
	Amount *big.Int
	Data   []byte
}

func ParseLogTokensApproved(contractAbi *abi.ABI, vLog *types.Log) (*LogTokensApproved, error) {
	var evt LogTokensApproved
	if len(vLog.Data) > 0 {
		err := contractAbi.UnpackIntoInterface(&evt, LogTokensApprovedName, vLog.Data)
		if err != nil {
			return nil, err
		}
	}

	evt.Sender = common.HexToAddress(vLog.Topics[1].Hex())

	return &evt, nil
}

const LogTokensReceivedName string = "TokensReceived"
const LogTokensReceivedSig string = "TokensReceived(address,address,uint256,bytes)"

type LogTokensReceived struct {
	Operator common.Address
	Sender   common.Address
	Amount   *big.Int
	Data     []byte
}

func ParseLogTokensReceived(contractAbi *abi.ABI, vLog *types.Log) (*LogTokensReceived, error) {
	var evt LogTokensReceived
	if len(vLog.Data) > 0 {
		err := contractAbi.UnpackIntoInterface(&evt, LogTokensReceivedName, vLog.Data)
		if err != nil {
			return nil, err
		}
	}

	evt.Operator = common.HexToAddress(vLog.Topics[1].Hex())

	evt.Sender = common.HexToAddress(vLog.Topics[2].Hex())

	return &evt, nil
}

const LogUnpausedName string = "Unpaused"
const LogUnpausedSig string = "Unpaused(address)"

type LogUnpaused struct {
	Account common.Address
}

func ParseLogUnpaused(contractAbi *abi.ABI, vLog *types.Log) (*LogUnpaused, error) {
	var evt LogUnpaused
	if len(vLog.Data) > 0 {
		err := contractAbi.UnpackIntoInterface(&evt, LogUnpausedName, vLog.Data)
		if err != nil {
			return nil, err
		}
	}

	return &evt, nil
}

const LogWithdrawnName string = "Withdrawn"
const LogWithdrawnSig string = "Withdrawn(address,uint256)"

type LogWithdrawn struct {
	User   common.Address
	Amount *big.Int
}

func ParseLogWithdrawn(contractAbi *abi.ABI, vLog *types.Log) (*LogWithdrawn, error) {
	var evt LogWithdrawn
	if len(vLog.Data) > 0 {
		err := contractAbi.UnpackIntoInterface(&evt, LogWithdrawnName, vLog.Data)
		if err != nil {
			return nil, err
		}
	}

	evt.User = common.HexToAddress(vLog.Topics[1].Hex())

	return &evt, nil
}

type parserHandler struct {
	name string
}

var _ event.ParserHandler = (*parserHandler)(nil)

func (p *parserHandler) Name() string {
	return p.name
}

func (p *parserHandler) ParseHandle(ctx context.Context, vLog *types.Log) error {
	contractAbi, err := abi.JSON(strings.NewReader(string(stakingrewards.ABI)))
	if err != nil {
		return err
	}

	ctx = context.WithValue(ctx, orbitContext.KeyTxHash, vLog.TxHash)

	for k, v := range hashMap {
		if k == vLog.Topics[0].Hex() {
			parser := logParserMap[v]
			handler := logHandlerMap[v]
			switch v {
			case LogExitedName:
				evt, err := parser.(func(*abi.ABI, *types.Log) (*LogExited, error))(&contractAbi, vLog)
				if err != nil {
					return err
				}
				return handler.(func(context.Context, *LogExited) error)(ctx, evt)
			case LogOwnershipTransferredName:
				evt, err := parser.(func(*abi.ABI, *types.Log) (*LogOwnershipTransferred, error))(&contractAbi, vLog)
				if err != nil {
					return err
				}
				return handler.(func(context.Context, *LogOwnershipTransferred) error)(ctx, evt)
			case LogPausedName:
				evt, err := parser.(func(*abi.ABI, *types.Log) (*LogPaused, error))(&contractAbi, vLog)
				if err != nil {
					return err
				}
				return handler.(func(context.Context, *LogPaused) error)(ctx, evt)
			case LogRewardsPaidName:
				evt, err := parser.(func(*abi.ABI, *types.Log) (*LogRewardsPaid, error))(&contractAbi, vLog)
				if err != nil {
					return err
				}
				return handler.(func(context.Context, *LogRewardsPaid) error)(ctx, evt)
			case LogStakedName:
				evt, err := parser.(func(*abi.ABI, *types.Log) (*LogStaked, error))(&contractAbi, vLog)
				if err != nil {
					return err
				}
				return handler.(func(context.Context, *LogStaked) error)(ctx, evt)
			case LogStoppedName:
				evt, err := parser.(func(*abi.ABI, *types.Log) (*LogStopped, error))(&contractAbi, vLog)
				if err != nil {
					return err
				}
				return handler.(func(context.Context, *LogStopped) error)(ctx, evt)
			case LogTokensApprovedName:
				evt, err := parser.(func(*abi.ABI, *types.Log) (*LogTokensApproved, error))(&contractAbi, vLog)
				if err != nil {
					return err
				}
				return handler.(func(context.Context, *LogTokensApproved) error)(ctx, evt)
			case LogTokensReceivedName:
				evt, err := parser.(func(*abi.ABI, *types.Log) (*LogTokensReceived, error))(&contractAbi, vLog)
				if err != nil {
					return err
				}
				return handler.(func(context.Context, *LogTokensReceived) error)(ctx, evt)
			case LogUnpausedName:
				evt, err := parser.(func(*abi.ABI, *types.Log) (*LogUnpaused, error))(&contractAbi, vLog)
				if err != nil {
					return err
				}
				return handler.(func(context.Context, *LogUnpaused) error)(ctx, evt)
			case LogWithdrawnName:
				evt, err := parser.(func(*abi.ABI, *types.Log) (*LogWithdrawn, error))(&contractAbi, vLog)
				if err != nil {
					return err
				}
				return handler.(func(context.Context, *LogWithdrawn) error)(ctx, evt)
			default:
				return errors.New("stakingrewards: event type is not supported")
			}
		}
	}
	return errors.New("stakingrewards: topic is not supported")
}

var logParserMap map[string]interface{}
var logHandlerMap map[string]interface{}
var hashMap map[string]string

func init() {
	logParserMap = make(map[string]interface{})
	logHandlerMap = make(map[string]interface{})
	hashMap = make(map[string]string)
	logParserMap[LogExitedName] = ParseLogExited
	logParserMap[LogOwnershipTransferredName] = ParseLogOwnershipTransferred
	logParserMap[LogPausedName] = ParseLogPaused
	logParserMap[LogRewardsPaidName] = ParseLogRewardsPaid
	logParserMap[LogStakedName] = ParseLogStaked
	logParserMap[LogStoppedName] = ParseLogStopped
	logParserMap[LogTokensApprovedName] = ParseLogTokensApproved
	logParserMap[LogTokensReceivedName] = ParseLogTokensReceived
	logParserMap[LogUnpausedName] = ParseLogUnpaused
	logParserMap[LogWithdrawnName] = ParseLogWithdrawn

	logHandlerMap[LogExitedName] = HandleLogExited
	logHandlerMap[LogOwnershipTransferredName] = HandleLogOwnershipTransferred
	logHandlerMap[LogPausedName] = HandleLogPaused
	logHandlerMap[LogRewardsPaidName] = HandleLogRewardsPaid
	logHandlerMap[LogStakedName] = HandleLogStaked
	logHandlerMap[LogStoppedName] = HandleLogStopped
	logHandlerMap[LogTokensApprovedName] = HandleLogTokensApproved
	logHandlerMap[LogTokensReceivedName] = HandleLogTokensReceived
	logHandlerMap[LogUnpausedName] = HandleLogUnpaused
	logHandlerMap[LogWithdrawnName] = HandleLogWithdrawn

	logExitedSigHash := crypto.Keccak256Hash([]byte(LogExitedSig))
	hashMap[logExitedSigHash.Hex()] = LogExitedName

	logOwnershipTransferredSigHash := crypto.Keccak256Hash([]byte(LogOwnershipTransferredSig))
	hashMap[logOwnershipTransferredSigHash.Hex()] = LogOwnershipTransferredName

	logPausedSigHash := crypto.Keccak256Hash([]byte(LogPausedSig))
	hashMap[logPausedSigHash.Hex()] = LogPausedName

	logRewardsPaidSigHash := crypto.Keccak256Hash([]byte(LogRewardsPaidSig))
	hashMap[logRewardsPaidSigHash.Hex()] = LogRewardsPaidName

	logStakedSigHash := crypto.Keccak256Hash([]byte(LogStakedSig))
	hashMap[logStakedSigHash.Hex()] = LogStakedName

	logStoppedSigHash := crypto.Keccak256Hash([]byte(LogStoppedSig))
	hashMap[logStoppedSigHash.Hex()] = LogStoppedName

	logTokensApprovedSigHash := crypto.Keccak256Hash([]byte(LogTokensApprovedSig))
	hashMap[logTokensApprovedSigHash.Hex()] = LogTokensApprovedName

	logTokensReceivedSigHash := crypto.Keccak256Hash([]byte(LogTokensReceivedSig))
	hashMap[logTokensReceivedSigHash.Hex()] = LogTokensReceivedName

	logUnpausedSigHash := crypto.Keccak256Hash([]byte(LogUnpausedSig))
	hashMap[logUnpausedSigHash.Hex()] = LogUnpausedName

	logWithdrawnSigHash := crypto.Keccak256Hash([]byte(LogWithdrawnSig))
	hashMap[logWithdrawnSigHash.Hex()] = LogWithdrawnName

	event.Register(&parserHandler{
		name: stakingrewards.Name,
	})
}
