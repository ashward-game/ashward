package get

import (
	"net/http"
	"net/rpc"
	"orbit_nft/api/context"
	"orbit_nft/api/util"
	"orbit_nft/crypto"
	"orbit_nft/db/service"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type handler struct {
	service *service.MakeRandService
}

func NewHandler(service *service.MakeRandService) *handler {
	return &handler{service: service}
}

func (h *handler) Commit(ctx *gin.Context) {
	signingServiceAddress := ctx.GetString(context.KeySigningServiceRPCAddress)

	worker, err := rpc.DialHTTP("tcp", signingServiceAddress)
	if err != nil {
		log.Error().Err(err).Msg("cannot connect to the signing service")
		util.APIResponseError(ctx, http.StatusInternalServerError)
		return
	}

	var com crypto.CommitmentWithSig
	if err := worker.Call("SigningService.Commit", 0, &com); err != nil {
		log.Error().Err(err).Msg("cannot call to SigningService.Commit")
		util.APIResponseError(ctx, http.StatusInternalServerError)
		return
	}
	if err := worker.Close(); err != nil {
		log.Error().Err(err).Msg("cannot close connection to signing service")
		util.APIResponseError(ctx, http.StatusInternalServerError)
		return
	}

	util.APIResponse(ctx, http.StatusOK, gin.H{
		"hash":      com.Commitment,
		"signature": com.Signature,
	})
}
