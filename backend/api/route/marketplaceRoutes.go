package route

import (
	"orbit_nft/api/handler/marketplace/get"
	"orbit_nft/db"
	"orbit_nft/db/repository"
	"orbit_nft/db/service"

	"github.com/gin-gonic/gin"
)

func marketplaceRoutes(db *db.Database, route *gin.RouterGroup) {
	groupRoute := route.Group("/marketplace")

	repo := repository.NewNftRepository(db)
	repoMarket := repository.NewMarketplaceRepository(db)
	nftService := service.NewNftService(repo)
	marketService := service.NewMarketplaceService(repoMarket, repo)

	getHandler := get.NewHandler(nftService, marketService)

	groupRoute.GET("/", getHandler.Get)
	groupRoute.GET("/:address/history", getHandler.GetTradingHistory)
	groupRoute.GET("/nft/:id", getHandler.ShowNft)
}
