package marketplace

import (
	"context"
	"orbit_nft/contract/abi/marketplace"
	orbitContext "orbit_nft/contract/context"
	eventerror "orbit_nft/contract/event/error"
)

func Err(ctx context.Context, message error) error {
	tx := ctx.Value(orbitContext.KeyTxHash).(string)
	return eventerror.New(marketplace.Name, tx, message.Error())
}
