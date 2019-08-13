package server

import (
	"reflect"
	"shield/rpc/codec"
	"shield/rpc/transport"
	"sync"
)

type RPCServer interface {
	Register(provider interface{}, metadata map[string]string) error
	Server(network string, addr string) error
}

type simpleServer struct {
	codec      codec.Codec
	serviceMap sync.Map
	tr         transport.ServerTransport
	mutex      sync.Mutex
	shutdown   bool
	option     Option
}

type methodType struct {
	method    reflect.Method
	ArgType   reflect.Type
	ReplyType reflect.Type
}

type service struct {
	name    string
	typ     reflect.Type
	rcvr    reflect.Value
	methods map[string]*methodType
}

func NewSimpleServer(option Option) RPCServer {
	s := new(simpleServer)
	s.option = option
	s.codec = codec.GetCodec(option.SerializeType)
	return s
}
