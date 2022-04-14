package vestingstaking

import (
	"math/big"
	"orbit_nft/contract/abi/vestingstaking"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type VestingStaking struct {
	VestingStakingCaller
	VestingStakingTransactor
}

type VestingStakingCaller struct {
	contract *bind.BoundContract
}

type VestingStakingTransactor struct {
	contract *bind.BoundContract
}

func NewVestingStaking(address common.Address, backend bind.ContractBackend) (*VestingStaking, error) {
	contract, err := bindVestingStaking(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VestingStaking{VestingStakingCaller: VestingStakingCaller{contract: contract}, VestingStakingTransactor: VestingStakingTransactor{contract: contract}}, nil
}

func bindVestingStaking(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(vestingstaking.ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_VestingStaking *VestingStakingTransactor) AddBeneficiaries(opts *bind.TransactOpts, beneficiaries []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _VestingStaking.contract.Transact(opts, "addBeneficiaries", beneficiaries, amounts)
}
