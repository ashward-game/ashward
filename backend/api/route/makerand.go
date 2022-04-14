package route

import (
	"orbit_nft/api/handler/makerand/get"
	"orbit_nft/db"
	"orbit_nft/db/repository"
	"orbit_nft/db/service"

	"github.com/gin-gonic/gin"
)

func makerandRoutes(db *db.Database, route *gin.RouterGroup) {
	groupRoute := route.Group("/makerand")

	repoMakeRand := repository.NewMakeRandRepository(db)
	makeRandService := service.NewMakeRandService(repoMakeRand)
	getHandler := get.NewHandler(makeRandService)
	groupRoute.GET("/commit", getHandler.Commit)
}
