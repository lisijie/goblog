package blog

import (
	"github.com/astaxie/beego/orm"
	"github.com/lisijie/goblog/models"
	"net/url"
	"strconv"
)

type MainController struct {
	baseController
}

//首页, 只显示前N条
func (this *MainController) Index() {
	var (
		list     []*models.Post
		post     models.Post
		pagesize int
		err      error
	)

	if pagesize, err = strconv.Atoi(this.options["pagesize"]); err != nil || pagesize < 1 {
		pagesize = 10
	}

	o := orm.NewOrm()
	o.QueryTable(&post).OrderBy("-id").Limit(pagesize).All(&list)

	this.Data["list"] = list

	this.display("index")
}

//文章显示
func (this *MainController) Show() {
	var (
		post models.Post
		err  error
	)

	urlname := this.Ctx.Input.Param(":urlname")
	if urlname != "" {
		post.Urlname = urlname
		err = post.Read("urlname")
	} else {
		id, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))
		post.Id = int64(id)
		err = post.Read()
	}
	if err != nil {
		this.Abort("404")
	}

	this.Data["post"] = post
	this.display("article")
}

//历史归档
func (this *MainController) Archives() {
	var (
		page     int
		pagesize int
		err      error
		count    int64
		result   map[string][]*models.Post
	)

	if page, err = strconv.Atoi(this.Ctx.Input.Param("page")); err != nil || page < 1 {
		page = 1
	}

	if pagesize, err = strconv.Atoi(this.options["pagesize"]); err != nil || pagesize < 1 {
		pagesize = 20
	} else {
		pagesize *= 2
	}

	o := orm.NewOrm()
	count, _ = o.QueryTable(&models.Post{}).Count()
	result = make(map[string][]*models.Post)
	if count > 0 {
		var list []*models.Post
		o.QueryTable(&models.Post{}).OrderBy("-id").Limit(pagesize, (page-1)*pagesize).All(&list)
		for _, v := range list {
			year := v.Posttime.Format("2006")
			if _, ok := result[year]; !ok {
				result[year] = make([]*models.Post, 0)
			}
			result[year] = append(result[year], v)
		}
	}

	this.Data["page"] = page
	this.Data["count"] = count
	this.Data["pagesize"] = pagesize
	this.Data["pagebar"] = models.Pager(int64(page), int64(count), int64(pagesize), "/archives")
	this.Data["result"] = result

	this.display("archives")
}

//分类查看
func (this *MainController) Category() {
	var (
		page     int
		pagesize int
		name     string
		err      error
		count    int64
		result   map[string][]*models.Post
	)
	name = this.Ctx.Input.Param(":name")
	if page, err = strconv.Atoi(this.Ctx.Input.Param(":page")); err != nil || page < 1 {
		page = 1
	}
	if pagesize, err = strconv.Atoi(this.options["pagesize"]); err != nil || pagesize < 1 {
		pagesize = 20
	} else {
		pagesize *= 2
	}

	o := orm.NewOrm()
	tag := new(models.Tag)
	tag.Name = name

	if o.Read(tag, "Name") != nil {
		this.Abort("404")
	}

	count, _ = o.QueryTable(&models.TagPost{}).Filter("tagid", tag.Id).Count()
	result = make(map[string][]*models.Post)
	if count > 0 {
		var tp []*models.TagPost
		var list []*models.Post
		var pids []int64 = make([]int64, 0)

		o.QueryTable(&models.TagPost{}).Filter("tagid", tag.Id).Limit(pagesize, (page-1)*pagesize).All(&tp)
		for _, v := range tp {
			pids = append(pids, v.Postid)
		}

		o.QueryTable(&models.Post{}).Filter("id__in", pids).All(&list)

		for _, v := range list {
			year := v.Posttime.Format("2006")
			if _, ok := result[year]; !ok {
				result[year] = make([]*models.Post, 0)
			}
			result[year] = append(result[year], v)
		}
	}

	this.Data["tag"] = tag
	this.Data["page"] = page
	this.Data["pagesize"] = pagesize
	this.Data["count"] = count
	this.Data["result"] = result
	this.Data["pagebar"] = models.Pager(int64(page), int64(count), int64(pagesize), "/category/"+url.QueryEscape(tag.Name))

	this.display("category")
}
