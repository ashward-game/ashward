package vestingadvisory

import (
	"math/big"
	"orbit_nft/contract/abi/vestingadvisory"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type VestingAdvisory struct {
	VestingAdvisoryCaller
	VestingAdvisoryTransactor
}

type VestingAdvisoryCaller struct {
	contract *bind.BoundContract
}

type VestingAdvisoryTransactor struct {
	contract *bind.BoundContract
}

func NewVestingAdvisory(address common.Address, backend bind.ContractBackend) (*VestingAdvisory, error) {
	contract, err := bindVestingAdvisory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VestingAdvisory{VestingAdvisoryCaller: VestingAdvisoryCaller{contract: contract}, VestingAdvisoryTransactor: VestingAdvisoryTransactor{contract: contract}}, nil
}

func bindVestingAdvisory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(vestingadvisory.ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_VestingAdvisory *VestingAdvisoryTransactor) AddBeneficiaries(opts *bind.TransactOpts, beneficiaries []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _VestingAdvisory.contract.Transact(opts, "addBeneficiaries", beneficiaries, amounts)
}
