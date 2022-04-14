package get

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"orbit_nft/api/util/validation"
	"orbit_nft/db/model"
	"orbit_nft/db/repository"
	"orbit_nft/db/service"
	"orbit_nft/testutil"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
)

func setupGetHandler(t *testing.T) (*handler, func()) {
	db, teardown := testutil.NewMockDB()
	repo := repository.NewNftRepository(db)
	service := service.NewNftService(repo)
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

func assertGetHandler(t *testing.T, expect gin.H, actual gin.H, keys ...string) {
	for _, key := range keys {
		assert.Equal(t, expect[key], actual[key])
	}
}

func TestListNftOfAddressWithAddressNotExists(t *testing.T) {
	getHandler, teardown := setupGetHandler(t)
	defer teardown()

	db, teardown := testutil.NewMockDB()
	defer teardown()
	nftRepo := repository.NewNftRepository(db)

	err := nftRepo.CreateNotForSaleToken(&model.NFT{
		Name:        "name nft",
		MetadataURI: "metadata NFT",
		TokenId:     0,
		Type:        "character",
		Owner:       "0xabcd",
		Image:       "image",
		Class:       "mage",
		Rarity:      "normal",
	})
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	url := fmt.Sprintf("/nft/%s/lists", "0x12345")
	req, err := http.NewRequest("GET", url, nil)
	assert.NoError(t, err)
	param := gin.Param{Key: "address", Value: "0x12345"}
	ctx.Params = []gin.Param{param}
	ctx.Request = req

	getHandler.GetTokenOfAddress(ctx)
	assert.Equal(t, http.StatusOK, w.Code)
	expected := gin.H{
		"total": float64(0),
	}
	var actual gin.H
	json.Unmarshal(w.Body.Bytes(), &actual)
	fmt.Println(actual)
	assertGetHandler(t, expected, actual, "total")
}

func TestListNftOfAddress(t *testing.T) {
	getHandler, teardown := setupGetHandler(t)
	defer teardown()

	db, teardown := testutil.NewMockDB()
	defer teardown()
	nftRepo := repository.NewNftRepository(db)

	owner := "0x12345"

	err := nftRepo.CreateNotForSaleToken(&model.NFT{
		Name:        "name nft",
		MetadataURI: "metadata NFT",
		TokenId:     1,
		Type:        "character",
		Owner:       owner,
		Image:       "image",
		Class:       "mage",
		Rarity:      "normal",
	})
	assert.NoError(t, err)

	err = nftRepo.CreateNotForSaleToken(&model.NFT{
		Name:        "name nft",
		MetadataURI: "metadata NFT",
		TokenId:     2,
		Type:        "character",
		Owner:       "0xabcd",
		Image:       "image",
		Class:       "mage",
		Rarity:      "normal",
	})
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	url := fmt.Sprintf("/nft/%s/lists", owner)
	req, err := http.NewRequest("GET", url, nil)
	assert.NoError(t, err)
	param := gin.Param{Key: "address", Value: owner}
	ctx.Params = []gin.Param{param}
	ctx.Request = req

	getHandler.GetTokenOfAddress(ctx)
	assert.Equal(t, http.StatusOK, w.Code)
	expected := gin.H{
		"total": float64(1),
	}
	var actual gin.H
	json.Unmarshal(w.Body.Bytes(), &actual)
	assertGetHandler(t, expected, actual, "total")
}
