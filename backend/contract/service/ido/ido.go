package ido

import (
	"strings"

	"orbit_nft/contract/abi/ido"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type Ido struct {
	IdoCaller
	IdoTransactor
}

type IdoCaller struct {
	contract *bind.BoundContract
}

type IdoTransactor struct {
	contract *bind.BoundContract
}

func NewIdo(address common.Address, backend bind.ContractBackend) (*Ido, error) {
	contract, err := bindIdo(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ido{IdoCaller: IdoCaller{contract: contract}, IdoTransactor: IdoTransactor{contract: contract}}, nil
}

func bindIdo(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ido.ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_Ido *IdoTransactor) PublicSale(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ido.contract.Transact(opts, "publicSale")
}

func (_Ido *IdoTransactor) AddSubscribers(opts *bind.TransactOpts, users []common.Address) (*types.Transaction, error) {
	return _Ido.contract.Transact(opts, "addSubscribers", users)
}

func (_Ido *IdoTransactor) Stop(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ido.contract.Transact(opts, "stop")
}

func (_Ido *IdoTransactor) CollectBUSD(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ido.contract.Transact(opts, "collectBUSD")
}
