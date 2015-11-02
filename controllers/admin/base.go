package admin

import (
	"github.com/astaxie/beego"
	"github.com/sndnvaps/goblog/models"
	"github.com/sndnvaps/goblog/models/option"
	"github.com/sndnvaps/goblog/util"
	"strconv"
	"strings"
	"time"
)

type baseController struct {
	beego.Controller
	userid         int
	username       string
	moduleName     string
	controllerName string
	actionName     string
	cache          *util.LruCache
}

func (this *baseController) Prepare() {
	controllerName, actionName := this.GetControllerAndAction()
	this.moduleName = "admin"
	this.controllerName = strings.ToLower(controllerName[0 : len(controllerName)-10])
	this.actionName = strings.ToLower(actionName)
	this.auth()
	this.checkPermission()
	cache, _ := util.Factory.Get("cache")
	this.cache = cache.(*util.LruCache)
}

//登录状态验证
func (this *baseController) auth() {
	arr := strings.Split(this.Ctx.GetCookie("auth"), "|")
	if len(arr) == 2 {
		idstr, password := arr[0], arr[1]
		userid, _ := strconv.Atoi(idstr)
		if userid > 0 {
			var user models.User
			user.Id = userid
			if user.Read() == nil && password == util.Md5([]byte(this.getClientIp()+"|"+user.Password)) {
				this.userid = user.Id
				this.username = user.UserName
			}
		}
	}

	if this.userid == 0 && (this.controllerName != "account" ||
		(this.controllerName == "account" && this.actionName != "logout" && this.actionName != "login")) {
		this.Redirect("/admin/login", 302)
	}
}

//渲染模版
func (this *baseController) display(tpl ...string) {
	var tplname string
	if len(tpl) == 1 {
		tplname = this.moduleName + "/" + tpl[0] + ".html"
	} else {
		tplname = this.moduleName + "/" + this.controllerName + "/" + this.actionName + ".html"
	}
	this.Data["version"] = beego.AppConfig.String("AppVer")
	this.Data["adminid"] = this.userid
	this.Data["adminname"] = this.username
	this.Layout = this.moduleName + "/layout.html"
	this.TplNames = tplname
}

//显示错误提示
func (this *baseController) showmsg(msg ...string) {
	if len(msg) == 1 {
		msg = append(msg, this.Ctx.Request.Referer())
	}
	this.Data["adminid"] = this.userid
	this.Data["adminname"] = this.username
	this.Data["msg"] = msg[0]
	this.Data["redirect"] = msg[1]
	this.Layout = this.moduleName + "/layout.html"
	this.TplNames = this.moduleName + "/" + "showmsg.html"
	this.Render()
	this.StopRun()
}

//是否post提交
func (this *baseController) isPost() bool {
	return this.Ctx.Request.Method == "POST"
}

//获取用户IP地址
func (this *baseController) getClientIp() string {
	s := strings.Split(this.Ctx.Request.RemoteAddr, ":")

	return s[0]
}

//权限验证
func (this *baseController) checkPermission() {
	if this.userid != 1 && this.controllerName == "user" {
		this.showmsg("抱歉，只有超级管理员才能进行该操作！")
	}
}

func (this *baseController) getTime() time.Time {
	timezone, _ := strconv.ParseFloat(option.Get("timezone"), 64)
	add := timezone * float64(time.Hour)
	return time.Now().UTC().Add(time.Duration(add))
}
