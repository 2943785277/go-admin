package main

import (
	"fmt"
	"os"

	"gopkg.in/ini.v1"
)

var Cfg *ini.File

type IniParser struct {
	conf_reader *ini.File // config reader
}

func (this *IniParser) Readini() {
	cfg, err := ini.Load("a.ini")
	//当有异常时：
	if err != nil {
		fmt.Println("读取失败:", err)
		//	退出读取 code == 0时，表示读取成功  code == 1时，表示读取失败退出
		os.Exit(1)
	}
	this.conf_reader = cfg
	// fmt.Println("userName ==>", Cfg.Section("").Key("userName").String())
}

func (this *IniParser) GetInit(section string, key string) int64 {
	data, _ := this.conf_reader.Section(section).Key(key).Int64()
	return data
}
func (this *IniParser) GetString(section string, key string) string {
	data := this.conf_reader.Section(section).Key(key).String()
	return data
}
func main() {
	ini_parser := IniParser{}
	ini_parser.Readini()
	ss := ini_parser.GetString("", "userName")
	fmt.Println(ss)

	// asdas := Cfg.Section("pas").Key("port").Int()
	// fmt.Println(asdas)
}
