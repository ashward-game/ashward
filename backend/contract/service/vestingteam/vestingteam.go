package vestingteam

import (
	"math/big"
	"orbit_nft/contract/abi/vestingteam"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type VestingTeam struct {
	VestingTeamCaller
	VestingTeamTransactor
}

type VestingTeamCaller struct {
	contract *bind.BoundContract
}

type VestingTeamTransactor struct {
	contract *bind.BoundContract
}

func NewVestingTeam(address common.Address, backend bind.ContractBackend) (*VestingTeam, error) {
	contract, err := bindVestingTeam(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VestingTeam{VestingTeamCaller: VestingTeamCaller{contract: contract}, VestingTeamTransactor: VestingTeamTransactor{contract: contract}}, nil
}

func bindVestingTeam(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(vestingteam.ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_VestingTeam *VestingTeamTransactor) AddBeneficiaries(opts *bind.TransactOpts, beneficiaries []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _VestingTeam.contract.Transact(opts, "addBeneficiaries", beneficiaries, amounts)
}
