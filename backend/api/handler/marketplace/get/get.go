package get

import (
	"net/http"
	"orbit_nft/api/util"
	"orbit_nft/api/util/pagination"
	"orbit_nft/db/service"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type handler struct {
	nftService    *service.NftService
	marketService *service.MarketplaceService
}

func NewHandler(nftService *service.NftService, marketService *service.MarketplaceService) *handler {
	return &handler{
		nftService:    nftService,
		marketService: marketService,
	}
}

func (h *handler) Get(ctx *gin.Context) {
	var input InputList
	if err := ctx.ShouldBind(&input); err != nil {
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

	pg, err := pagination.PaginationFromRequest(ctx)
	if err != nil {
		log.Error().Err(err).Str("query", ctx.Request.URL.RequestURI()).Msg("pagination.PaginationFromRequest")
		util.APIResponseError(ctx, http.StatusInternalServerError)
		return
	}

	total, result, err := h.nftService.GetOnSaleTokensWithFilter(pg, filter)
	if err != nil {
		log.Error().Err(err).Str("query", ctx.Request.URL.RequestURI()).Msg("h.nftService.GetOnSaleTokensWithFilter")
		util.APIResponseError(ctx, http.StatusInternalServerError)
		return
	}

	util.APIResponse(ctx, http.StatusOK, gin.H{
		"total": total,
		"data":  result,
	})
}

func (h *handler) GetTradingHistory(ctx *gin.Context) {
	var input InputTradingHistory
	if err := ctx.ShouldBindUri(&input); err != nil {
		util.ValidatorErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	if err := ctx.ShouldBind(&input); err != nil {
		util.ValidatorErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	pg, err := pagination.PaginationFromRequest(ctx)
	if err != nil {
		log.Error().Err(err).Str("query", ctx.Request.URL.RequestURI()).Msg("pagination.PaginationFromRequest")
		util.APIResponseError(ctx, http.StatusInternalServerError)
		return
	}

	total, result, err := h.marketService.GetTradingOfAddressWithPg(pg, input.Address)
	if err != nil {
		log.Error().Err(err).Str("query", ctx.Request.URL.RequestURI()).Msg("h.service.GetTradingOfAddressWithPg")
		util.APIResponseError(ctx, http.StatusInternalServerError)
		return
	}

	util.APIResponse(ctx, http.StatusOK, gin.H{
		"total": total,
		"data":  result,
	})
}

func (h *handler) ShowNft(ctx *gin.Context) {
	var input InputShowNft
	if err := ctx.ShouldBindUri(&input); err != nil {
		util.ValidatorErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	result, err := h.nftService.GetTokenById(uint(*input.Id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			util.APIResponseError(ctx, http.StatusNotFound)
			return
		}
		log.Error().Err(err).Str("query", ctx.Request.URL.RequestURI()).Msg("h.service.GetTokenById")
		util.APIResponseError(ctx, http.StatusInternalServerError)
		return
	}

	util.APIResponse(ctx, http.StatusOK, result)
}
