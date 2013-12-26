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
	this.moduleName = "blog"
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

func (this *baseController) getOption(name string) string {
	if v, ok := this.options[name]; ok {
		return v
	} else {
		return ""
	}
}

func (this *baseController) setHeadMetas(params ...string) {
	title_buf := make([]string, 0, 3)
	if len(params) == 0 && this.getOption("subtitle") != "" {
		title_buf = append(title_buf, this.getOption("subtitle"))
	}
	if len(params) > 0 {
		title_buf = append(title_buf, params[0])
	}
	title_buf = append(title_buf, this.getOption("sitename"))
	this.Data["title"] = strings.Join(title_buf, " - ")

	if len(params) > 1 {
		this.Data["keywords"] = params[1]
	} else {
		this.Data["keywords"] = this.getOption("keywords")
	}

	if len(params) > 2 {
		this.Data["description"] = params[2]
	} else {
		this.Data["description"] = this.getOption("description")
	}

}
