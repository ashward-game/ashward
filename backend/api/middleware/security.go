package middleware

import (
	"github.com/gin-contrib/secure"
	"github.com/gin-gonic/gin"
)

func Security() gin.HandlerFunc {
	// Currently use `DefaultConfig` as a starting point
	// with strict security settings
	config := secure.DefaultConfig()
	return secure.New(config)
}
