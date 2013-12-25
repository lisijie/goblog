package models

import (
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"net/url"
	"strings"
)

//配置项表
type Option struct {
	Id    int64
	Name  string
	Value string
}

func init() {
	dbhost := beego.AppConfig.String("dbhost")
	dbport := beego.AppConfig.String("dbport")
	dbuser := beego.AppConfig.String("dbuser")
	dbpassword := beego.AppConfig.String("dbpassword")
	dbname := beego.AppConfig.String("dbname")
	dbprefix := beego.AppConfig.String("dbprefix")
	if dbport == "" {
		dbport = "3306"
	}
	dsn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8"
	orm.RegisterDataBase("default", "mysql", dsn)
	orm.RegisterModelWithPrefix(dbprefix, new(User))
	orm.RegisterModelWithPrefix(dbprefix, new(Post))
	orm.RegisterModelWithPrefix(dbprefix, new(Tag))
	orm.RegisterModelWithPrefix(dbprefix, new(Option))
	orm.RegisterModelWithPrefix(dbprefix, new(TagPost))
}

func Md5(buf []byte) string {
	hash := md5.New()
	hash.Write(buf)
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func Rawurlencode(str string) string {
	return strings.Replace(url.QueryEscape(str), "+", "%20", -1)
}
