// This module provides functions to check if an address is whitelisted.
// If an address is whitelisted, a related event might be ignored during the
// handling process.
// An example is as follows.
// When an NFT is transfer to our marketplace (by opening an offer), the owner
// is not changed (to the marketplace's address) in the database.
// This is designed on purpose, to optimize SQL queries (that we do not need
// to make a join queries for many operations).
// Consequently, this event should be ignored.
package whitelist

import (
	"errors"
	"fmt"
	"orbit_nft/contract/abi/marketplace"
	"orbit_nft/util"

	"github.com/ethereum/go-ethereum/common"
)

var tokenTransferFrom map[string]bool
var tokenTransferTo map[string]bool

func Setup(addressFile string) error {
	mkpAddress, err := util.GetContractAddress(addressFile, marketplace.Name)
	if err != nil {
		s := fmt.Sprintf("whitelist: error while reading contract's address: %s", err)
		return errors.New(s)
	}

	tokenTransferFrom = map[string]bool{
		mkpAddress: true,
	}

	tokenTransferTo = map[string]bool{
		mkpAddress: true,
	}
	return nil
}

func IsTokenTransferFrom(address common.Address) bool {
	return tokenTransferFrom[address.String()]
}

func IsTokenTransferTo(address common.Address) bool {
	return tokenTransferTo[address.String()]
}
