package marketplace

import (
	"context"
	orbitContext "orbit_nft/contract/context"
	"orbit_nft/db"
	"orbit_nft/db/repository"
	"orbit_nft/db/service"
)

func newMarketplaceService(ctx context.Context) *service.MarketplaceService {
	sqlDB := ctx.Value(orbitContext.KeyDB).(*db.Database)
	marketRepo := repository.NewMarketplaceRepository(sqlDB)
	nftRepo := repository.NewNftRepository(sqlDB)

	return service.NewMarketplaceService(marketRepo, nftRepo)
}

func HandleLogOfferCanceled(ctx context.Context, evt *LogOfferCanceled) error {
	marketService := newMarketplaceService(ctx)

	tokenId := uint(evt.TokenId.Uint64())
	seller := evt.Seller.String()

	err := marketService.CancelOffer(tokenId, seller)
	if err != nil {
		return Err(ctx, err)
	}
	return nil
}

func HandleLogOfferCreated(ctx context.Context, evt *LogOfferCreated) error {
	marketService := newMarketplaceService(ctx)

	tokenId := uint(evt.TokenId.Uint64())
	seller := evt.Seller.String()
	price := evt.Price

	err := marketService.OpenOffer(tokenId, seller, price)
	if err != nil {
		return Err(ctx, err)
	}
	return nil
}

func HandleLogOwnershipTransferred(ctx context.Context, evt *LogOwnershipTransferred) error {
	return nil
}

func HandleLogTokenPurchased(ctx context.Context, evt *LogTokenPurchased) error {
	marketService := newMarketplaceService(ctx)

	tokenId := uint(evt.TokenId.Uint64())
	buyer := evt.Buyer.String()

	err := marketService.Purchase(tokenId, buyer)
	if err != nil {
		return Err(ctx, err)
	}
	return nil
}

func HandleLogTokensApproved(ctx context.Context, evt *LogTokensApproved) error {
	return nil
}

func HandleLogTokensReceived(ctx context.Context, evt *LogTokensReceived) error {
	return nil
}
