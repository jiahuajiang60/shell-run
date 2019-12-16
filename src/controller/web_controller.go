package controller

import (
	"bufio"
	"fmt"
	"github.com/astaxie/beego"
	"io"
	"os"
	"os/exec"
	"publish/src/utils"
	"time"
)

type WebController struct {
	beego.Controller
}

func (c *WebController) Get() {
	utils.CreateLogFile()
	dir, _ := os.Getwd()
	utils.WriteLogFile(dir, []byte("==========================================================\r\n"))
	utils.WriteLogFile(dir, []byte("开始部署，时间："+time.Now().Format("2006-01-02 15:04:05")+"\r\n"))
	utils.WriteLogFile(dir, []byte("==========================================================\r\n"))
	command := exec.Command("/bin/bash", "static/shell.sh")
	//StdoutPipe方法返回一个在命令Start后与命令标准输出关联的管道。Wait方法获知命令结束后会关闭这个管道，一般不需要显式的关闭该管道。
	stdout, err := command.StdoutPipe()
	if err != nil {
		utils.WriteLogFile(dir, []byte(err.Error()))
	}
	if err := command.Start(); err != nil {
		fmt.Println("Execute failed when Start:" + err.Error())
		return
	}
	//创建一个流来读取管道内内容，这里逻辑是通过一行一行的读取的
	reader := bufio.NewReader(stdout)
	//实时循环读取输出流中的一行内容
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}
		utils.WriteLogFile(dir, []byte(line))
	}
	//阻塞直到该命令执行完成，该命令必须是被Start方法开始执行的
	command.Wait()
	utils.WriteLogFile(dir, []byte("==========================================================\r\n"))
	utils.WriteLogFile(dir, []byte("部署成功，时间："+time.Now().Format("2006-01-02 15:04:05")+"\r\n"))
	utils.WriteLogFile(dir, []byte("==========================================================\r\n"))
	data := &JSONS{"200", "部署成功"}
	c.Data["json"] = data
	c.ServeJSON()
}
