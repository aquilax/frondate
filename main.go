package main

import (
	_ "github.com/aquilax/frondate/routers"
	"github.com/astaxie/beego"
)

func main() {
	// Session
	beego.SessionOn = true
	beego.SessionName = "fd"
	beego.SessionGCMaxLifetime = 2592000 // 1 month
	beego.Run()
}
