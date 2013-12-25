package blog

import (
	"github.com/astaxie/beego"
	"github.com/lisijie/goblog/models"
	"os"
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
	this.options = models.GetOptions()
	this.Data["options"] = this.options
}

func (this *baseController) display(tpl string) {
	var theme string
	if v, ok := this.options["theme"]; ok && v != "" {
		theme = v
	} else {
		theme = "default"
	}
	if _, err := os.Stat(beego.ViewsPath + "/" + theme + "/layout.html"); err == nil {
		this.Layout = theme + "/layout.html"
	}
	this.TplNames = theme + "/" + tpl + ".html"
}
