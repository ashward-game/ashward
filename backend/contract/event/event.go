package event

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/core/types"
)

type ParserHandler interface {
	ParseHandle(ctx context.Context, vLog *types.Log) error
	Name() string
}

var ParserHandlers = make(map[string]ParserHandler)

func Register(c ParserHandler) {
	if _, ok := ParserHandlers[c.Name()]; ok {
		panic(fmt.Sprintf("%s is already registered", c))
	}
	ParserHandlers[c.Name()] = c
}
