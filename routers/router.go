package routers

import (
	"github.com/aquilax/frondate/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.AutoRouter(&controllers.AuthController{})

	beego.InsertFilter("/^auth", beego.BeforeRouter, func(ctx *context.Context) {
		_, ok := ctx.Input.Session("uid").(int)
		if !ok {
			ctx.Redirect(302, "/auth/login")
		}
	})
}
