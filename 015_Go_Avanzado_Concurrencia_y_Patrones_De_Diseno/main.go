package main

import (
	"github.com/jsovalles/Cursos_Platzi/cache"
	"github.com/jsovalles/Cursos_Platzi/design_patterns"
	"github.com/jsovalles/Cursos_Platzi/design_patterns/observer"
	"github.com/jsovalles/Cursos_Platzi/net"
	"github.com/jsovalles/Cursos_Platzi/sync"
)

func main() {
	if !true {
		sync.SyncExample()
		cache.CacheExample()
		cache.MultipleCacheExample()
		design_patterns.FactoryExample()
		design_patterns.SingletonExample()
		design_patterns.AdapterExample()
		observer.ObserverExample()
		design_patterns.StrategyExample()
		net.NetExample()
		net.AsyncNetExample()
	}
	net.NetCat()
}
