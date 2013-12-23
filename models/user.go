package models

import (
	"github.com/astaxie/beego/orm"
)

//用户表模型
type User struct {
	Id         int64
	Username   string `orm:"unique;size(15)"`
	Password   string `orm:"size(32)"`
	Lastlogin  int64
	Logincount int64
	Lastip     string
	Authkey    string `orm:"size(10)"`
}

func (m *User) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *User) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *User) Update(fields ...string) error {
	fields = append(fields, "Updated")
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *User) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}
