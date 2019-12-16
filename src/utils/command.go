package utils

import (
	"bufio"
	"io"
	"os/exec"
	"strings"
)

///命令执行
func ExecCmd(cmd string, dir string) {
	cmds := strings.Split(cmd, " ")
	command := exec.Command(cmds[0], cmds[1:]...)

	//StdoutPipe方法返回一个在命令Start后与命令标准输出关联的管道。Wait方法获知命令结束后会关闭这个管道，一般不需要显式的关闭该管道。
	stdout, err := command.StdoutPipe()
	if err != nil {
		WriteLogFile(dir, []byte(err.Error()))
	}
	command.Start()

	//创建一个流来读取管道内内容，这里逻辑是通过一行一行的读取的
	reader := bufio.NewReader(stdout)

	//实时循环读取输出流中的一行内容
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}
		WriteLogFile(dir, []byte(line))
	}
	//阻塞直到该命令执行完成，该命令必须是被Start方法开始执行的
	command.Wait()
}
