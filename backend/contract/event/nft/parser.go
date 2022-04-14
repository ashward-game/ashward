package nft

import (
	"context"
	"errors"
	"math/big"
	"orbit_nft/contract/abi/nft"
	orbitContext "orbit_nft/contract/context"
	"orbit_nft/contract/event"
	"orbit_nft/util"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

const LogApprovalName string = "Approval"
const LogApprovalSig string = "Approval(address,address,uint256)"

type LogApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
}

func ParseLogApproval(contractAbi *abi.ABI, vLog *types.Log) (*LogApproval, error) {
	var evt LogApproval
	if len(vLog.Data) > 0 {
		err := contractAbi.UnpackIntoInterface(&evt, LogApprovalName, vLog.Data)
		if err != nil {
			return nil, err
		}
	}

	evt.Owner = common.HexToAddress(vLog.Topics[1].Hex())

	evt.Approved = common.HexToAddress(vLog.Topics[2].Hex())

	evt.TokenId = util.HexToBigInt(vLog.Topics[3].Hex())

	return &evt, nil
}

const LogApprovalForAllName string = "ApprovalForAll"
const LogApprovalForAllSig string = "ApprovalForAll(address,address,bool)"

type LogApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
}

func ParseLogApprovalForAll(contractAbi *abi.ABI, vLog *types.Log) (*LogApprovalForAll, error) {
	var evt LogApprovalForAll
	if len(vLog.Data) > 0 {
		err := contractAbi.UnpackIntoInterface(&evt, LogApprovalForAllName, vLog.Data)
		if err != nil {
			return nil, err
		}
	}

	evt.Owner = common.HexToAddress(vLog.Topics[1].Hex())

	evt.Operator = common.HexToAddress(vLog.Topics[2].Hex())

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

const LogTransferName string = "Transfer"
const LogTransferSig string = "Transfer(address,address,uint256)"

type LogTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
}

func ParseLogTransfer(contractAbi *abi.ABI, vLog *types.Log) (*LogTransfer, error) {
	var evt LogTransfer
	if len(vLog.Data) > 0 {
		err := contractAbi.UnpackIntoInterface(&evt, LogTransferName, vLog.Data)
		if err != nil {
			return nil, err
		}
	}

	evt.From = common.HexToAddress(vLog.Topics[1].Hex())

	evt.To = common.HexToAddress(vLog.Topics[2].Hex())

	evt.TokenId = util.HexToBigInt(vLog.Topics[3].Hex())

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
	contractAbi, err := abi.JSON(strings.NewReader(string(nft.ABI)))
	if err != nil {
		return err
	}

	ctx = context.WithValue(ctx, orbitContext.KeyTxHash, vLog.TxHash)

	for k, v := range hashMap {
		if k == vLog.Topics[0].Hex() {
			parser := logParserMap[v]
			handler := logHandlerMap[v]
			switch v {
			case LogApprovalName:
				evt, err := parser.(func(*abi.ABI, *types.Log) (*LogApproval, error))(&contractAbi, vLog)
				if err != nil {
					return err
				}
				return handler.(func(context.Context, *LogApproval) error)(ctx, evt)
			case LogApprovalForAllName:
				evt, err := parser.(func(*abi.ABI, *types.Log) (*LogApprovalForAll, error))(&contractAbi, vLog)
				if err != nil {
					return err
				}
				return handler.(func(context.Context, *LogApprovalForAll) error)(ctx, evt)
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
			case LogTransferName:
				evt, err := parser.(func(*abi.ABI, *types.Log) (*LogTransfer, error))(&contractAbi, vLog)
				if err != nil {
					return err
				}
				return handler.(func(context.Context, *LogTransfer) error)(ctx, evt)
			default:
				return errors.New("nft: event type is not supported")
			}
		}
	}
	return errors.New("nft: topic is not supported")
}

var logParserMap map[string]interface{}
var logHandlerMap map[string]interface{}
var hashMap map[string]string

func init() {
	logParserMap = make(map[string]interface{})
	logHandlerMap = make(map[string]interface{})
	hashMap = make(map[string]string)
	logParserMap[LogApprovalName] = ParseLogApproval
	logParserMap[LogApprovalForAllName] = ParseLogApprovalForAll
	logParserMap[LogRoleAdminChangedName] = ParseLogRoleAdminChanged
	logParserMap[LogRoleGrantedName] = ParseLogRoleGranted
	logParserMap[LogRoleRevokedName] = ParseLogRoleRevoked
	logParserMap[LogTransferName] = ParseLogTransfer

	logHandlerMap[LogApprovalName] = HandleLogApproval
	logHandlerMap[LogApprovalForAllName] = HandleLogApprovalForAll
	logHandlerMap[LogRoleAdminChangedName] = HandleLogRoleAdminChanged
	logHandlerMap[LogRoleGrantedName] = HandleLogRoleGranted
	logHandlerMap[LogRoleRevokedName] = HandleLogRoleRevoked
	logHandlerMap[LogTransferName] = HandleLogTransfer

	logApprovalSigHash := crypto.Keccak256Hash([]byte(LogApprovalSig))
	hashMap[logApprovalSigHash.Hex()] = LogApprovalName

	logApprovalForAllSigHash := crypto.Keccak256Hash([]byte(LogApprovalForAllSig))
	hashMap[logApprovalForAllSigHash.Hex()] = LogApprovalForAllName

	logRoleAdminChangedSigHash := crypto.Keccak256Hash([]byte(LogRoleAdminChangedSig))
	hashMap[logRoleAdminChangedSigHash.Hex()] = LogRoleAdminChangedName

	logRoleGrantedSigHash := crypto.Keccak256Hash([]byte(LogRoleGrantedSig))
	hashMap[logRoleGrantedSigHash.Hex()] = LogRoleGrantedName

	logRoleRevokedSigHash := crypto.Keccak256Hash([]byte(LogRoleRevokedSig))
	hashMap[logRoleRevokedSigHash.Hex()] = LogRoleRevokedName

	logTransferSigHash := crypto.Keccak256Hash([]byte(LogTransferSig))
	hashMap[logTransferSigHash.Hex()] = LogTransferName

	event.Register(&parserHandler{
		name: nft.Name,
	})
}
