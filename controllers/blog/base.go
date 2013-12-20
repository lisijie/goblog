package blog

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/lisijie/goblog/models"
	"strings"
)

type baseController struct {
	beego.Controller
	moduleName     string
	controllerName string
	actionName     string
	options        map[string]string
}

func (this *baseController) Prepare() {
	controllerName, actionName := this.GetControllerAndAction()
	this.moduleName = "admin"
	this.controllerName = strings.ToLower(controllerName[0 : len(controllerName)-10])
	this.actionName = strings.ToLower(actionName)
	this.options = this.getOptions()
	this.Data["options"] = this.options
}

func (this *baseController) display(tpl string) {
	this.Layout = "blog/layout.html"
	this.TplNames = "blog/" + tpl + ".html"
}

func (this *baseController) getOptions() map[string]string {
	if !models.Cache.IsExist("options") {
		var result []*models.Option
		o := orm.NewOrm()
		o.QueryTable(&models.Option{}).All(&result)
		options := make(map[string]string)
		for _, v := range result {
			options[v.Name] = v.Value
		}
		models.Cache.Put("options", options, 0)
	}
	v := models.Cache.Get("options")
	return v.(map[string]string)
}
