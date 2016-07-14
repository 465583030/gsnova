package proxy

import (
	"log"

	"github.com/yinqiwen/gsnova/common/event"
)

type Proxy interface {
	Init() error
	Serve(session *ProxySession, ev event.Event) error
}

var proxyTable = make(map[string]Proxy)

func RegisterProxy(name string, p Proxy) error {
	proxyTable[name] = p
	return nil
}

func getProxyByName(name string) Proxy {
	p, exist := proxyTable[name]
	if exist {
		return p
	}
	return nil
}

func Init() error {
	for name, p := range proxyTable {
		err := p.Init()
		if nil != err {
			log.Printf("Failed to init proxy:%s", name)
		} else {
			log.Printf("Proxy:%s init success.", name)
		}
	}
	startLocalServers()
	return nil
}