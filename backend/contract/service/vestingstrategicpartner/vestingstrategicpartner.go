package vestingstrategicpartner

import (
	"math/big"
	"orbit_nft/contract/abi/vestingstrategicpartner"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type VestingStrategicPartner struct {
	VestingStrategicPartnerCaller
	VestingStrategicPartnerTransactor
}

type VestingStrategicPartnerCaller struct {
	contract *bind.BoundContract
}

type VestingStrategicPartnerTransactor struct {
	contract *bind.BoundContract
}

func NewVestingStrategicPartner(address common.Address, backend bind.ContractBackend) (*VestingStrategicPartner, error) {
	contract, err := bindVestingStrategicPartner(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VestingStrategicPartner{VestingStrategicPartnerCaller: VestingStrategicPartnerCaller{contract: contract}, VestingStrategicPartnerTransactor: VestingStrategicPartnerTransactor{contract: contract}}, nil
}

func bindVestingStrategicPartner(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(vestingstrategicpartner.ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_VestingStrategicPartner *VestingStrategicPartnerTransactor) AddBeneficiaries(opts *bind.TransactOpts, beneficiaries []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _VestingStrategicPartner.contract.Transact(opts, "addBeneficiaries", beneficiaries, amounts)
}
