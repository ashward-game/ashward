package openboxgenesis

import (
	"context"
	"errors"
	"math/big"
	"orbit_nft/contract/abi/openboxgenesis"
	orbitContext "orbit_nft/contract/context"
	"orbit_nft/contract/event"
	"orbit_nft/util"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

const LogBoxBoughtName string = "BoxBought"
const LogBoxBoughtSig string = "BoxBought(address,uint8,bytes32,bytes32)"

type LogBoxBought struct {
	Buyer        common.Address
	BoxGrade     uint8
	ServerHash   common.Hash
	ClientRandom common.Hash
}

func ParseLogBoxBought(contractAbi *abi.ABI, vLog *types.Log) (*LogBoxBought, error) {
	var evt LogBoxBought
	if len(vLog.Data) > 0 {
		err := contractAbi.UnpackIntoInterface(&evt, LogBoxBoughtName, vLog.Data)
		if err != nil {
			return nil, err
		}
	}

	evt.Buyer = common.HexToAddress(vLog.Topics[1].Hex())

	evt.BoxGrade = uint8(util.HexToBigInt(vLog.Topics[2].Hex()).Uint64())

	return &evt, nil
}

const LogBoxOpenedName string = "BoxOpened"
const LogBoxOpenedSig string = "BoxOpened(address,uint8,bytes32,bool,uint256)"

type LogBoxOpened struct {
	Buyer      common.Address
	BoxGrade   uint8
	ServerHash common.Hash
	IsEmpty    bool
	TokenID    *big.Int
}

func ParseLogBoxOpened(contractAbi *abi.ABI, vLog *types.Log) (*LogBoxOpened, error) {
	var evt LogBoxOpened
	if len(vLog.Data) > 0 {
		err := contractAbi.UnpackIntoInterface(&evt, LogBoxOpenedName, vLog.Data)
		if err != nil {
			return nil, err
		}
	}

	evt.Buyer = common.HexToAddress(vLog.Topics[1].Hex())

	evt.BoxGrade = uint8(util.HexToBigInt(vLog.Topics[2].Hex()).Uint64())

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

type parserHandler struct {
	name string
}

var _ event.ParserHandler = (*parserHandler)(nil)

func (p *parserHandler) Name() string {
	return p.name
}

func (p *parserHandler) ParseHandle(ctx context.Context, vLog *types.Log) error {
	contractAbi, err := abi.JSON(strings.NewReader(string(openboxgenesis.ABI)))
	if err != nil {
		return err
	}

	ctx = context.WithValue(ctx, orbitContext.KeyTxHash, vLog.TxHash)

	for k, v := range hashMap {
		if k == vLog.Topics[0].Hex() {
			parser := logParserMap[v]
			handler := logHandlerMap[v]
			switch v {
			case LogBoxBoughtName:
				evt, err := parser.(func(*abi.ABI, *types.Log) (*LogBoxBought, error))(&contractAbi, vLog)
				if err != nil {
					return err
				}
				return handler.(func(context.Context, *LogBoxBought) error)(ctx, evt)
			case LogBoxOpenedName:
				evt, err := parser.(func(*abi.ABI, *types.Log) (*LogBoxOpened, error))(&contractAbi, vLog)
				if err != nil {
					return err
				}
				return handler.(func(context.Context, *LogBoxOpened) error)(ctx, evt)
			case LogOwnershipTransferredName:
				evt, err := parser.(func(*abi.ABI, *types.Log) (*LogOwnershipTransferred, error))(&contractAbi, vLog)
				if err != nil {
					return err
				}
				return handler.(func(context.Context, *LogOwnershipTransferred) error)(ctx, evt)
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
			case LogSubscriberRegisteredName:
				evt, err := parser.(func(*abi.ABI, *types.Log) (*LogSubscriberRegistered, error))(&contractAbi, vLog)
				if err != nil {
					return err
				}
				return handler.(func(context.Context, *LogSubscriberRegistered) error)(ctx, evt)
			default:
				return errors.New("openboxgenesis: event type is not supported")
			}
		}
	}
	return errors.New("openboxgenesis: topic is not supported")
}

var logParserMap map[string]interface{}
var logHandlerMap map[string]interface{}
var hashMap map[string]string

func init() {
	logParserMap = make(map[string]interface{})
	logHandlerMap = make(map[string]interface{})
	hashMap = make(map[string]string)
	logParserMap[LogBoxBoughtName] = ParseLogBoxBought
	logParserMap[LogBoxOpenedName] = ParseLogBoxOpened
	logParserMap[LogOwnershipTransferredName] = ParseLogOwnershipTransferred
	logParserMap[LogPublicSaleOpenedName] = ParseLogPublicSaleOpened
	logParserMap[LogRoleAdminChangedName] = ParseLogRoleAdminChanged
	logParserMap[LogRoleGrantedName] = ParseLogRoleGranted
	logParserMap[LogRoleRevokedName] = ParseLogRoleRevoked
	logParserMap[LogSubscriberRegisteredName] = ParseLogSubscriberRegistered

	logHandlerMap[LogBoxBoughtName] = HandleLogBoxBought
	logHandlerMap[LogBoxOpenedName] = HandleLogBoxOpened
	logHandlerMap[LogOwnershipTransferredName] = HandleLogOwnershipTransferred
	logHandlerMap[LogPublicSaleOpenedName] = HandleLogPublicSaleOpened
	logHandlerMap[LogRoleAdminChangedName] = HandleLogRoleAdminChanged
	logHandlerMap[LogRoleGrantedName] = HandleLogRoleGranted
	logHandlerMap[LogRoleRevokedName] = HandleLogRoleRevoked
	logHandlerMap[LogSubscriberRegisteredName] = HandleLogSubscriberRegistered

	logBoxBoughtSigHash := crypto.Keccak256Hash([]byte(LogBoxBoughtSig))
	hashMap[logBoxBoughtSigHash.Hex()] = LogBoxBoughtName

	logBoxOpenedSigHash := crypto.Keccak256Hash([]byte(LogBoxOpenedSig))
	hashMap[logBoxOpenedSigHash.Hex()] = LogBoxOpenedName

	logOwnershipTransferredSigHash := crypto.Keccak256Hash([]byte(LogOwnershipTransferredSig))
	hashMap[logOwnershipTransferredSigHash.Hex()] = LogOwnershipTransferredName

	logPublicSaleOpenedSigHash := crypto.Keccak256Hash([]byte(LogPublicSaleOpenedSig))
	hashMap[logPublicSaleOpenedSigHash.Hex()] = LogPublicSaleOpenedName

	logRoleAdminChangedSigHash := crypto.Keccak256Hash([]byte(LogRoleAdminChangedSig))
	hashMap[logRoleAdminChangedSigHash.Hex()] = LogRoleAdminChangedName

	logRoleGrantedSigHash := crypto.Keccak256Hash([]byte(LogRoleGrantedSig))
	hashMap[logRoleGrantedSigHash.Hex()] = LogRoleGrantedName

	logRoleRevokedSigHash := crypto.Keccak256Hash([]byte(LogRoleRevokedSig))
	hashMap[logRoleRevokedSigHash.Hex()] = LogRoleRevokedName

	logSubscriberRegisteredSigHash := crypto.Keccak256Hash([]byte(LogSubscriberRegisteredSig))
	hashMap[logSubscriberRegisteredSigHash.Hex()] = LogSubscriberRegisteredName

	event.Register(&parserHandler{
		name: openboxgenesis.Name,
	})
}
