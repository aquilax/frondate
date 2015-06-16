package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

var siteTitle string

func (this *BaseController) storeErrorToFlash(err error) {
	msg := fmt.Sprintf("Error: ", err)
	flash := beego.NewFlash()
	beego.Error(msg)
	flash.Error(msg)
	flash.Store(&this.Controller)
}

func (this *BaseController) storeErrorToFlashString(msg string) {
	flash := beego.NewFlash()
	beego.Error(msg)
	flash.Error(msg)
	flash.Store(&this.Controller)
}

func (this *BaseController) storeNoticeToFlash(msg string) {
	flash := beego.NewFlash()
	flash.Notice(msg)
	flash.Store(&this.Controller)
}

func (this *BaseController) readAndSetFlashes() {
	flash := beego.ReadFromRequest(&this.Controller)
	if ok := flash.Data["error"]; ok != "" {
		this.Data["flash"] = ok
	}
}

func (this *BaseController) SetDefaultLayout(tplNames string) {
	this.Data["SiteTitle"] = siteTitle

	this.Layout = "main-layout.html"
	this.TplNames = tplNames

	this.LayoutSections = make(map[string]string)
	//this.LayoutSections["MainScripts"] = "js-main.html"

	// v := this.GetSession(beego.SessionName)
	// m := v.(map[string]interface{})

	// this.Data["userLogin"] = m["userlogin"].(string)
	// this.Data["userStrKey"] = m["userstrkey"].(string)

	// if (m["userfname"].(string) != "") || (m["userlname"].(string) != "") {
	// 	this.Data["userNames"] = m["userlname"].(string) + " " + m["userfname"].(string)
	// } else {
	// 	this.Data["userNames"] = m["userlogin"].(string)
	// }
}

func init() {
	siteTitle = beego.AppConfig.String("SiteTitle")
}
