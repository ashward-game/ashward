package openboxgenesis

import (
	"orbit_nft/contract/abi/openboxgenesis"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type Openboxgenesis struct {
	OpenboxgenesisCaller
	OpenboxgenesisTransactor
}
type OpenboxgenesisCaller struct {
	contract *bind.BoundContract
}

type OpenboxgenesisTransactor struct {
	contract *bind.BoundContract
}

func NewOpenboxgenesis(address common.Address, backend bind.ContractBackend) (*Openboxgenesis, error) {
	contract, err := bindOpenboxgenesis(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Openboxgenesis{OpenboxgenesisCaller: OpenboxgenesisCaller{contract: contract}, OpenboxgenesisTransactor: OpenboxgenesisTransactor{contract: contract}}, nil
}

func bindOpenboxgenesis(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(openboxgenesis.ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_Openboxgenesis *OpenboxgenesisTransactor) OpenBox(opts *bind.TransactOpts, sHash [32]byte, isEmpty bool, tokenURI string) (*types.Transaction, error) {
	return _Openboxgenesis.contract.Transact(opts, "openBox", sHash, isEmpty, tokenURI)
}

func (_Openboxgenesis *OpenboxgenesisTransactor) AddSubscribers(opts *bind.TransactOpts, users []common.Address) (*types.Transaction, error) {
	return _Openboxgenesis.contract.Transact(opts, "addSubscribers", users)
}

func (_Openboxgenesis *OpenboxgenesisTransactor) BuyBox(opts *bind.TransactOpts, grade uint8, serverHash [32]byte, serverSig []byte, clientRandom [32]byte) (*types.Transaction, error) {
	return _Openboxgenesis.contract.Transact(opts, "buyBox", grade, serverHash, serverSig, clientRandom)
}

func (_Openboxgenesis *OpenboxgenesisTransactor) SetupOpener(opts *bind.TransactOpts, opener common.Address) (*types.Transaction, error) {
	return _Openboxgenesis.contract.Transact(opts, "setupOpener", opener)
}
