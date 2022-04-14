package pagination

import (
	"fmt"
	"orbit_nft/db/util"
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	DefaultLimit = 20
	DefaultPage  = 1
	DefaultSort  = "asc"
)

type InputPagination struct {
	Page  *int `form:"page" json:"page" binding:"omitempty,min=1"`
	Limit *int `form:"limit" json:"limit" binding:"omitempty,min=1"`
}

func PaginationFromRequest(ctx *gin.Context) (*util.Pagination, error) {
	limitStr := ctx.DefaultQuery("limit", fmt.Sprint(DefaultLimit))
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		return nil, err
	}
	pageStr := ctx.DefaultQuery("page", fmt.Sprint(DefaultPage))
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		return nil, err
	}
	offset := limit * (page - 1)
	return &util.Pagination{
		Offset: offset,
		Limit:  limit,
	}, nil
}
