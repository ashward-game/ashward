package vestingido

import (
	"math/big"
	"orbit_nft/contract/abi/vestingido"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type VestingIDO struct {
	VestingIDOCaller
	VestingIDOTransactor
}

type VestingIDOCaller struct {
	contract *bind.BoundContract
}

type VestingIDOTransactor struct {
	contract *bind.BoundContract
}

func NewVestingIDO(address common.Address, backend bind.ContractBackend) (*VestingIDO, error) {
	contract, err := bindVestingIDO(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VestingIDO{VestingIDOCaller: VestingIDOCaller{contract: contract}, VestingIDOTransactor: VestingIDOTransactor{contract: contract}}, nil
}

func bindVestingIDO(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(vestingido.ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_VestingIDO *VestingIDOTransactor) AddBeneficiaries(opts *bind.TransactOpts, beneficiaries []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _VestingIDO.contract.Transact(opts, "addBeneficiaries", beneficiaries, amounts)
}
