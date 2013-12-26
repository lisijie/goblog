package models

import (
	"github.com/astaxie/beego/orm"
)

//配置项表
type Option struct {
	Id    int64
	Name  string
	Value string
}

func (m *Option) TableName() string {
	return TableName("option")
}

func (m *Option) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	Cache.Delete("options")
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
	Cache.Delete("options")
	return nil
}

func (m *Option) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	Cache.Delete("options")
	return nil
}

func (m *Option) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}
