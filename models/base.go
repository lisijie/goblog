package models

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"math"
	"net/url"
	"strings"
)

//配置项表
type Option struct {
	Id    int64
	Name  string
	Value string
}

func init() {
	dbhost := beego.AppConfig.String("dbhost")
	dbport := beego.AppConfig.String("dbport")
	dbuser := beego.AppConfig.String("dbuser")
	dbpassword := beego.AppConfig.String("dbpassword")
	dbname := beego.AppConfig.String("dbname")
	dbprefix := beego.AppConfig.String("dbprefix")
	if dbport == "" {
		dbport = "3306"
	}
	dsn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8"
	orm.RegisterDataBase("default", "mysql", dsn)
	orm.RegisterModelWithPrefix(dbprefix, new(User))
	orm.RegisterModelWithPrefix(dbprefix, new(Post))
	orm.RegisterModelWithPrefix(dbprefix, new(Tag))
	orm.RegisterModelWithPrefix(dbprefix, new(Option))
	orm.RegisterModelWithPrefix(dbprefix, new(TagPost))
}

func Md5(buf []byte) string {
	hash := md5.New()
	hash.Write(buf)
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func Pager(page int64, totalnum int64, pagesize int64, url string) string {
	var buf bytes.Buffer
	var from, to, linknum, offset, totalpage int64

	offset = 5
	linknum = 10
	if totalnum > pagesize {
		totalpage = int64(math.Ceil(float64(totalnum / pagesize)))
		if totalpage < linknum {
			from = 1
			to = totalpage
		} else {
			from = page - offset
			to = from + linknum
			if from < 1 {
				from = 1
				to = from + linknum - 1
			} else if to > totalpage {
				to = totalpage
				from = totalpage - linknum + 1
			}
		}

		buf.WriteString("<div class=\"pagination\"><ul>")
		if page > 1 {
			buf.WriteString(fmt.Sprintf("<li><a href=\"%s/page/%d\">&laquo;</a></li>", url, page-1))
		} else {
			buf.WriteString("<li class=\"disabled\"><span>&laquo;</span></li>")
		}

		if page > linknum {
			buf.WriteString(fmt.Sprintf("<li><a href=\"%s/page/1\">1...</a></li>", url))
		}

		for i := from; i <= to; i++ {
			if i == page {
				buf.WriteString(fmt.Sprintf("<li class=\"active\"><span>%d</span></li>", i))
			} else {
				buf.WriteString(fmt.Sprintf("<li><a href=\"%s/page/%d\">%d</a></li>", url, i, i))
			}
		}

		if totalpage > to {
			buf.WriteString(fmt.Sprintf("<li><a href=\"%s/page/%d\">...%d</a></li>", url, totalpage, totalpage))
		}

		if page < totalpage {
			buf.WriteString(fmt.Sprintf("<li><a href=\"%s/page/%d\">&raquo;</a></li>", url, page+1))
		} else {
			buf.WriteString(fmt.Sprintf("<li class=\"disabled\"><span>&raquo;</span></li>"))
		}
		buf.WriteString("</ul></div>")
	}

	return buf.String()
}

//文章标签转换为带链接的模版html
func Tags2html(tags string) template.HTML {
	var buf bytes.Buffer
	arr := strings.Split(tags, ",")
	for k, v := range arr {
		if k > 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(fmt.Sprintf("<a class=\"category\" href=\"/category/%s\">%s</a>", url.QueryEscape(v), v))
	}
	return template.HTML(buf.String())
}
