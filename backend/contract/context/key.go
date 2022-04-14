package context

import "context"

type key int

const (
	// context keys that pass to handler
	KeyAddressFile key = iota
	KeyChainId
	KeyDB
	KeyOpenboxRPCAddress

	// context keys that use around handler and parser
	KeyTxHash

	// context keys for data
	KeyEthEndpoint
)

func WithOpenboxRPCAddress(ctx context.Context, openboxGenesisRPCAddr string) context.Context {
	ctx = context.WithValue(ctx, KeyOpenboxRPCAddress, openboxGenesisRPCAddr)
	return ctx
}
