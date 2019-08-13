package rpc

import (
	"context"
	"io"
)

type Call struct {
	ServiceMethod string
	Args          interface{}
	Reply         interface{}
	Error         error
	Done          chan *Call
}

type RPCClient interface {
	Go(ctx context.Context, serviceMethod string, arg interface{}, reply interface{}, done chan *Call) *Call
	Call(ctx context.Context, serviceMethod string, arg interface{}, reply interface{}) error
}

type Listener interface {
	Listen(network, addr string) error
	Accept() (Transport, error)
	io.Closer
}
