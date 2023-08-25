package main

import (
	"fmt"
	"go/bootstrap"
	"go/global"
	"sync"
)

// CreateContainersFactory 创建一个容器工厂
func CreateContainersFactory() *containers {
	return &containers{}
}

// 定义一个容器结构体
type containers struct {
}

var sMap sync.Map
var containerFactory = CreateContainersFactory()

func (c *containers) KeyIsExists(key string) (interface{}, bool) {
	return sMap.Load(key)
}

func (c *containers) Get(key string) interface{} {
	if value, exists := c.KeyIsExists(key); exists {
		return value
	}
	return nil
}
func main() {
	// log.Fatal("666")
	// sMap.Store("dd.dd", "value")

	// data := containerFactory.Get("dd.dd")
	// fmt.Println(data)
	// bootstrap.InitializeConfig()
	global.ConfigYml = bootstrap.InitializeConfigs()
	fmt.Println(global.ConfigYml.Get("database.password"))
	ConfigYml := global.ConfigYml
	// dsn := ConfigYml.GetString("database.username") + ":"
	dsn := ConfigYml.GetString("database.username") + ":" + ConfigYml.GetString("database.password") + "@(" + ConfigYml.GetString("database.host") + ":" + ConfigYml.GetString("database.port") + ")/" +
		ConfigYml.GetString("database.database") + "?charset=" + ConfigYml.GetString("database.charset") + "&parseTime=True&loc=Local"

	fmt.Println(dsn)
	// c.Get("database")
	// fmt.Println(global.App.Config.Database.UserName)
	// fmt.Println(Database)
	// dsn := Database.UserName + ":" + Database.Password + "@(" + Database.Host + ":" + strconv.Itoa(Database.Port) + ")/" +
	// 	Database.Database + "?charset=" + Database.Charset + "&parseTime=True&loc=Local"
	// fmt.Println("======================")
	// fmt.Println(dsn)
	// routers.Init()
}
