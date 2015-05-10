package admin

import (
	"github.com/astaxie/beego"
	"github.com/lionel0806/goblog/models"
	"os"
	"runtime"
)

type IndexController struct {
	baseController
}

func (this *IndexController) Index() {

	this.Data["hostname"], _ = os.Hostname()
	this.Data["version"] = beego.AppConfig.String("AppVer")
	this.Data["gover"] = runtime.Version()
	this.Data["os"] = runtime.GOOS
	this.Data["cpunum"] = runtime.NumCPU()
	this.Data["arch"] = runtime.GOARCH

	this.Data["postnum"], _ = new(models.Post).Query().Count()
	this.Data["tagnum"], _ = new(models.Tag).Query().Count()
	this.Data["usernum"], _ = new(models.User).Query().Count()

	this.display()
}
