package route

import (
	"orbit_nft/api/handler/nft/get"
	"orbit_nft/db"
	"orbit_nft/db/repository"
	"orbit_nft/db/service"

	"github.com/gin-gonic/gin"
)

func nftRoutes(db *db.Database, route *gin.RouterGroup) {
	groupRoute := route.Group("/nft")

	repo := repository.NewNftRepository(db)
	service := service.NewNftService(repo)

	getHandler := get.NewHandler(service)
	groupRoute.GET("/:address/lists", getHandler.GetTokenOfAddress)
}
