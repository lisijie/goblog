package admin

import (
	"github.com/astaxie/beego/orm"
	"github.com/lisijie/goblog/models"
	"os"
	"runtime"
)

type IndexController struct {
	baseController
}

func (this *IndexController) Index() {
	o := orm.NewOrm()

	this.Data["hostname"], _ = os.Hostname()
	this.Data["version"] = "0.1.0"
	this.Data["gover"] = runtime.Version()
	this.Data["os"] = runtime.GOOS
	this.Data["cpunum"] = runtime.NumCPU()
	this.Data["arch"] = runtime.GOARCH

	this.Data["postnum"], _ = o.QueryTable(&models.Post{}).Count()
	this.Data["tagnum"], _ = o.QueryTable(&models.Tag{}).Count()
	this.Data["usernum"], _ = o.QueryTable(&models.User{}).Count()
	this.display()
}
