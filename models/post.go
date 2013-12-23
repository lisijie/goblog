package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Post struct {
	Id       int64
	Userid   int64
	Author   string
	Title    string
	Urlname  string
	Content  string `orm:"type(text)"`
	Tags     string
	Posttime time.Time `orm:"auto_now_add;type(datetime)"`
	Views    int64
	Status   int64
	Updated  time.Time `orm:"auto_now_add;type(datetime)"`
}

func (m *Post) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *Post) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Post) Update(fields ...string) error {
	fields = append(fields, "Updated")
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Post) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}
