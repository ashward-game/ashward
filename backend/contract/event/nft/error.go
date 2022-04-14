package nft

import (
	"context"
	"orbit_nft/contract/abi/nft"
	orbitContext "orbit_nft/contract/context"
	eventerror "orbit_nft/contract/event/error"

	"github.com/ethereum/go-ethereum/common"
)

func Err(ctx context.Context, message error) error {
	tx := ctx.Value(orbitContext.KeyTxHash).(common.Hash)
	return eventerror.New(nft.Name, tx.String(), message.Error())
}
