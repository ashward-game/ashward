package nft

import (
	"context"
	"orbit_nft/constant"
	orbitContext "orbit_nft/contract/context"
	"orbit_nft/contract/event/whitelist"
	"orbit_nft/db"
	"orbit_nft/db/repository"
	"orbit_nft/db/service"
)

func HandleLogApproval(ctx context.Context, evt *LogApproval) error {
	return nil
}

func HandleLogApprovalForAll(ctx context.Context, evt *LogApprovalForAll) error {
	return nil
}

func HandleLogRoleAdminChanged(ctx context.Context, evt *LogRoleAdminChanged) error {
	return nil
}

func HandleLogRoleGranted(ctx context.Context, evt *LogRoleGranted) error {
	return nil
}

func HandleLogRoleRevoked(ctx context.Context, evt *LogRoleRevoked) error {
	return nil
}

func HandleLogTransfer(ctx context.Context, evt *LogTransfer) error {
	sqlDB := ctx.Value(orbitContext.KeyDB).(*db.Database)
	repo := repository.NewNftRepository(sqlDB)
	service := service.NewNftService(repo)

	isMintEvent := evt.From.String() == constant.AddressZero().String()
	if isMintEvent {
		metadataURI, err := getTokenURI(ctx, evt.TokenId)
		if err != nil {
			return Err(ctx, err)
		}

		rawStringMetadata, err := getTokenMetadata(ctx, metadataURI)
		if err != nil {
			return Err(ctx, err)
		}

		err = service.CreateNotForSaleToken(
			uint(evt.TokenId.Uint64()),
			evt.To.String(),
			metadataURI, rawStringMetadata)
		if err != nil {
			return Err(ctx, err)
		}
		return nil
	}

	if whitelist.IsTokenTransferTo(evt.To) {
		// this is a marketplace open offer event
		return nil
	}

	var from string
	if whitelist.IsTokenTransferFrom(evt.From) {
		// this is a marketplace cancel offer or purchase event
		from = "" // ignore the sender, as it is not the real owner
	} else {
		from = evt.From.String()
	}

	err := service.TransferNFT(
		uint(evt.TokenId.Uint64()),
		from,
		evt.To.String())
	if err != nil {
		return Err(ctx, err)
	}
	return nil
}
