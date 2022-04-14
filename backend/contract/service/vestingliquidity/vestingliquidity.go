package vestingliquidity

import (
	"math/big"
	"orbit_nft/contract/abi/vestingliquidity"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type VestingLiquidity struct {
	VestingLiquidityCaller
	VestingLiquidityTransactor
}

type VestingLiquidityCaller struct {
	contract *bind.BoundContract
}

type VestingLiquidityTransactor struct {
	contract *bind.BoundContract
}

func NewVestingLiquidity(address common.Address, backend bind.ContractBackend) (*VestingLiquidity, error) {
	contract, err := bindVestingLiquidity(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VestingLiquidity{VestingLiquidityCaller: VestingLiquidityCaller{contract: contract}, VestingLiquidityTransactor: VestingLiquidityTransactor{contract: contract}}, nil
}

func bindVestingLiquidity(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(vestingliquidity.ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_VestingLiquidity *VestingLiquidityTransactor) AddBeneficiaries(opts *bind.TransactOpts, beneficiaries []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _VestingLiquidity.contract.Transact(opts, "addBeneficiaries", beneficiaries, amounts)
}
