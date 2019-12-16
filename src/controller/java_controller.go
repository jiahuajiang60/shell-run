package controller

import (
	"github.com/astaxie/beego"
	"os"
	"publish/src/utils"
	"strings"
)

type JavaController struct {
	beego.Controller
}

type JSONS struct {
	//必须的大写开头
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

func (c *JavaController) Get() {
	dir, _ := os.Getwd()
	str := utils.ReadFile("web.cmd")
	cmds := strings.Split(str, "\n")
	for _, cmd := range cmds {
		utils.WriteLogFile(dir, []byte(cmd+"\n"))
		i := strings.Index(strings.ToLower(cmd), "cd")
		if i == 0 {
			_ = os.Chdir(strings.Split(cmd, " ")[1])
		} else {
			utils.ExecCmd(cmd, dir)
		}
	}
	data := &JSONS{"200", "部署成功"}
	c.Data["json"] = data
	utils.WriteLogFile(dir, []byte("部署成功"))
	os.Chdir(dir)
	c.ServeJSON()
}
