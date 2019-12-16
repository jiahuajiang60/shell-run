package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	constant "publish/src/constant"
)

///删除日志文件
func RemoveLogFile() string {
	cd := exec.Command("rm", "-rf", constant.LogFile)
	buf, err := cd.Output()
	if err != nil {
		return err.Error()
	}
	return string(buf)
}

///创建日志文件
func CreateLogFile() bool {
	dir, _ := os.Getwd()
	path := dir + "/" + constant.LogFile
	_, err := os.Create(path)
	if err != nil {
		fmt.Errorf(err.Error())
		return false
	}
	return true
}

///读取文件
func ReadFile(name string) string {
	dir, err := os.Getwd()
	data, err := ioutil.ReadFile(dir + "/" + name)
	if err != nil {
		fmt.Println("File reading error", err)
	}
	return string(data)
}

///写入日志文件
func WriteLogFile(dir string, bytes []byte) {
	file, err := os.OpenFile(dir+"/"+constant.LogFile, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("File write error", err)
	}
	_, err = file.WriteString(string(bytes))
	if err != nil {
		fmt.Println(err)
		fmt.Println(err.Error())
	}
}
