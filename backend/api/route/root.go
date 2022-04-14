package route

import (
	"orbit_nft/db"

	"github.com/gin-gonic/gin"
)

func Init(db *db.Database, route *gin.Engine) {
	groupRoute := route.Group("/api/v1")

	// define all routes
	metadataRoutes(db, groupRoute)
	marketplaceRoutes(db, groupRoute)
	nftRoutes(db, groupRoute)
	makerandRoutes(db, groupRoute)
}
