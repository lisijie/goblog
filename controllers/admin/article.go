package admin

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/lisijie/goblog/models"
	"strconv"
	"strings"
)

type ArticleController struct {
	baseController
}

//管理
func (this *ArticleController) List() {
	status, _ := this.GetInt("status")
	page, _ := strconv.ParseInt(this.Ctx.Input.Param(":page"), 10, 0)
	if page < 1 {
		page = 1
	}
	pagesize := int64(10)
	offset := (page - 1) * pagesize

	var list []*models.Post
	var post models.Post
	o := orm.NewOrm()
	count, _ := o.QueryTable(&post).Filter("status", status).Count()
	if count > 0 {
		o.QueryTable(&post).Filter("status", status).OrderBy("-istop", "-posttime").Limit(pagesize, offset).All(&list)
	}

	this.Data["count_1"], _ = o.QueryTable(&post).Filter("status", 1).Count()
	this.Data["count_2"], _ = o.QueryTable(&post).Filter("status", 2).Count()
	this.Data["status"] = status
	this.Data["count"] = count
	this.Data["list"] = list
	this.Data["pagebar"] = models.NewPager(page, count, pagesize, fmt.Sprintf("/admin/article/list?status=%d", status)).ToString()
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
	color := this.GetString("color")
	status, _ := this.GetInt("status")
	var istop int8 = 0
	var urltype int8 = 1
	if this.GetString("istop") == "1" {
		istop = 1
	}
	if this.GetString("urltype") == "2" {
		urltype = 2
	}
	if status != 1 && status != 2 {
		status = 0
	}

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
		post.Posttime = this.getTime()
		post.Insert()
	} else {
		post.Id = id
		if post.Read() != nil {
			goto RD
		}
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
			tp := models.TagPost{Tagid: tag.Id, Postid: post.Id, Poststatus: int8(status), Posttime: post.Posttime}
			tp.Insert()
		}
		post.Tags = strings.Join(addtags, ",")
	}

	post.Status = int8(status)
	post.Title = title
	post.Color = color
	post.Istop = istop
	post.Content = content
	post.Urlname = urlname
	post.Urltype = urltype
	post.Updated = this.getTime()
	post.Update()

RD:
	this.Redirect("/admin/article/list", 302)
}

//删除
func (this *ArticleController) Delete() {
	id, _ := this.GetInt("id")
	post := models.Post{Id: id}
	if post.Read() == nil {
		post.Delete()
	}
	this.Redirect("/admin/article/list", 302)
}

//批处理
func (this *ArticleController) Batch() {
	ids := this.GetStrings("ids[]")
	op := this.GetString("op")

	idarr := make([]int64, 0)
	for _, v := range ids {
		if id, _ := strconv.Atoi(v); id > 0 {
			idarr = append(idarr, int64(id))
		}
	}

	switch op {
	case "topub": //移到已发布
		orm.NewOrm().QueryTable(&models.Post{}).Filter("id__in", idarr).Update(orm.Params{"status": 0})
	case "todrafts": //移到草稿箱
		orm.NewOrm().QueryTable(&models.Post{}).Filter("id__in", idarr).Update(orm.Params{"status": 1})
	case "totrash": //移到回收站
		orm.NewOrm().QueryTable(&models.Post{}).Filter("id__in", idarr).Update(orm.Params{"status": 2})
	case "delete": //批量删除
		for _, id := range idarr {
			post := models.Post{Id: id}
			if post.Read() == nil {
				post.Delete()
			}
		}
	}

	this.Redirect(this.Ctx.Request.Referer(), 302)
}
