package option

import (
	"github.com/astaxie/beego"
	bcache "github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/orm"
	"github.com/lisijie/goblog/models/cache"
)

var mc bcache.Cache

func init() {
	mc = cache.Instance()
	orm.RegisterModel(new(Option))
}

//配置项表
type Option struct {
	Id    int64
	Name  string
	Value string
}

func (m *Option) TableName() string {
	return beego.AppConfig.String("dbprefix") + "option"
}

func (m *Option) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	mc.Delete("options")
	return nil
}

func (m *Option) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Option) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	mc.Delete("options")
	return nil
}

func (m *Option) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	mc.Delete("options")
	return nil
}

func (m *Option) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}
