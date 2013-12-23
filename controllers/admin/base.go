package admin

import (
	"github.com/astaxie/beego"
	"github.com/lisijie/goblog/models"
	"strconv"
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
	this.auth()
}

func (this *baseController) auth() {
	if this.controllerName == "account" && (this.actionName == "login" || this.actionName == "logout") {

	} else {
		arr := strings.Split(this.Ctx.GetCookie("auth"), "|")
		if len(arr) == 2 {
			idstr, password := arr[0], arr[1]
			userid, _ := strconv.ParseInt(idstr, 10, 0)
			if userid > 0 {
				var user models.User
				user.Id = userid
				if user.Read() == nil && password == models.Md5([]byte(this.getClientIp()+"|"+user.Password)) {
					this.userid = user.Id
					this.username = user.Username
				}
			}
		}
		if this.userid == 0 {
			this.Redirect("/admin/login", 302)
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

func (this *baseController) getClientIp() string {
	s := strings.Split(this.Ctx.Request.RemoteAddr, ":")
	return s[0]
}
