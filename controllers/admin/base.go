package admin

import (
	"github.com/astaxie/beego"
	"strings"
)

type baseController struct {
	beego.Controller
	userid         int64
	username       string
	moduleName     string
	controllerName string
	actionName     string
}

func (this *baseController) Prepare() {
	controllerName, actionName := this.GetControllerAndAction()
	this.moduleName = "admin"
	this.controllerName = strings.ToLower(controllerName[0 : len(controllerName)-10])
	this.actionName = strings.ToLower(actionName)

	if this.controllerName == "account" && (this.actionName == "login" || this.actionName == "logout") {

	} else {
		adminid := this.GetSession("adminid")
		if adminid == nil {
			//this.Redirect("/admin/login", 302)
		} else {
			this.userid = adminid.(int64)
			this.username = this.GetSession("adminname").(string)
		}
	}
}

func (this *baseController) display() {
	this.Data["adminid"] = this.userid
	this.Data["adminname"] = this.username
	this.Layout = this.moduleName + "/layout.html"
	this.TplNames = this.moduleName + "/" + this.controllerName + "_" + this.actionName + ".html"
}

func (this *baseController) showerr(err error) {
	this.Ctx.WriteString(err.Error())
}

//是否post提交
func (this *baseController) isPost() bool {
	return this.Ctx.Request.Method == "POST"
}
