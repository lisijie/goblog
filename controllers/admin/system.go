package admin

import (
	"github.com/astaxie/beego/orm"
	"github.com/lisijie/goblog/models"
)

type SystemController struct {
	baseController
}

//系统设置
func (this *SystemController) Setting() {
	var result []*models.Option
	o := orm.NewOrm()
	o.QueryTable(&models.Option{}).All(&result)

	options := make(map[string]string)
	mp := make(map[string]*models.Option)
	for _, v := range result {
		options[v.Name] = v.Value
		mp[v.Name] = v
	}

	if this.Ctx.Request.Method == "POST" {
		o := orm.NewOrm()
		keys := []string{"sitename", "siteurl", "subtitle", "pagesize", "keywords", "description", "email", "theme", "timezone"}
		for _, key := range keys {
			val := this.GetString(key)
			if _, ok := mp[key]; !ok {
				option := new(models.Option)
				option.Name = key
				option.Value = val
				options[key] = val
				o.Insert(option)
			} else {
				option := mp[key]
				option.Value = val
				o.Update(option, "Value")
			}
		}
		models.Cache.Delete("options")
		this.Redirect("/admin/system/setting", 302)
	}

	this.Data["now"] = this.getTime()
	this.Data["options"] = options
	this.display()
}

func (this *SystemController) Stat() {

}
