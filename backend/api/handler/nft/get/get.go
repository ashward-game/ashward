package get

import (
	"net/http"
	"orbit_nft/api/util"
	"orbit_nft/api/util/pagination"
	"orbit_nft/db/service"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type handler struct {
	service *service.NftService
}

func NewHandler(service *service.NftService) *handler {
	return &handler{service: service}
}

func (h *handler) GetTokenOfAddress(ctx *gin.Context) {
	var validate InputListNftOfAddress
	if err := ctx.ShouldBindUri(&validate); err != nil {
		util.ValidatorErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	if err := ctx.ShouldBind(&validate); err != nil {
		util.ValidatorErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	filter := service.TokenFilter{
		Type:         ctx.Query("type"),
		Class:        ctx.Query("class"),
		Rarity:       ctx.Query("rarity"),
		Search:       strings.TrimSpace(ctx.Query("search")),
		OrderByPrice: ctx.DefaultQuery("order_by_price", pagination.DefaultSort),
	}
	address := validate.Address

	pg, err := pagination.PaginationFromRequest(ctx)
	if err != nil {
		log.Error().Err(err).Str("query", ctx.Request.URL.RequestURI()).Msg("pagination.PaginationFromRequest")
		util.APIResponseError(ctx, http.StatusInternalServerError)
		return
	}

	total, result, err := h.service.GetTokensOfAddressWithFilter(pg, address, filter)
	if err != nil {
		log.Error().Err(err).Str("query", ctx.Request.URL.RequestURI()).Msg("h.service.GetTokensOfAddressWithFilter")
		util.APIResponseError(ctx, http.StatusInternalServerError)
		return
	}

	util.APIResponse(ctx, http.StatusOK, gin.H{
		"total": total,
		"data":  result,
	})
}
