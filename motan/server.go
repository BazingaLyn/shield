package motan

import (
	"flag"
	motan "shield/motan/core"
	"sync"
)

type MSContext struct {
	confFile     string
	context      *motan.Context
	extFactory   motan.ExtensionFactory
	portService  map[int]motan.Exporter
	portServer   map[int]motan.Server
	serviceImpls map[string]interface{}
	registries   map[string]motan.Registry
}

var (
	serverContextMap   = make(map[string]*MSContext, 8)
	serverContextMutex sync.Mutex
)

func GetMotanServerContext(confFile string) *MSContext {
	if !flag.Parsed() {
		flag.Parse()
	}

	serverContextMutex.Lock()
	defer serverContextMutex.Unlock()
	ms := serverContextMap[confFile]

	if ms == nil {
		ms = &MSContext{confFile: confFile}
		serverContextMap[confFile] = ms
		motan.Initialize(ms)
	}

	return nil
}

func (m *MSContext) RegisterService(s interface{}, sid string) error {
	return nil
}

func (m *MSContext) Start(extfactory motan.ExtensionFactory) {

}

func (m *MSContext) ServicesAvailable() {

}
