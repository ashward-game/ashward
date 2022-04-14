package get

import (
	"orbit_nft/api/util"
	"orbit_nft/db/service"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service *service.MetadataService
}

func NewHandler(service *service.MetadataService) *handler {
	return &handler{service: service}
}

func (h *handler) Get(ctx *gin.Context) {
	// validate request input
	// call to service
	// return response

	util.APIResponse(ctx, 200, "pong")
}
