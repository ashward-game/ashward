// This module adds to gin.Context addresses of other RPC services that are needed
// for API handlers.
// For example, the signing service (commitment) for openbox feature.

package middleware

import (
	"orbit_nft/api/context"

	"github.com/gin-gonic/gin"
)

func SetSigningServiceRPC(address string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(context.KeySigningServiceRPCAddress, address)
	}
}
