package controllers

type MainController struct {
	BaseController
}

func (this *MainController) Get() {
	this.Data["Website"] = "beego.me"
	this.SetDefaultLayout("home/index.tpl")
}
