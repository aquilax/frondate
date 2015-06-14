package controllers

import (
	"github.com/astaxie/beego"
	"html/template"
)

type AuthController struct {
	beego.Controller
}

func (c *AuthController) Login() {
	c.Data["xsrfdata"] = template.HTML(c.XsrfFormHtml())
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "auth/login.html"
}

func (c *AuthController) Register() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "index.tpl"
}
