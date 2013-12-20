package admin

type IndexController struct {
	baseController
}

func (this *IndexController) Index() {
	this.Data["adminid"] = this.userid
	this.Data["adminname"] = this.username
	this.display()
}
