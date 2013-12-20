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
	Alias    string
	Content  string `orm:"type(text)"`
	Tags     string
	Posttime time.Time `orm:"auto_now_add;type(datetime)"`
	Views    int64
	Status   int64
}

func GetPost(id int64) (*Post, error) {
	p := new(Post)
	p.Id = id
	err := orm.NewOrm().Read(p)
	return p, err
}

func DelPost(id int64) bool {
	p, err := GetPost(id)
	if err == orm.ErrNoRows {
		return false
	}
	orm.NewOrm().Delete(p)
	return true
}
