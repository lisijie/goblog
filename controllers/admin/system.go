package admin

import (
	"github.com/lisijie/goblog/models"
)

type SystemController struct {
	baseController
}

//系统设置
func (this *SystemController) Setting() {
	var result []*models.Option
	new(models.Option).Query().All(&result)

	options := make(map[string]string)
	mp := make(map[string]*models.Option)
	for _, v := range result {
		options[v.Name] = v.Value
		mp[v.Name] = v
	}

	if this.Ctx.Request.Method == "POST" {
		keys := []string{"sitename", "siteurl", "subtitle", "pagesize", "keywords", "description", "email", "theme", "timezone", "stat"}
		for _, key := range keys {
			val := this.GetString(key)
			if _, ok := mp[key]; !ok {
				option := new(models.Option)
				option.Name = key
				option.Value = val
				options[key] = val
				option.Insert()
			} else {
				option := mp[key]
				option.Value = val
				option.Update("Value")
			}
		}
		this.Redirect("/admin/system/setting", 302)
	}

	this.Data["now"] = this.getTime()
	this.Data["options"] = options
	this.display()
}
