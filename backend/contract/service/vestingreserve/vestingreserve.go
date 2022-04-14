package vestingreserve

import (
	"math/big"
	"orbit_nft/contract/abi/vestingreserve"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type VestingReserve struct {
	VestingReserveCaller
	VestingReserveTransactor
}

type VestingReserveCaller struct {
	contract *bind.BoundContract
}

type VestingReserveTransactor struct {
	contract *bind.BoundContract
}

func NewVestingReserve(address common.Address, backend bind.ContractBackend) (*VestingReserve, error) {
	contract, err := bindVestingReserve(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VestingReserve{VestingReserveCaller: VestingReserveCaller{contract: contract}, VestingReserveTransactor: VestingReserveTransactor{contract: contract}}, nil
}

func bindVestingReserve(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(vestingreserve.ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_VestingReserve *VestingReserveTransactor) AddBeneficiaries(opts *bind.TransactOpts, beneficiaries []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _VestingReserve.contract.Transact(opts, "addBeneficiaries", beneficiaries, amounts)
}
