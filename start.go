package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

//判断文件是否存在
func IsFileExist(filename string) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return true
}
func main() {
	// fmt.Println(captcha.StdHeight)
	// captcha.Server(captcha.StdWidth, captcha.StdHeight)
	ShutDownEXE()
	// ex, _ := os.Executable()
	// os.Mkdir("./1", 0666)
	// fmt.Println(os.TempDir())
	// fmt.Println(IsFileExist("./utils"))
	// fmt.Println(time.Now().Format("20060102"))
	// fmt.Println(time.Now().Unix())
	// fmt.Println(strconv.FormatInt(time.Now().Unix(), 10))
	// fmt.Println(rand.Intn(999999 - 100000))
	// fmt.Println(strconv.Itoa(rand.Intn(999999-100000) + 100000))
}

func ShutDownEXE() {

	var files []string
	root := "http"
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}
	for _, file := range files {

		fmt.Println(file)
	}

	// D:/ruanjian/Redis/redis-server.exe
	// D:/软件/HBuilderX/HBuilderX.exe
	// D:/phpstudy_pro/COM/phpstudy_pro.exe

	fileName := "1.txt"                         // txt文件路径
	data, err_read := ioutil.ReadFile(fileName) // 读取文件
	if err_read != nil {
		fmt.Println("文件读取失败！")
	}
	dataLine := strings.Split(string(data), "\n") // 将文件内容作为string按行切片
	for _, line := range dataLine {
		Start(line)
	}
	// Start("D:/软件/HBuilderX/HBuilderX.exe")
}

//启动程序
func Start(name string) {

	ps, _ := exec.LookPath(name)
	// arg := []string{"-s", "-t", "20"}  arg...
	// arg := []string{"CD D:\项目\eleemnt-plus\element-plus-admin"}

	cmd := exec.Command(ps)
	// // cmd.Run()
	err := cmd.Start()
	fmt.Println(name)
	fmt.Println("------------")
	fmt.Println(err)
	// d, err := cmd.CombinedOutput()
	// if err != nil {
	// 	fmt.Println("error")
	// 	fmt.Println("Error:", err)
	// 	// return
	// }
	// fmt.Println("开启成功")
	// fmt.Println(d)
	// fmt.Println(string(d))
	// return
}
