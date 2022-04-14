package ido

import (
	"context"
	"errors"
	"math/big"
	"orbit_nft/contract/abi/ido"
	orbitContext "orbit_nft/contract/context"
	"orbit_nft/contract/event"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

const LogBuyName string = "Buy"
const LogBuySig string = "Buy(address,uint256)"

type LogBuy struct {
	Buyer  common.Address
	Amount *big.Int
}

func ParseLogBuy(contractAbi *abi.ABI, vLog *types.Log) (*LogBuy, error) {
	var evt LogBuy
	if len(vLog.Data) > 0 {
		err := contractAbi.UnpackIntoInterface(&evt, LogBuyName, vLog.Data)
		if err != nil {
			return nil, err
		}
	}

	evt.Buyer = common.HexToAddress(vLog.Topics[1].Hex())

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

const LogPublicSaleOpenedName string = "PublicSaleOpened"
const LogPublicSaleOpenedSig string = "PublicSaleOpened()"

type LogPublicSaleOpened struct {
}

func ParseLogPublicSaleOpened(contractAbi *abi.ABI, vLog *types.Log) (*LogPublicSaleOpened, error) {
	var evt LogPublicSaleOpened
	if len(vLog.Data) > 0 {
		err := contractAbi.UnpackIntoInterface(&evt, LogPublicSaleOpenedName, vLog.Data)
		if err != nil {
			return nil, err
		}
	}

	return &evt, nil
}

const LogRoleAdminChangedName string = "RoleAdminChanged"
const LogRoleAdminChangedSig string = "RoleAdminChanged(bytes32,bytes32,bytes32)"

type LogRoleAdminChanged struct {
	Role              common.Hash
	PreviousAdminRole common.Hash
	NewAdminRole      common.Hash
}

func ParseLogRoleAdminChanged(contractAbi *abi.ABI, vLog *types.Log) (*LogRoleAdminChanged, error) {
	var evt LogRoleAdminChanged
	if len(vLog.Data) > 0 {
		err := contractAbi.UnpackIntoInterface(&evt, LogRoleAdminChangedName, vLog.Data)
		if err != nil {
			return nil, err
		}
	}

	evt.Role = common.HexToHash(vLog.Topics[1].Hex())

	evt.PreviousAdminRole = common.HexToHash(vLog.Topics[2].Hex())

	evt.NewAdminRole = common.HexToHash(vLog.Topics[3].Hex())

	return &evt, nil
}

const LogRoleGrantedName string = "RoleGranted"
const LogRoleGrantedSig string = "RoleGranted(bytes32,address,address)"

type LogRoleGranted struct {
	Role    common.Hash
	Account common.Address
	Sender  common.Address
}

func ParseLogRoleGranted(contractAbi *abi.ABI, vLog *types.Log) (*LogRoleGranted, error) {
	var evt LogRoleGranted
	if len(vLog.Data) > 0 {
		err := contractAbi.UnpackIntoInterface(&evt, LogRoleGrantedName, vLog.Data)
		if err != nil {
			return nil, err
		}
	}

	evt.Role = common.HexToHash(vLog.Topics[1].Hex())

	evt.Account = common.HexToAddress(vLog.Topics[2].Hex())

	evt.Sender = common.HexToAddress(vLog.Topics[3].Hex())

	return &evt, nil
}

const LogRoleRevokedName string = "RoleRevoked"
const LogRoleRevokedSig string = "RoleRevoked(bytes32,address,address)"

type LogRoleRevoked struct {
	Role    common.Hash
	Account common.Address
	Sender  common.Address
}

func ParseLogRoleRevoked(contractAbi *abi.ABI, vLog *types.Log) (*LogRoleRevoked, error) {
	var evt LogRoleRevoked
	if len(vLog.Data) > 0 {
		err := contractAbi.UnpackIntoInterface(&evt, LogRoleRevokedName, vLog.Data)
		if err != nil {
			return nil, err
		}
	}

	evt.Role = common.HexToHash(vLog.Topics[1].Hex())

	evt.Account = common.HexToAddress(vLog.Topics[2].Hex())

	evt.Sender = common.HexToAddress(vLog.Topics[3].Hex())

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

const LogSubscriberRegisteredName string = "SubscriberRegistered"
const LogSubscriberRegisteredSig string = "SubscriberRegistered(address)"

type LogSubscriberRegistered struct {
	Subscriber common.Address
}

func ParseLogSubscriberRegistered(contractAbi *abi.ABI, vLog *types.Log) (*LogSubscriberRegistered, error) {
	var evt LogSubscriberRegistered
	if len(vLog.Data) > 0 {
		err := contractAbi.UnpackIntoInterface(&evt, LogSubscriberRegisteredName, vLog.Data)
		if err != nil {
			return nil, err
		}
	}

	evt.Subscriber = common.HexToAddress(vLog.Topics[1].Hex())

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

type parserHandler struct {
	name string
}

var _ event.ParserHandler = (*parserHandler)(nil)

func (p *parserHandler) Name() string {
	return p.name
}

func (p *parserHandler) ParseHandle(ctx context.Context, vLog *types.Log) error {
	contractAbi, err := abi.JSON(strings.NewReader(string(ido.ABI)))
	if err != nil {
		return err
	}

	ctx = context.WithValue(ctx, orbitContext.KeyTxHash, vLog.TxHash)

	for k, v := range hashMap {
		if k == vLog.Topics[0].Hex() {
			parser := logParserMap[v]
			handler := logHandlerMap[v]
			switch v {
			case LogBuyName:
				evt, err := parser.(func(*abi.ABI, *types.Log) (*LogBuy, error))(&contractAbi, vLog)
				if err != nil {
					return err
				}
				return handler.(func(context.Context, *LogBuy) error)(ctx, evt)
			case LogPausedName:
				evt, err := parser.(func(*abi.ABI, *types.Log) (*LogPaused, error))(&contractAbi, vLog)
				if err != nil {
					return err
				}
				return handler.(func(context.Context, *LogPaused) error)(ctx, evt)
			case LogPublicSaleOpenedName:
				evt, err := parser.(func(*abi.ABI, *types.Log) (*LogPublicSaleOpened, error))(&contractAbi, vLog)
				if err != nil {
					return err
				}
				return handler.(func(context.Context, *LogPublicSaleOpened) error)(ctx, evt)
			case LogRoleAdminChangedName:
				evt, err := parser.(func(*abi.ABI, *types.Log) (*LogRoleAdminChanged, error))(&contractAbi, vLog)
				if err != nil {
					return err
				}
				return handler.(func(context.Context, *LogRoleAdminChanged) error)(ctx, evt)
			case LogRoleGrantedName:
				evt, err := parser.(func(*abi.ABI, *types.Log) (*LogRoleGranted, error))(&contractAbi, vLog)
				if err != nil {
					return err
				}
				return handler.(func(context.Context, *LogRoleGranted) error)(ctx, evt)
			case LogRoleRevokedName:
				evt, err := parser.(func(*abi.ABI, *types.Log) (*LogRoleRevoked, error))(&contractAbi, vLog)
				if err != nil {
					return err
				}
				return handler.(func(context.Context, *LogRoleRevoked) error)(ctx, evt)
			case LogStoppedName:
				evt, err := parser.(func(*abi.ABI, *types.Log) (*LogStopped, error))(&contractAbi, vLog)
				if err != nil {
					return err
				}
				return handler.(func(context.Context, *LogStopped) error)(ctx, evt)
			case LogSubscriberRegisteredName:
				evt, err := parser.(func(*abi.ABI, *types.Log) (*LogSubscriberRegistered, error))(&contractAbi, vLog)
				if err != nil {
					return err
				}
				return handler.(func(context.Context, *LogSubscriberRegistered) error)(ctx, evt)
			case LogUnpausedName:
				evt, err := parser.(func(*abi.ABI, *types.Log) (*LogUnpaused, error))(&contractAbi, vLog)
				if err != nil {
					return err
				}
				return handler.(func(context.Context, *LogUnpaused) error)(ctx, evt)
			default:
				return errors.New("ido: event type is not supported")
			}
		}
	}
	return errors.New("ido: topic is not supported")
}

var logParserMap map[string]interface{}
var logHandlerMap map[string]interface{}
var hashMap map[string]string

func init() {
	logParserMap = make(map[string]interface{})
	logHandlerMap = make(map[string]interface{})
	hashMap = make(map[string]string)
	logParserMap[LogBuyName] = ParseLogBuy
	logParserMap[LogPausedName] = ParseLogPaused
	logParserMap[LogPublicSaleOpenedName] = ParseLogPublicSaleOpened
	logParserMap[LogRoleAdminChangedName] = ParseLogRoleAdminChanged
	logParserMap[LogRoleGrantedName] = ParseLogRoleGranted
	logParserMap[LogRoleRevokedName] = ParseLogRoleRevoked
	logParserMap[LogStoppedName] = ParseLogStopped
	logParserMap[LogSubscriberRegisteredName] = ParseLogSubscriberRegistered
	logParserMap[LogUnpausedName] = ParseLogUnpaused

	logHandlerMap[LogBuyName] = HandleLogBuy
	logHandlerMap[LogPausedName] = HandleLogPaused
	logHandlerMap[LogPublicSaleOpenedName] = HandleLogPublicSaleOpened
	logHandlerMap[LogRoleAdminChangedName] = HandleLogRoleAdminChanged
	logHandlerMap[LogRoleGrantedName] = HandleLogRoleGranted
	logHandlerMap[LogRoleRevokedName] = HandleLogRoleRevoked
	logHandlerMap[LogStoppedName] = HandleLogStopped
	logHandlerMap[LogSubscriberRegisteredName] = HandleLogSubscriberRegistered
	logHandlerMap[LogUnpausedName] = HandleLogUnpaused

	logBuySigHash := crypto.Keccak256Hash([]byte(LogBuySig))
	hashMap[logBuySigHash.Hex()] = LogBuyName

	logPausedSigHash := crypto.Keccak256Hash([]byte(LogPausedSig))
	hashMap[logPausedSigHash.Hex()] = LogPausedName

	logPublicSaleOpenedSigHash := crypto.Keccak256Hash([]byte(LogPublicSaleOpenedSig))
	hashMap[logPublicSaleOpenedSigHash.Hex()] = LogPublicSaleOpenedName

	logRoleAdminChangedSigHash := crypto.Keccak256Hash([]byte(LogRoleAdminChangedSig))
	hashMap[logRoleAdminChangedSigHash.Hex()] = LogRoleAdminChangedName

	logRoleGrantedSigHash := crypto.Keccak256Hash([]byte(LogRoleGrantedSig))
	hashMap[logRoleGrantedSigHash.Hex()] = LogRoleGrantedName

	logRoleRevokedSigHash := crypto.Keccak256Hash([]byte(LogRoleRevokedSig))
	hashMap[logRoleRevokedSigHash.Hex()] = LogRoleRevokedName

	logStoppedSigHash := crypto.Keccak256Hash([]byte(LogStoppedSig))
	hashMap[logStoppedSigHash.Hex()] = LogStoppedName

	logSubscriberRegisteredSigHash := crypto.Keccak256Hash([]byte(LogSubscriberRegisteredSig))
	hashMap[logSubscriberRegisteredSigHash.Hex()] = LogSubscriberRegisteredName

	logUnpausedSigHash := crypto.Keccak256Hash([]byte(LogUnpausedSig))
	hashMap[logUnpausedSigHash.Hex()] = LogUnpausedName

	event.Register(&parserHandler{
		name: ido.Name,
	})
}
