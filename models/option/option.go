package option

import (
	"github.com/lisijie/goblog/models"
	"github.com/lisijie/goblog/util/singleton"
	"strconv"
)

func GetOptions() map[string]string {
	if !singleton.Cache().IsExist("options") {
		var result []*models.Option
		new(models.Option).Query().All(&result)
		options := make(map[string]string)
		for _, v := range result {
			options[v.Name] = v.Value
		}
		singleton.Cache().Put("options", options, 0)
	}
	v := singleton.Cache().Get("options")
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
