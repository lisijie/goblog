package admin

import (
	"github.com/lisijie/goblog/models"
)

type AccountController struct {
	baseController
}

//登录
func (this *AccountController) Login() {
	if this.GetString("dosubmit") == "yes" {
		account := this.GetString("account")
		password := this.GetString("password")
		if account != "" && password != "" {
			var user models.User
			user.Username = account
			if user.Read("username") != nil || user.Password != models.Md5([]byte(password)) {
				this.Data["errmsg"] = "帐号或密码错误"
			} else {
				this.SetSession("adminid", user.Id)
				this.SetSession("adminname", user.Username)
				this.Redirect("/admin", 302)
			}
		}
	}
	this.TplNames = "admin/account_login.html"
}

//退出登录
func (this *AccountController) Logout() {
	this.DelSession("adminid")
	this.DelSession("adminname")
	this.Redirect("/admin/login", 302)
}

//资料修改
func (this *AccountController) Profile() {
	var user models.User
	user.Id = this.userid
	if err := user.Read(); err != nil {
		this.showerr(err)
		return
	}

	this.Data["user"] = user
	this.display()
}
