package admin

import (
	"github.com/astaxie/beego/orm"
	"github.com/lisijie/goblog/models"
	"strconv"
	"strings"
	"time"
)

type ArticleController struct {
	baseController
}

//管理
func (this *ArticleController) List() {
	page, _ := strconv.ParseInt(this.Ctx.Input.Param(":page"), 10, 0)
	if page < 1 {
		page = 1
	}
	pagesize := int64(10)
	offset := (page - 1) * pagesize

	var list []*models.Post
	var post models.Post
	o := orm.NewOrm()
	count, _ := o.QueryTable(&post).Count()
	if count > 0 {
		o.QueryTable(&post).OrderBy("-id").Limit(pagesize, offset).All(&list)
	}

	this.Data["count"] = count
	this.Data["list"] = list
	this.Data["pagebar"] = models.Pager(page, count, pagesize, "/admin/article/list")
	this.display()
}

//添加
func (this *ArticleController) Add() {
	this.display()
}

//编辑
func (this *ArticleController) Edit() {
	id, _ := this.GetInt("id")
	post := models.Post{Id: id}
	if post.Read() != nil {
		this.Abort("404")
	}
	this.Data["post"] = post
	this.display()
}

//保存
func (this *ArticleController) Save() {
	o := orm.NewOrm()

	id, _ := this.GetInt("id")
	title := this.GetString("title")
	content := this.GetString("content")
	tags := this.GetString("tags")
	urlname := this.GetString("urlname")

	addtags := make([]string, 0)
	//标签过滤
	if tags != "" {
		tagarr := strings.Split(tags, ",")
		for _, v := range tagarr {
			if tag := strings.TrimSpace(v); tag != "" {
				exists := false
				for _, vv := range addtags {
					if vv == tag {
						exists = true
						break
					}
				}
				if !exists {
					addtags = append(addtags, tag)
				}
			}
		}
	}

	var post models.Post
	if id < 1 {
		post.Userid = this.userid
		post.Author = this.username
		post.Posttime = time.Now()
		post.Title = title
		post.Content = content
		post.Urlname = urlname
		post.Insert()
	} else {
		post.Id = id
		if post.Read() != nil {
			goto RD
		}
		post.Title = title
		post.Content = content
		post.Urlname = urlname
		if post.Tags != "" {
			oldtags := strings.Split(post.Tags, ",")
			//标签统计-1
			o.QueryTable(&models.Tag{}).Filter("name__in", oldtags).Update(orm.Params{"count": orm.ColValue(orm.Col_Minus, 1)})
			//删掉tag_post表的记录
			o.QueryTable(&models.TagPost{}).Filter("postid", post.Id).Delete()
		}
	}

	if len(addtags) > 0 {
		for _, v := range addtags {
			tag := models.Tag{Name: v}
			if tag.Read("Name") == orm.ErrNoRows {
				tag.Count = 1
				tag.Insert()
			} else {
				tag.Count += 1
				tag.Update("Count")
			}
			tp := models.TagPost{Tagid: tag.Id, Postid: post.Id}
			tp.Insert()
		}
		post.Tags = strings.Join(addtags, ",")
	}
	post.Update()

RD:
	this.Redirect("/admin/article/list", 302)
}

//删除
func (this *ArticleController) Delete() {
	id, _ := this.GetInt("id")
	o := orm.NewOrm()
	post := models.Post{Id: id}
	if post.Read() == nil {
		if post.Tags != "" {
			oldtags := strings.Split(post.Tags, ",")
			//标签统计-1
			o.QueryTable(&models.Tag{}).Filter("name__in", oldtags).Update(orm.Params{"count": orm.ColValue(orm.Col_Minus, 1)})
			//删掉tag_post表的记录
			o.QueryTable(&models.TagPost{}).Filter("postid", post.Id).Delete()
		}
		post.Delete()
	}

	this.Redirect("/admin/article/list", 302)
}
