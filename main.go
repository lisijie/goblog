package main

import (
	"flag"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/lisijie/goblog/controllers/admin"
	"github.com/lisijie/goblog/controllers/blog"
	"github.com/lisijie/goblog/models"
	"os"
	"path/filepath"
)

func init() {
	var config_file string
	flag.StringVar(&config_file, "config-file", "", "the path of the config file")
	flag.Parse()
	if config_file != "" {
		beego.AppConfigPath, _ = filepath.Abs(config_file)
		beego.ParseConfig()
	} else {
		if config_file = os.Getenv("BEEGO_APP_CONFIG_FILE"); config_file != "" {
			beego.AppConfigPath, _ = filepath.Abs(config_file)
			beego.ParseConfig()
		}
	}

	models.Init()
}

func main() {
	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}

	//前台路由
	beego.Router("/", &blog.MainController{}, "*:Index")
	beego.Router("/:urlname([^/]+)", &blog.MainController{}, "*:Show") //别名访问
	beego.Router("/page/:page:int", &blog.MainController{}, "*:Index")
	beego.Router("/article/:id:int", &blog.MainController{}, "*:Show")         //ID访问
	beego.Router("/article/:urlname([^/]+)", &blog.MainController{}, "*:Show") //别名访问
	beego.Router("/archives", &blog.MainController{}, "*:Archives")
	beego.Router("/archives/page/:page:int", &blog.MainController{}, "*:Archives")
	beego.Router("/category/:name([^/]+)", &blog.MainController{}, "*:Category")
	beego.Router("/category/:name([^/]+)/page/:page:int", &blog.MainController{}, "*:Category")

	//后台路由
	beego.Router("/admin", &admin.IndexController{}, "*:Index")
	beego.Router("/admin/login", &admin.AccountController{}, "*:Login")
	beego.Router("/admin/logout", &admin.AccountController{}, "*:Logout")
	beego.Router("/admin/account/profile", &admin.AccountController{}, "*:Profile")
	//系统管理
	beego.Router("/admin/system/setting", &admin.SystemController{}, "*:Setting")
	//内容管理
	beego.Router("/admin/article/list", &admin.ArticleController{}, "*:List")
	beego.Router("/admin/article/add", &admin.ArticleController{}, "*:Add")
	beego.Router("/admin/article/edit", &admin.ArticleController{}, "*:Edit")
	beego.Router("/admin/article/save", &admin.ArticleController{}, "post:Save")
	beego.Router("/admin/article/delete", &admin.ArticleController{}, "*:Delete")
	beego.Router("/admin/article/batch", &admin.ArticleController{}, "*:Batch")
	beego.Router("/admin/article/upload", &admin.ArticleController{}, "*:Upload")
	beego.Router("/admin/tag", &admin.TagController{}, "*:Index")
	//用户管理
	beego.Router("/admin/user/list", &admin.UserController{}, "*:List")
	beego.Router("/admin/user/add", &admin.UserController{}, "*:Add")
	beego.Router("/admin/user/edit", &admin.UserController{}, "*:Edit")
	beego.Router("/admin/user/delete", &admin.UserController{}, "*:Delete")

	beego.Run()
}
