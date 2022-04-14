package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func APIResponse(ctx *gin.Context, StatusCode int, Message interface{}) {
	if StatusCode >= 400 {
		ctx.JSON(StatusCode, Message)
		defer ctx.AbortWithStatus(StatusCode)
	} else {
		ctx.JSON(StatusCode, Message)
	}
}

func APIResponseError(ctx *gin.Context, StatusCode int) {
	APIResponse(ctx, StatusCode, http.StatusText(StatusCode))
}

func ValidatorErrorResponse(ctx *gin.Context, StatusCode int, Error interface{}) {
	ctx.JSON(StatusCode, Error)
	defer ctx.AbortWithStatus(StatusCode)
}
