package vestingprivate

import (
	"math/big"
	"orbit_nft/contract/abi/vestingprivate"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type VestingPrivate struct {
	VestingPrivateCaller
	VestingPrivateTransactor
}

type VestingPrivateCaller struct {
	contract *bind.BoundContract
}

type VestingPrivateTransactor struct {
	contract *bind.BoundContract
}

func NewVestingPrivate(address common.Address, backend bind.ContractBackend) (*VestingPrivate, error) {
	contract, err := bindVestingPrivate(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VestingPrivate{VestingPrivateCaller: VestingPrivateCaller{contract: contract}, VestingPrivateTransactor: VestingPrivateTransactor{contract: contract}}, nil
}

func bindVestingPrivate(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(vestingprivate.ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_VestingPrivate *VestingPrivateTransactor) AddBeneficiaries(opts *bind.TransactOpts, beneficiaries []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _VestingPrivate.contract.Transact(opts, "addBeneficiaries", beneficiaries, amounts)
}
