package rpc

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/rpc"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	records     map[string]map[string][]string
	urlEndpoint *endpoint
)

type endpoint struct {
	mu           sync.Mutex
	currentIndex int
}

func Initialize(configFile string) {
	jsonFile, err := os.Open(configFile)
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()
	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal(jsonData, &records); err != nil {
		panic(err)
	}
	urlEndpoint = &endpoint{
		currentIndex: 0,
	}
}

func newEndpoint(chain string, chainId string) string {
	if records == nil {
		panic("rpc: need to call Initialize first")
	}
	endpoints, ok := records[chain][chainId]
	if !ok {
		err := fmt.Sprintf("rpc: no endpoint for %s of chain %s", chainId, chain)
		panic(err)
	}

	urlEndpoint.mu.Lock()
	defer urlEndpoint.mu.Unlock()
	if urlEndpoint.currentIndex == len(endpoints) {
		urlEndpoint.currentIndex = 0
	}
	urlEndpoint.currentIndex += 1
	return endpoints[urlEndpoint.currentIndex-1]
}

// Dial tries to connect to blockchain using the pre-defined list of RPC endpoints.
// It will return an error if it fails `maxTry` times (for now, this is 5 times).
// The error returned is the error from the last try.
func Dial(chain string, chainId string) (client *ethclient.Client, err error) {
	for i := 0; i < maxTry; i++ {
		endpoint := newEndpoint(chain, chainId)
		client, err = dial(endpoint)
		if err == nil {
			return client, nil
		}
	}
	return nil, err
}

func dial(endpoint string) (*ethclient.Client, error) {
	c, err := dialContext(context.Background(), endpoint)
	if err != nil {
		return nil, err
	}
	return ethclient.NewClient(c), nil
}

func dialContext(ctx context.Context, rawurl string) (*rpc.Client, error) {
	u, err := url.Parse(rawurl)
	if err != nil {
		return nil, err
	}
	switch u.Scheme {
	case "http", "https":
		return rpc.DialHTTPWithClient(rawurl, newNetClient())
	case "ws", "wss":
		return rpc.DialWebsocket(ctx, rawurl, "")
	case "stdio":
		return rpc.DialStdIO(ctx)
	case "":
		return rpc.DialIPC(ctx, rawurl)
	default:
		return nil, fmt.Errorf("no known transport for URL scheme %q", u.Scheme)
	}
}

func newNetClient() *http.Client {
	return &http.Client{
		Timeout: 60 * time.Second,
	}
}
