package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var PathData string

func init() {

}

func main() {
	RunEXE("hello.exe")
}

//启动exe
func RunEXE(strExeName string) {
	fmt.Println("启动 exe：", strExeName)
	arg := []string{"Golang", "test"}
	_ = arg
	//实现cmd启动 command: 参数1：绝对路径，参数2：启动exe的传入数据
	//获取当前路径
	strpath := getCurrentPath()
	fmt.Println("当前路径：", strpath)
	strpath = strpath + strExeName
	fmt.Println("当前路径文件：", strpath)
	cmd := exec.Command(strpath, arg...)
	if err := cmd.Run(); err != nil {
		fmt.Printf("Error:", err)
		return
	}
}

func getCurrentPath() string {
	if len(PathData) == 0 {
		s, _ := exec.LookPath(os.Args[0])
		i := strings.LastIndex(s, "\\")
		path := string(s[0 : i+1])
		PathData = path
		return path
	}
	return PathData
}
