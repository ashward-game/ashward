package get

import (
	"encoding/json"
	"fmt"
	"math/big"
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

func TestInputValidationRule(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	err := validation.Register()
	assert.NoError(t, err)

	req, err := http.NewRequest("GET", "/test/", nil)
	assert.NoError(t, err)

	q := req.URL.Query()
	req.URL.RawQuery = q.Encode()
	ctx.Request = req

	var validate InputList

	// no type
	err = ctx.ShouldBind(&validate)
	assert.EqualError(t, err, "Key: 'InputList.Type' Error:Field validation for 'Type' failed on the 'required' tag")

	// invalid type type
	ctx.Request.Form.Add("type", "invalid")
	err = ctx.ShouldBind(&validate)
	assert.EqualError(t, err, "Key: 'InputList.Type' Error:Field validation for 'Type' failed on the 'oneof' tag")

	// with valid type
	ctx.Request.Form.Del("type")
	ctx.Request.Form.Add("type", "character")
	err = ctx.ShouldBind(&validate)
	assert.NoError(t, err)
}

func setupGetHandler(t *testing.T) (*handler, func()) {
	db, teardown := testutil.NewMockDB()
	repoNft := repository.NewNftRepository(db)
	serviceNft := service.NewNftService(repoNft)
	repoMarket := repository.NewMarketplaceRepository(db)
	serviceMarket := service.NewMarketplaceService(repoMarket, repoNft)
	getHandler := NewHandler(serviceNft, serviceMarket)

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

func TestGetNoType(t *testing.T) {
	getHandler, teardown := setupGetHandler(t)
	defer teardown()

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	req, err := http.NewRequest("GET", "/test/", nil)
	assert.NoError(t, err)
	q := req.URL.Query()
	req.URL.RawQuery = q.Encode()
	ctx.Request = req

	getHandler.Get(ctx)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestTypeNotExists(t *testing.T) {
	getHandler, teardown := setupGetHandler(t)
	defer teardown()

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	url := fmt.Sprintf("/marketplace?type=%s", "test")
	req, err := http.NewRequest("GET", url, nil)
	assert.NoError(t, err)
	q := req.URL.Query()
	req.URL.RawQuery = q.Encode()
	ctx.Request = req

	getHandler.Get(ctx)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestTypeExists(t *testing.T) {
	getHandler, teardown := setupGetHandler(t)
	defer teardown()

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	url := fmt.Sprintf("/marketplace?type=%s", "character")
	req, err := http.NewRequest("GET", url, nil)
	assert.NoError(t, err)
	q := req.URL.Query()
	req.URL.RawQuery = q.Encode()
	ctx.Request = req

	getHandler.Get(ctx)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestClassSupport(t *testing.T) {
	getHandler, teardown := setupGetHandler(t)
	defer teardown()

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	url := fmt.Sprintf("/marketplace?type=%s&class=%s", "character", "mage")
	req, err := http.NewRequest("GET", url, nil)
	assert.NoError(t, err)
	q := req.URL.Query()
	req.URL.RawQuery = q.Encode()
	ctx.Request = req

	getHandler.Get(ctx)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestClassNotSupport(t *testing.T) {
	getHandler, teardown := setupGetHandler(t)
	defer teardown()

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	url := fmt.Sprintf("/marketplace?type=%s&class=%s", "character", "abc")
	req, err := http.NewRequest("GET", url, nil)
	assert.NoError(t, err)
	q := req.URL.Query()
	req.URL.RawQuery = q.Encode()
	ctx.Request = req

	getHandler.Get(ctx)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestRarityNotExists(t *testing.T) {
	getHandler, teardown := setupGetHandler(t)
	defer teardown()

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	url := fmt.Sprintf("/marketplace?type=%s&rarity=%s", "character", "test")
	req, err := http.NewRequest("GET", url, nil)
	assert.NoError(t, err)
	q := req.URL.Query()
	req.URL.RawQuery = q.Encode()
	ctx.Request = req

	getHandler.Get(ctx)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestRarityExists(t *testing.T) {
	getHandler, teardown := setupGetHandler(t)
	defer teardown()

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	url := fmt.Sprintf("/marketplace?type=%s&rarity=%s", "character", "normal")
	req, err := http.NewRequest("GET", url, nil)
	assert.NoError(t, err)
	q := req.URL.Query()
	req.URL.RawQuery = q.Encode()
	ctx.Request = req

	getHandler.Get(ctx)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDoNotSearchSymbol(t *testing.T) {
	getHandler, teardown := setupGetHandler(t)
	defer teardown()

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	url := fmt.Sprintf("/marketplace?type=%s&search=%s", "character", "@")
	req, err := http.NewRequest("GET", url, nil)
	assert.NoError(t, err)
	q := req.URL.Query()
	req.URL.RawQuery = q.Encode()
	ctx.Request = req

	getHandler.Get(ctx)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestSearchNumneric(t *testing.T) {
	getHandler, teardown := setupGetHandler(t)
	defer teardown()

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	url := fmt.Sprintf("/marketplace?type=%s&search=%s", "character", "123")
	req, err := http.NewRequest("GET", url, nil)
	assert.NoError(t, err)
	q := req.URL.Query()
	req.URL.RawQuery = q.Encode()
	ctx.Request = req

	getHandler.Get(ctx)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestSearchStringWithSpace(t *testing.T) {
	getHandler, teardown := setupGetHandler(t)
	defer teardown()

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	url := fmt.Sprintf("/marketplace?type=%s&search=%s", "character", "token 1")
	req, err := http.NewRequest("GET", url, nil)
	assert.NoError(t, err)
	q := req.URL.Query()
	req.URL.RawQuery = q.Encode()
	ctx.Request = req

	getHandler.Get(ctx)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestSearchString(t *testing.T) {
	getHandler, teardown := setupGetHandler(t)
	defer teardown()

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	url := fmt.Sprintf("/marketplace?type=%s&search=%s", "character", "token")
	req, err := http.NewRequest("GET", url, nil)
	assert.NoError(t, err)
	q := req.URL.Query()
	req.URL.RawQuery = q.Encode()
	ctx.Request = req

	getHandler.Get(ctx)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestOrderPriceNotContainsAscOrDesc(t *testing.T) {
	getHandler, teardown := setupGetHandler(t)
	defer teardown()

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	url := fmt.Sprintf("/marketplace?type=%s&order_by_price=%s", "character", "test")
	req, err := http.NewRequest("GET", url, nil)
	assert.NoError(t, err)
	q := req.URL.Query()
	req.URL.RawQuery = q.Encode()
	ctx.Request = req

	getHandler.Get(ctx)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestOrderPriceIsAscOrDesc(t *testing.T) {
	getHandler, teardown := setupGetHandler(t)
	defer teardown()

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	url := fmt.Sprintf("/marketplace?type=%s&order_by_price=%s", "character", "desc")
	req, err := http.NewRequest("GET", url, nil)
	assert.NoError(t, err)
	q := req.URL.Query()
	req.URL.RawQuery = q.Encode()
	ctx.Request = req

	getHandler.Get(ctx)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestPageIsNotNumeric(t *testing.T) {
	getHandler, teardown := setupGetHandler(t)
	defer teardown()

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	url := fmt.Sprintf("/marketplace?type=%s&page=%s", "character", "test")
	req, err := http.NewRequest("GET", url, nil)
	assert.NoError(t, err)
	q := req.URL.Query()
	req.URL.RawQuery = q.Encode()
	ctx.Request = req

	getHandler.Get(ctx)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestPageIsZero(t *testing.T) {
	getHandler, teardown := setupGetHandler(t)
	defer teardown()

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	url := fmt.Sprintf("/marketplace?type=%s&page=%s", "character", "0")
	req, err := http.NewRequest("GET", url, nil)
	assert.NoError(t, err)
	q := req.URL.Query()
	req.URL.RawQuery = q.Encode()
	ctx.Request = req

	getHandler.Get(ctx)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestPageLtZero(t *testing.T) {
	getHandler, teardown := setupGetHandler(t)
	defer teardown()

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	url := fmt.Sprintf("/marketplace?type=%s&page=%s", "character", "-1")
	req, err := http.NewRequest("GET", url, nil)
	assert.NoError(t, err)
	q := req.URL.Query()
	req.URL.RawQuery = q.Encode()
	ctx.Request = req

	getHandler.Get(ctx)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestPageEqualOne(t *testing.T) {
	getHandler, teardown := setupGetHandler(t)
	defer teardown()

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	url := fmt.Sprintf("/marketplace?type=%s&page=%s", "character", "1")
	req, err := http.NewRequest("GET", url, nil)
	assert.NoError(t, err)
	q := req.URL.Query()
	req.URL.RawQuery = q.Encode()
	ctx.Request = req

	getHandler.Get(ctx)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestLimitLtZero(t *testing.T) {
	getHandler, teardown := setupGetHandler(t)
	defer teardown()

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	url := fmt.Sprintf("/marketplace?type=%s&limit=%s", "character", "-1")
	req, err := http.NewRequest("GET", url, nil)
	assert.NoError(t, err)
	q := req.URL.Query()
	req.URL.RawQuery = q.Encode()
	ctx.Request = req

	getHandler.Get(ctx)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestLimitIsZero(t *testing.T) {
	getHandler, teardown := setupGetHandler(t)
	defer teardown()

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	url := fmt.Sprintf("/marketplace?type=%s&limit=%s", "character", "0")
	req, err := http.NewRequest("GET", url, nil)
	assert.NoError(t, err)
	q := req.URL.Query()
	req.URL.RawQuery = q.Encode()
	ctx.Request = req

	getHandler.Get(ctx)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestLimitIsNotNumeric(t *testing.T) {
	getHandler, teardown := setupGetHandler(t)
	defer teardown()

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	url := fmt.Sprintf("/marketplace?type=%s&limit=%s", "character", "test")
	req, err := http.NewRequest("GET", url, nil)
	assert.NoError(t, err)
	q := req.URL.Query()
	req.URL.RawQuery = q.Encode()
	ctx.Request = req

	getHandler.Get(ctx)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestLimitEqualOne(t *testing.T) {
	getHandler, teardown := setupGetHandler(t)
	defer teardown()

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	url := fmt.Sprintf("/marketplace?type=%s&limit=%s", "character", "1")
	req, err := http.NewRequest("GET", url, nil)
	assert.NoError(t, err)
	q := req.URL.Query()
	req.URL.RawQuery = q.Encode()
	ctx.Request = req

	getHandler.Get(ctx)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestInputTokenIdIsZero(t *testing.T) {
	getHandler, teardown := setupGetHandler(t)
	defer teardown()

	db, teardown := testutil.NewMockDB()
	defer teardown()
	nftRepo := repository.NewNftRepository(db)

	err := nftRepo.CreateForSaleToken(&model.NFT{
		Name:        "name nft",
		MetadataURI: "metadata NFT",
		TokenId:     0,
		Type:        "character",
		Owner:       "0x000",
		Image:       "image",
		Class:       "Mage",
		Rarity:      "normal",
	}, big.NewInt(1))
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	url := fmt.Sprintf("/marketplace/nft/%s", "0")
	req, err := http.NewRequest("GET", url, nil)
	assert.NoError(t, err)
	param := gin.Param{Key: "id", Value: "0"}
	ctx.Params = []gin.Param{param}
	ctx.Request = req

	getHandler.ShowNft(ctx)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestInputTokenIdNotFound(t *testing.T) {
	getHandler, teardown := setupGetHandler(t)
	defer teardown()

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	url := fmt.Sprintf("/marketplace/nft/%s", "1000")
	req, err := http.NewRequest("GET", url, nil)
	assert.NoError(t, err)
	param := gin.Param{Key: "id", Value: "1000"}
	ctx.Params = []gin.Param{param}
	ctx.Request = req

	getHandler.ShowNft(ctx)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestInputTokenIdLtZero(t *testing.T) {
	getHandler, teardown := setupGetHandler(t)
	defer teardown()

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	url := fmt.Sprintf("/marketplace/nft/%s", "-1")
	req, err := http.NewRequest("GET", url, nil)
	assert.NoError(t, err)
	param := gin.Param{Key: "id", Value: "-1"}
	ctx.Params = []gin.Param{param}
	ctx.Request = req

	getHandler.ShowNft(ctx)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestInputTokenIsNotNumeric(t *testing.T) {
	getHandler, teardown := setupGetHandler(t)
	defer teardown()

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	url := fmt.Sprintf("/marketplace/nft/%s", "test")
	req, err := http.NewRequest("GET", url, nil)
	assert.NoError(t, err)
	param := gin.Param{Key: "id", Value: "test"}
	ctx.Params = []gin.Param{param}
	ctx.Request = req

	getHandler.ShowNft(ctx)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func assertGetHandler(t *testing.T, expect gin.H, actual gin.H, keys ...string) {
	for _, key := range keys {
		assert.Equal(t, expect[key], actual[key])
	}
}

func TestTradingOfAddressWithAddressNotExists(t *testing.T) {
	getHandler, teardown := setupGetHandler(t)
	defer teardown()

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	url := fmt.Sprintf("/marketplace/%s/history", "Oxabcd")
	req, err := http.NewRequest("GET", url, nil)
	assert.NoError(t, err)
	param := gin.Param{Key: "address", Value: "Oxabcd"}
	ctx.Params = []gin.Param{param}
	ctx.Request = req

	getHandler.GetTradingHistory(ctx)
	assert.Equal(t, http.StatusOK, w.Code)

	var actual gin.H
	expect := gin.H{
		"data":  []interface{}{},
		"total": float64(0),
	}
	err = json.Unmarshal(w.Body.Bytes(), &actual)
	assert.NoError(t, err)

	assertGetHandler(t, expect, actual, "data", "total")
}

func TestTradingOfAddressHasTrade(t *testing.T) {
	getHandler, teardown := setupGetHandler(t)
	defer teardown()

	owner := "Oxabcd"
	buyer := "0x12345"

	db, teardown := testutil.NewMockDB()
	defer teardown()
	nftRepo := repository.NewNftRepository(db)

	err := nftRepo.CreateNotForSaleToken(&model.NFT{
		Name:        "name nft",
		MetadataURI: "metadata NFT",
		TokenId:     1,
		Type:        "character",
		Owner:       owner,
		Image:       "image",
		Class:       "Mage",
		Rarity:      "normal",
	})
	assert.NoError(t, err)

	err = getHandler.marketService.OpenOffer(1, owner, big.NewInt(1))
	assert.NoError(t, err)

	err = getHandler.marketService.Purchase(1, buyer)
	assert.NoError(t, err)

	err = getHandler.marketService.OpenOffer(1, buyer, big.NewInt(1))
	assert.NoError(t, err)

	err = getHandler.marketService.Purchase(1, owner)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	url := fmt.Sprintf("/marketplace/%s/history", owner)
	req, err := http.NewRequest("GET", url, nil)
	assert.NoError(t, err)
	param := gin.Param{Key: "address", Value: owner}
	ctx.Params = []gin.Param{param}
	ctx.Request = req

	getHandler.GetTradingHistory(ctx)
	assert.Equal(t, http.StatusOK, w.Code)

	var actual gin.H
	expect := gin.H{
		"total": float64(2),
	}
	err = json.Unmarshal(w.Body.Bytes(), &actual)
	assert.NoError(t, err)

	assertGetHandler(t, expect, actual, "total")
}
