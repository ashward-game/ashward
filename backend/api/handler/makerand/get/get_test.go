package get

import (
	"encoding/json"
	"net"
	"net/http"
	"net/http/httptest"
	"net/rpc"
	"orbit_nft/api/context"
	"orbit_nft/api/util/validation"
	"orbit_nft/crypto"
	"orbit_nft/db/repository"
	"orbit_nft/db/service"
	"orbit_nft/testutil"
	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
)

type SigningService struct {
}

var signingAddress = "localhost:9001"

func (ss *SigningService) Commit(ignore int, com *crypto.CommitmentWithSig) error {
	*com = crypto.CommitmentWithSig{
		Commitment: "0x00",
		Signature:  "0x01",
	}
	return nil
}

func setupMockSigningService(t *testing.T) {
	ss := &SigningService{}
	rpc.Register(ss)
	rpc.HandleHTTP()
	l, err := net.Listen("tcp", signingAddress)
	assert.NoError(t, err)
	http.Serve(l, nil)
}

func setupGetHandler(t *testing.T) (*handler, func()) {
	db, teardown := testutil.NewMockDB()
	repo := repository.NewMakeRandRepository(db)
	service := service.NewMakeRandService(repo)
	getHandler := NewHandler(service)

	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	log.Logger = log.Output(
		zerolog.ConsoleWriter{
			Out:     os.Stderr,
			NoColor: false,
		},
	)

	gin.SetMode(gin.TestMode)

	err := validation.Register()
	assert.NoError(t, err)

	return getHandler, func() {
		teardown()
	}
}

func TestNewCommit(t *testing.T) {
	getHandler, teardown := setupGetHandler(t)
	defer teardown()

	go setupMockSigningService(t)

	// delay ensuring server already started
	time.Sleep(time.Second * 2)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Set(context.KeySigningServiceRPCAddress, signingAddress)

	req, err := http.NewRequest("GET", "/makerand/commit", nil)
	assert.NoError(t, err)
	q := req.URL.Query()
	req.URL.RawQuery = q.Encode()
	ctx.Request = req

	getHandler.Commit(ctx)
	assert.Equal(t, http.StatusOK, w.Code)
	var actual gin.H
	json.Unmarshal(w.Body.Bytes(), &actual)
	assert.NotEmpty(t, actual["hash"])
	assert.NotEmpty(t, actual["signature"])
}
