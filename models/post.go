package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

type Post struct {
	Id       int64
	Userid   int64     `orm:"index"`
	Author   string    `orm:"size(15)"`
	Title    string    `orm:"size(100)"`
	Urlname  string    `orm:"size(100);index"`
	Content  string    `orm:"type(text)"`
	Tags     string    `orm:"size(100)"`
	Posttime time.Time `orm:"auto_now_add;type(datetime);index"`
	Views    int64
	Status   int8
	Updated  time.Time `orm:"auto_now;type(datetime)"`
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

//内容URL
func (m *Post) Link() string {
	if m.Urlname != "" {
		return fmt.Sprintf("/article/%s", Rawurlencode(m.Urlname))
	} else {
		return fmt.Sprintf("/article/%d", m.Id)
	}
}

//带链接的标签
func (m *Post) TagsLink() string {
	return Tags2html(m.Tags)
}
