package admin

import (
	"strconv"
	"strings"

	"github.com/lisijie/goblog/models"
)

type AccountController struct {
	baseController
}

//登录
func (this *AccountController) Login() {
	if this.GetString("dosubmit") == "yes" {
		account := strings.TrimSpace(this.GetString("account"))
		password := strings.TrimSpace(this.GetString("password"))
		remember := this.GetString("remember")
		if account != "" && password != "" {
			var user models.User
			user.Username = account
			if user.Read("username") != nil || user.Password != models.Md5([]byte(password)) {
				this.Data["errmsg"] = "帐号或密码错误"
			} else if user.Active == 0 {
				this.Data["errmsg"] = "该帐号未激活"
			} else {
				user.Logincount += 1
				user.Lastip = this.getClientIp()
				user.Lastlogin = this.getTime()
				user.Update()
				authkey := models.Md5([]byte(this.getClientIp() + "|" + user.Password))
				if remember == "yes" {
					this.Ctx.SetCookie("auth", strconv.Itoa(user.Id)+"|"+authkey, 7*86400)
				} else {
					this.Ctx.SetCookie("auth", strconv.Itoa(user.Id)+"|"+authkey)
				}

				this.Redirect("/admin", 302)
			}
		}
	}
	this.TplNames = "admin/account_login.html"
}

//退出登录
func (this *AccountController) Logout() {
	this.Ctx.SetCookie("auth", "")
	this.Redirect("/admin/login", 302)
}

//资料修改
func (this *AccountController) Profile() {
	user := models.User{Id: this.userid}
	if err := user.Read(); err != nil {
		this.showmsg(err.Error())
	}
	if this.Ctx.Request.Method == "POST" {
		errmsg := make(map[string]string)
		password := strings.TrimSpace(this.GetString("password"))
		newpassword := strings.TrimSpace(this.GetString("newpassword"))
		newpassword2 := strings.TrimSpace(this.GetString("newpassword2"))
		updated := false
		if newpassword != "" {
			if password == "" || models.Md5([]byte(password)) != user.Password {
				errmsg["password"] = "当前密码错误"
			} else if len(newpassword) < 6 {
				errmsg["newpassword"] = "密码长度不能少于6个字符"
			} else if newpassword != newpassword2 {
				errmsg["newpassword2"] = "两次输入的密码不一致"
			}
			if len(errmsg) == 0 {
				user.Password = models.Md5([]byte(newpassword))
				user.Update("password")
				updated = true
			}
		}
		this.Data["updated"] = updated
		this.Data["errmsg"] = errmsg
	}
	this.Data["user"] = user
	this.display()
}
