package vestingmarketing

import (
	"math/big"
	"orbit_nft/contract/abi/vestingmarketing"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type VestingMarketing struct {
	VestingMarketingCaller
	VestingMarketingTransactor
}

type VestingMarketingCaller struct {
	contract *bind.BoundContract
}

type VestingMarketingTransactor struct {
	contract *bind.BoundContract
}

func NewVestingMarketing(address common.Address, backend bind.ContractBackend) (*VestingMarketing, error) {
	contract, err := bindVestingMarketing(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VestingMarketing{VestingMarketingCaller: VestingMarketingCaller{contract: contract}, VestingMarketingTransactor: VestingMarketingTransactor{contract: contract}}, nil
}

func bindVestingMarketing(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(vestingmarketing.ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_VestingMarketing *VestingMarketingTransactor) AddBeneficiaries(opts *bind.TransactOpts, beneficiaries []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _VestingMarketing.contract.Transact(opts, "addBeneficiaries", beneficiaries, amounts)
}
