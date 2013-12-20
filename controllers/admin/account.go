package admin

import (
	"github.com/astaxie/beego/orm"
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
			var err error
			o := orm.NewOrm()
			user := new(models.User)
			user.Username = account
			err = o.Read(user, "Username")
			if err == orm.ErrNoRows || user.Password != models.Md5([]byte(password)) {
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
	o := orm.NewOrm()
	user := new(models.User)
	user.Id = this.userid
	err := o.Read(user)
	if err != nil {
		this.showerr(err)
		return
	}

	this.Data["user"] = user
	this.display()
}
