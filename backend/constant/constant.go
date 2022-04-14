package constant

import (
	"github.com/ethereum/go-ethereum/common"
)

const EthDerivationPath = "m/44'/60'/0'/0/0"

func AddressZero() common.Address {
	return common.HexToAddress("0x0000000000000000000000000000000000000000")
}
