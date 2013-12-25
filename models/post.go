package models

import (
	"bytes"
	"fmt"
	"github.com/astaxie/beego/orm"
	"strings"
	"time"
)

type Post struct {
	Id       int64
	Userid   int64  `orm:"index"`
	Author   string `orm:"size(15)"`
	Title    string `orm:"size(100)"`
	Color    string `orm:"size(7)"`
	Urlname  string `orm:"size(100);index"`
	Urltype  int8
	Excerpt  string    `orm:"type(text)"`
	Content  string    `orm:"type(text)"`
	Tags     string    `orm:"size(100)"`
	Posttime time.Time `orm:"auto_now_add;type(datetime);index"`
	Views    int64
	Status   int8
	Updated  time.Time `orm:"auto_now;type(datetime)"`
	Istop    int8
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

//带颜色的标题
func (m *Post) ColorTitle() string {
	if m.Color != "" {
		return fmt.Sprintf("<span style=\"color:%s\">%s</span>", m.Color, m.Title)
	} else {
		return m.Title
	}
}

//内容URL
func (m *Post) Link() string {
	if m.Urlname != "" {
		if m.Urltype == 2 {
			return fmt.Sprintf("/article/%s", Rawurlencode(m.Urlname))
		} else {
			return fmt.Sprintf("/%s", Rawurlencode(m.Urlname))
		}
	} else {
		return fmt.Sprintf("/article/%d", m.Id)
	}
}

//带链接的标签
func (m *Post) TagsLink() string {
	if m.Tags == "" {
		return ""
	}
	var buf bytes.Buffer
	arr := strings.Split(m.Tags, ",")
	for k, v := range arr {
		if k > 0 {
			buf.WriteString(", ")
		}
		tag := Tag{Name: v}
		buf.WriteString(tag.Link())
	}
	return buf.String()
}
