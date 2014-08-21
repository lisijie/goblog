package singleton

import (
	"github.com/lisijie/goblog/util"
)

func Cache() *util.LruCache {

	obj, _ := NewSingleton(func() (interface{}, error) {
		mc := util.NewLruCache(1000)
		return mc, nil
	}).Get()

	return obj.(*util.LruCache)
}
