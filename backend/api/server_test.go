package api

import (
	"io/ioutil"
	"orbit_nft/testutil"
	"path"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestServer(t *testing.T) {
	db, dbTeardown := testutil.NewMockDB()
	defer dbTeardown()

	dir, teardown := testutil.CreateTLSCertForTest(t)
	defer teardown()

	server := NewServer("development", db)
	go func() {
		server.Run(":3000", "", path.Join(dir, "server.pem"), path.Join(dir, "server.key"))
	}()

	// delay ensuring server already started
	time.Sleep(time.Second * 2)

	client := testutil.NewHTTPSClient(t)
	resp, err := client.Get("https://localhost:3000/api/v1/metadata/newblock")
	assert.NoError(t, err)

	data, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "\"pong\"", string(data))

	server.Shutdown()
}
