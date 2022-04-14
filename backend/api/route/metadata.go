package route

import (
	"orbit_nft/api/handler/metadata/get"
	"orbit_nft/db"
	"orbit_nft/db/repository"
	"orbit_nft/db/service"

	"github.com/gin-gonic/gin"
)

func metadataRoutes(db *db.Database, route *gin.RouterGroup) {
	groupRoute := route.Group("/metadata")

	repo := repository.NewMetadataRepository(db)
	svc := service.NewMetadataService(repo)

	getHandler := get.NewHandler(svc)
	groupRoute.GET("/newblock", getHandler.Get)
}
