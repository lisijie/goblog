package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/lisijie/goblog/util"
	"strconv"
	"strings"
)

//标签表
type Tag struct {
	Id    int
	Name  string `orm:"size(20);index"`
	Count int
}

func (m *Tag) TableName() string {
	return TableName("tag")
}

func (m *Tag) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *Tag) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Tag) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

//删除
func (m *Tag) Delete() error {
	var list []*TagPost
	table := new(Post).TableName()
	new(TagPost).Query().Filter("tagid", m.Id).All(&list)
	if len(list) > 0 {
		ids := make([]string, 0, len(list))
		for _, v := range list {
			ids = append(ids, strconv.Itoa(v.PostId))
		}
		orm.NewOrm().Raw("UPDATE "+table+" SET tags = REPLACE(tags, ?,',') WHERE id IN ("+strings.Join(ids, ",")+")", ","+m.Name+",").Exec()
		new(TagPost).Query().Filter("tag_id", m.Id).Delete()
	}
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

//表查询
func (m *Tag) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}

//标签连接
func (m *Tag) Link() string {
	return fmt.Sprintf("<a class=\"category\" href=\"/category/%s\">%s</a>", util.Rawurlencode(m.Name), m.Name)
}

//更新统计
func (m *Tag) UpCount() {
	c, _ := new(TagPost).Query().Filter("tag_id", m.Id).Count()
	m.Count = int(c)
	m.Update("count")
}

//合并到另一个标签
func (m *Tag) MergeTo(to *Tag) {
	var list []*TagPost
	var tp TagPost
	tp.Query().Filter("tag_id", m.Id).All(&list)
	if len(list) > 0 {
		ids := make([]string, 0, len(list))
		for _, v := range list {
			ids = append(ids, strconv.Itoa(v.PostId))
		}
		tp.Query().Filter("tag_id", m.Id).Update(orm.Params{"tag_id": to.Id})
		orm.NewOrm().Raw("UPDATE "+new(Post).TableName()+" SET tags = REPLACE(tags, ?, ?) WHERE id IN ("+strings.Join(ids, ",")+")", ","+m.Name+",", ","+to.Name+",").Exec()
	}
}
