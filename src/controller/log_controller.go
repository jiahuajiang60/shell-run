package controller

import (
	"github.com/astaxie/beego"
	"publish/src/utils"
)

type LogController struct {
	beego.Controller
}

func (c *LogController) Get() {
	data := &JSONS{"200", utils.ReadFile("deploy.log")}
	c.Data["json"] = data
	c.ServeJSON()
}
