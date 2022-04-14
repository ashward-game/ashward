package nft

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"orbit_nft/contract/abi/nft"

	"github.com/ethereum/go-ethereum/core/types"
)

type NFToken struct {
	NftCaller
	NftTransactor
}

type NftCaller struct {
	contract *bind.BoundContract
}

type NftTransactor struct {
	contract *bind.BoundContract
}

type NftSession struct {
	Contract     *NFToken          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

func NewNFToken(address common.Address, backend bind.ContractBackend) (*NFToken, error) {
	contract, err := bindNft(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NFToken{NftCaller: NftCaller{contract: contract}, NftTransactor: NftTransactor{contract: contract}}, nil
}

func bindNft(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(nft.ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Mint Solidity: function mint(string uri)
func (_Nft *NftTransactor) Mint(opts *bind.TransactOpts, uri string) (*types.Transaction, error) {
	return _Nft.contract.Transact(opts, "mint", uri)
}

// MintAndTransfer Solidity: function mintAndTransfer(string memory uri, address owner)
func (_Nft *NftTransactor) MintAndTransfer(opts *bind.TransactOpts, uri string, owner common.Address) (*types.Transaction, error) {
	return _Nft.contract.Transact(opts, "mintAndTransfer", uri, owner)
}

// TokenURI Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Nft *NftCaller) TokenURI(tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Nft.contract.Call(nil, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err
}

func (_Nft *NftCaller) OwnerOf(tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Nft.contract.Call(nil, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_Nft *NftTransactor) SetupMinter(opts *bind.TransactOpts, minter common.Address) (*types.Transaction, error) {
	return _Nft.contract.Transact(opts, "setupMinter", minter)
}
