package models

import (
	"bytes"
	"fmt"
	"math"
	"strings"
)

type Pager struct {
	Page     int64
	Totalnum int64
	Pagesize int64
	urlpath  string
	urlquery string
}

func NewPager(page, totalnum, pagesize int64, url string) *Pager {
	p := new(Pager)
	p.Page = page
	p.Totalnum = totalnum
	p.Pagesize = pagesize

	arr := strings.Split(url, "?")
	p.urlpath = arr[0]
	if len(arr) > 1 {
		p.urlquery = "?" + arr[1]
	}

	return p
}

func (this *Pager) url(page int64) string {
	return fmt.Sprintf("%s/page/%d%s", this.urlpath, page, this.urlquery)
}

func (this *Pager) ToString() string {
	var buf bytes.Buffer
	var from, to, linknum, offset, totalpage int64

	offset = 5
	linknum = 10
	if this.Totalnum > this.Pagesize {
		totalpage = int64(math.Ceil(float64(this.Totalnum / this.Pagesize)))
		if totalpage < linknum {
			from = 1
			to = totalpage
		} else {
			from = this.Page - offset
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
		if this.Page > 1 {
			buf.WriteString(fmt.Sprintf("<li><a href=\"%s\">&laquo;</a></li>", this.url(this.Page-1)))
		} else {
			buf.WriteString("<li class=\"disabled\"><span>&laquo;</span></li>")
		}

		if this.Page > linknum {
			buf.WriteString(fmt.Sprintf("<li><a href=\"%s\">1...</a></li>", this.url(1)))
		}

		for i := from; i <= to; i++ {
			if i == this.Page {
				buf.WriteString(fmt.Sprintf("<li class=\"active\"><span>%d</span></li>", i))
			} else {
				buf.WriteString(fmt.Sprintf("<li><a href=\"%s\">%d</a></li>", this.url(i), i))
			}
		}

		if totalpage > to {
			buf.WriteString(fmt.Sprintf("<li><a href=\"%s\">...%d</a></li>", this.url(totalpage), totalpage))
		}

		if this.Page < totalpage {
			buf.WriteString(fmt.Sprintf("<li><a href=\"%s\">&raquo;</a></li>", this.url(this.Page+1)))
		} else {
			buf.WriteString(fmt.Sprintf("<li class=\"disabled\"><span>&raquo;</span></li>"))
		}
		buf.WriteString("</ul></div>")
	}

	return buf.String()
}
