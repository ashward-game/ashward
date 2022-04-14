package token

import (
	"orbit_nft/contract/abi/token"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/core/types"
)

type Token struct {
	TokenCaller
	TokenTransactor
}

type TokenCaller struct {
	contract *bind.BoundContract
}

type TokenTransactor struct {
	contract *bind.BoundContract
}

type TokenSession struct {
	Contract     *Token            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

func NewToken(address common.Address, backend bind.ContractBackend) (*Token, error) {
	contract, err := bindToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Token{TokenCaller: TokenCaller{contract: contract}, TokenTransactor: TokenTransactor{contract: contract}}, nil
}

func bindToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(token.ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_token *TokenTransactor) AddSellingAddress(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _token.contract.Transact(opts, "addSellingAddress", addr)
}

func (_token *TokenTransactor) AddNoTaxAddress(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _token.contract.Transact(opts, "addNoTaxAddress", addr)
}

func (_token *TokenTransactor) RemoveNoTaxAddress(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _token.contract.Transact(opts, "removeNoTaxAddress", addr)
}
