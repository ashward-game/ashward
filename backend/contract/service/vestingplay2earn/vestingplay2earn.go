package vestingplay2earn

import (
	"math/big"
	"orbit_nft/contract/abi/vestingplay2earn"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type VestingPlay2Earn struct {
	VestingPlay2EarnCaller
	VestingPlay2EarnTransactor
}

type VestingPlay2EarnCaller struct {
	contract *bind.BoundContract
}

type VestingPlay2EarnTransactor struct {
	contract *bind.BoundContract
}

func NewVestingPlay2Earn(address common.Address, backend bind.ContractBackend) (*VestingPlay2Earn, error) {
	contract, err := bindVestingPlay2Earn(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VestingPlay2Earn{VestingPlay2EarnCaller: VestingPlay2EarnCaller{contract: contract}, VestingPlay2EarnTransactor: VestingPlay2EarnTransactor{contract: contract}}, nil
}

func bindVestingPlay2Earn(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(vestingplay2earn.ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_VestingPlay2Earn *VestingPlay2EarnTransactor) AddBeneficiaries(opts *bind.TransactOpts, beneficiaries []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _VestingPlay2Earn.contract.Transact(opts, "addBeneficiaries", beneficiaries, amounts)
}
