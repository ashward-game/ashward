package marketplace

import (
	"context"
	"errors"
	"math/big"
	"orbit_nft/contract/abi/marketplace"
	orbitContext "orbit_nft/contract/context"
	"orbit_nft/contract/event"
	"orbit_nft/util"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

const LogOfferCanceledName string = "OfferCanceled"
const LogOfferCanceledSig string = "OfferCanceled(address,uint256)"

type LogOfferCanceled struct {
	Seller  common.Address
	TokenId *big.Int
}

func ParseLogOfferCanceled(contractAbi *abi.ABI, vLog *types.Log) (*LogOfferCanceled, error) {
	var evt LogOfferCanceled
	if len(vLog.Data) > 0 {
		err := contractAbi.UnpackIntoInterface(&evt, LogOfferCanceledName, vLog.Data)
		if err != nil {
			return nil, err
		}
	}

	evt.Seller = common.HexToAddress(vLog.Topics[1].Hex())

	evt.TokenId = util.HexToBigInt(vLog.Topics[2].Hex())

	return &evt, nil
}

const LogOfferCreatedName string = "OfferCreated"
const LogOfferCreatedSig string = "OfferCreated(address,uint256,uint256)"

type LogOfferCreated struct {
	Seller  common.Address
	TokenId *big.Int
	Price   *big.Int
}

func ParseLogOfferCreated(contractAbi *abi.ABI, vLog *types.Log) (*LogOfferCreated, error) {
	var evt LogOfferCreated
	if len(vLog.Data) > 0 {
		err := contractAbi.UnpackIntoInterface(&evt, LogOfferCreatedName, vLog.Data)
		if err != nil {
			return nil, err
		}
	}

	evt.Seller = common.HexToAddress(vLog.Topics[1].Hex())

	evt.TokenId = util.HexToBigInt(vLog.Topics[2].Hex())

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

const LogTokenPurchasedName string = "TokenPurchased"
const LogTokenPurchasedSig string = "TokenPurchased(address,uint256)"

type LogTokenPurchased struct {
	Buyer   common.Address
	TokenId *big.Int
}

func ParseLogTokenPurchased(contractAbi *abi.ABI, vLog *types.Log) (*LogTokenPurchased, error) {
	var evt LogTokenPurchased
	if len(vLog.Data) > 0 {
		err := contractAbi.UnpackIntoInterface(&evt, LogTokenPurchasedName, vLog.Data)
		if err != nil {
			return nil, err
		}
	}

	evt.Buyer = common.HexToAddress(vLog.Topics[1].Hex())

	evt.TokenId = util.HexToBigInt(vLog.Topics[2].Hex())

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

type parserHandler struct {
	name string
}

var _ event.ParserHandler = (*parserHandler)(nil)

func (p *parserHandler) Name() string {
	return p.name
}

func (p *parserHandler) ParseHandle(ctx context.Context, vLog *types.Log) error {
	contractAbi, err := abi.JSON(strings.NewReader(string(marketplace.ABI)))
	if err != nil {
		return err
	}

	ctx = context.WithValue(ctx, orbitContext.KeyTxHash, vLog.TxHash)

	for k, v := range hashMap {
		if k == vLog.Topics[0].Hex() {
			parser := logParserMap[v]
			handler := logHandlerMap[v]
			switch v {
			case LogOfferCanceledName:
				evt, err := parser.(func(*abi.ABI, *types.Log) (*LogOfferCanceled, error))(&contractAbi, vLog)
				if err != nil {
					return err
				}
				return handler.(func(context.Context, *LogOfferCanceled) error)(ctx, evt)
			case LogOfferCreatedName:
				evt, err := parser.(func(*abi.ABI, *types.Log) (*LogOfferCreated, error))(&contractAbi, vLog)
				if err != nil {
					return err
				}
				return handler.(func(context.Context, *LogOfferCreated) error)(ctx, evt)
			case LogOwnershipTransferredName:
				evt, err := parser.(func(*abi.ABI, *types.Log) (*LogOwnershipTransferred, error))(&contractAbi, vLog)
				if err != nil {
					return err
				}
				return handler.(func(context.Context, *LogOwnershipTransferred) error)(ctx, evt)
			case LogTokenPurchasedName:
				evt, err := parser.(func(*abi.ABI, *types.Log) (*LogTokenPurchased, error))(&contractAbi, vLog)
				if err != nil {
					return err
				}
				return handler.(func(context.Context, *LogTokenPurchased) error)(ctx, evt)
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
			default:
				return errors.New("marketplace: event type is not supported")
			}
		}
	}
	return errors.New("marketplace: topic is not supported")
}

var logParserMap map[string]interface{}
var logHandlerMap map[string]interface{}
var hashMap map[string]string

func init() {
	logParserMap = make(map[string]interface{})
	logHandlerMap = make(map[string]interface{})
	hashMap = make(map[string]string)
	logParserMap[LogOfferCanceledName] = ParseLogOfferCanceled
	logParserMap[LogOfferCreatedName] = ParseLogOfferCreated
	logParserMap[LogOwnershipTransferredName] = ParseLogOwnershipTransferred
	logParserMap[LogTokenPurchasedName] = ParseLogTokenPurchased
	logParserMap[LogTokensApprovedName] = ParseLogTokensApproved
	logParserMap[LogTokensReceivedName] = ParseLogTokensReceived

	logHandlerMap[LogOfferCanceledName] = HandleLogOfferCanceled
	logHandlerMap[LogOfferCreatedName] = HandleLogOfferCreated
	logHandlerMap[LogOwnershipTransferredName] = HandleLogOwnershipTransferred
	logHandlerMap[LogTokenPurchasedName] = HandleLogTokenPurchased
	logHandlerMap[LogTokensApprovedName] = HandleLogTokensApproved
	logHandlerMap[LogTokensReceivedName] = HandleLogTokensReceived

	logOfferCanceledSigHash := crypto.Keccak256Hash([]byte(LogOfferCanceledSig))
	hashMap[logOfferCanceledSigHash.Hex()] = LogOfferCanceledName

	logOfferCreatedSigHash := crypto.Keccak256Hash([]byte(LogOfferCreatedSig))
	hashMap[logOfferCreatedSigHash.Hex()] = LogOfferCreatedName

	logOwnershipTransferredSigHash := crypto.Keccak256Hash([]byte(LogOwnershipTransferredSig))
	hashMap[logOwnershipTransferredSigHash.Hex()] = LogOwnershipTransferredName

	logTokenPurchasedSigHash := crypto.Keccak256Hash([]byte(LogTokenPurchasedSig))
	hashMap[logTokenPurchasedSigHash.Hex()] = LogTokenPurchasedName

	logTokensApprovedSigHash := crypto.Keccak256Hash([]byte(LogTokensApprovedSig))
	hashMap[logTokensApprovedSigHash.Hex()] = LogTokensApprovedName

	logTokensReceivedSigHash := crypto.Keccak256Hash([]byte(LogTokensReceivedSig))
	hashMap[logTokensReceivedSigHash.Hex()] = LogTokensReceivedName

	event.Register(&parserHandler{
		name: marketplace.Name,
	})
}
