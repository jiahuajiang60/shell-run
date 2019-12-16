package main

import (
	"github.com/astaxie/beego"
	"publish/src/controller"
	"publish/src/utils"
)

///初始化
func Init() {
	utils.CreateLogFile()
}

func main() {
	Init()
	beego.Router("/deploy/java", &controller.JavaController{})
	beego.Router("/deploy/web", &controller.WebController{})
	beego.Router("/deploy/log", &controller.LogController{})
	beego.SetStaticPath("/", "static")
	beego.BConfig.Listen.HTTPPort = 9999
	beego.Run()
}
