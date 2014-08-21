package option

import (
	"github.com/lisijie/goblog/models"
	"github.com/lisijie/goblog/util/cache"
	"strconv"
)

var mc *cache.LruCache

func init() {
	mc = cache.Instance(1000)
}

func GetOptions() map[string]string {
	if !mc.IsExist("options") {
		var result []*models.Option
		new(models.Option).Query().All(&result)
		options := make(map[string]string)
		for _, v := range result {
			options[v.Name] = v.Value
		}
		mc.Put("options", options, 0)
	}
	v := mc.Get("options")
	return v.(map[string]string)
}

func Get(key string) string {
	options := GetOptions()
	if v, ok := options[key]; ok {
		return v
	}
	return ""
}

func GetInt(key string) int {
	v, _ := strconv.Atoi(Get(key))
	return v
}
