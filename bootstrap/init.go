package bootstrap

import "go/global"

func Init() {
	global.ConfigYml = InitializeConfigs()
	global.Redis = InitializeRedis()
}
