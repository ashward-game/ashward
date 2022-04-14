package openboxgenesis

import (
	"context"
	"net"
	"net/http"
	"net/rpc"
	"orbit_nft/cmd/openbox"
	orbitContext "orbit_nft/contract/context"
	"orbit_nft/testutil"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

var mockBoxBought = &LogBoxBought{
	Buyer:        common.HexToAddress("0x123abc"),
	BoxGrade:     0,
	ServerHash:   common.HexToHash("0xserverhash"),
	ClientRandom: common.HexToHash("0xclientrandom"),
}
var rpcOpenboxService = "localhost:1234"

type OpenboxService struct{}

func (b *OpenboxService) OpenBox(params *openbox.HashWithClientRandom, ignore *int) error {
	return nil
}

func setupMockOpenboxService(t *testing.T) {
	ss := &OpenboxService{}
	err := rpc.Register(ss)
	assert.NoError(t, err)
	rpc.HandleHTTP()
	l, err := net.Listen("tcp", rpcOpenboxService)
	assert.NoError(t, err)
	err = http.Serve(l, nil)
	assert.NoError(t, err)
}

func TestBoxBoughtOpenboxServiceNotRunning(t *testing.T) {
	db, teardown := testutil.NewMockDB()
	defer teardown()

	ctx := context.Background()
	ctx = context.WithValue(ctx, orbitContext.KeyDB, db)
	ctx = context.WithValue(ctx, orbitContext.KeyTxHash, common.HexToHash("0x00"))
	ctx = context.WithValue(ctx, orbitContext.KeyOpenboxRPCAddress, rpcOpenboxService)

	err := HandleLogBoxBought(ctx, mockBoxBought)
	assert.Error(t, err)
}

func TestConnectToOpenboxService(t *testing.T) {
	go setupMockOpenboxService(t)

	// delay ensuring server already start
	time.Sleep(time.Second * 2)

	ctx := context.Background()
	ctx = context.WithValue(ctx, orbitContext.KeyOpenboxRPCAddress, rpcOpenboxService)

	err := HandleLogBoxBought(ctx, mockBoxBought)
	assert.NoError(t, err)
}
