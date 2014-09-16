package util

import (
	"errors"
	"github.com/lisijie/goblog/util/singleton"
	"sync"
)

var Factory factory

func init() {
	Factory.instances = make(map[string]singleton.Singleton)
}

type factory struct {
	instances map[string]singleton.Singleton
	lock      sync.Mutex
}

func (f *factory) Set(name string, init singleton.SingletonInitFunc) bool {
	f.lock.Lock()
	defer f.lock.Unlock()
	if _, ok := f.instances[name]; !ok {
		f.instances[name] = singleton.NewSingleton(init)
		return true
	}
	return false
}

func (f *factory) Get(name string) (interface{}, error) {
	if _, ok := f.instances[name]; ok {
		return f.instances[name].Get()
	}
	return nil, errors.New("factory get error : " + name + " not exists.")
}
