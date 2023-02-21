package jsonrpc

import "github.com/osamingo/jsonrpc/v2"

// Method представляет собой JSONRPC метод
type Method interface {
	jsonrpc.Handler
	Name() string
	Params() any
	Result() any
}
