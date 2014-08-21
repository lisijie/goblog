package option

import (
	"strconv"
)

func GetOptions() map[string]string {
	if !mc.IsExist("options") {
		var result []*Option
		new(Option).Query().All(&result)
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
