package main

import (
	"go/Model"
	"go/bootstrap"
	"go/routers"
)

func main() {
	// global.ConfigYml = bootstrap.InitializeConfigs()
	bootstrap.Init() //配置项初始化

	Model.Init() //数据库初始化

	routers.Init() //路由初始化

}
