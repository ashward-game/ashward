package nft

import (
	"context"
	"errors"
	"math/big"
	"orbit_nft/constant"
	orbitContext "orbit_nft/contract/context"
	"orbit_nft/db/service"
	"orbit_nft/testutil"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

type mockLogTransfer struct {
	*LogTransfer
	Uri string
}

var mintEvent = &mockLogTransfer{
	&LogTransfer{
		From:    constant.AddressZero(),
		To:      common.HexToAddress("0x0dead"),
		TokenId: common.Big1,
	},
	"mintUri",
}

var transferEvent = &mockLogTransfer{
	&LogTransfer{
		From:    mintEvent.To,
		To:      common.HexToAddress("0x0beef"),
		TokenId: mintEvent.TokenId,
	},
	"transferUri",
}

var invalidTransferEvent = &mockLogTransfer{
	&LogTransfer{
		From:    common.HexToAddress("0x0baca"),
		To:      common.HexToAddress("0x0beef"),
		TokenId: mintEvent.TokenId,
	},
	"invalidTransferUri",
}

var mintInvalidMetadataEvent = &mockLogTransfer{
	&LogTransfer{
		From:    constant.AddressZero(),
		To:      common.HexToAddress("0x0beef"),
		TokenId: big.NewInt(7),
	},
	"invalidMetadataURI",
}

var mintMeta = `{}`
var invalidMetadata = `{}`

// - on event `Transfer(oldOwner, newOwner, tokenId)` with `oldOwner = 0x00` (ie, minted): update db with new `(tokenId, owner)` + `(metadataURI)` (get from contract)
// - on event `Transfer(oldOwner, newOwner, tokenId)` with `oldOwner != 0x00` (ie, transferred): update db for `tokenId` with new `newOwner`
//
// - validation `oldOwner` = `owner` (stored in db) when updating db with `Transfer`
func TestHandleLogTransfer(t *testing.T) {
	// redefine worker
	_getTokenURI = func(ctx context.Context, tokenId *big.Int) (string, error) {
		if tokenId.Cmp(mintEvent.TokenId) == 0 {
			return mintEvent.Uri, nil
		}
		if tokenId.Cmp(transferEvent.TokenId) == 0 {
			return transferEvent.Uri, nil
		}
		if tokenId.Cmp(mintInvalidMetadataEvent.TokenId) == 0 {
			return mintInvalidMetadataEvent.Uri, nil
		}
		return "", errors.New("tokenId not found")
	}

	_getTokenMetadata = func(ctx context.Context, tokenUri string) (string, error) {
		if tokenUri == mintEvent.Uri {
			return mintMeta, nil
		}
		if tokenUri == transferEvent.Uri {
			return mintMeta, nil
		}
		if tokenUri == mintInvalidMetadataEvent.Uri {
			return invalidMetadata, nil
		}
		return "", errors.New("tokenUri not found")
	}

	// setup
	db, teardown := testutil.NewMockDB()
	defer teardown()

	ctx := context.Background()
	ctx = context.WithValue(ctx, orbitContext.KeyDB, db)
	ctx = context.WithValue(ctx, orbitContext.KeyTxHash, common.HexToHash("0x00"))

	// test
	err := HandleLogTransfer(ctx, mintEvent.LogTransfer)
	assert.NoError(t, err)
	err = HandleLogTransfer(ctx, transferEvent.LogTransfer)
	assert.NoError(t, err)
	err = HandleLogTransfer(ctx, invalidTransferEvent.LogTransfer)
	expect := Err(ctx, errors.New(service.ErrNftTransferFromNotFound))
	assert.EqualError(t, err, expect.Error())
	err = HandleLogTransfer(ctx, mintInvalidMetadataEvent.LogTransfer)
	assert.NoError(t, err)
}
