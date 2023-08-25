package bootstrap

import (
	"fmt"
	"go/global"
	"os"
	"sync"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	viper *viper.Viper
	mu    *sync.Mutex
}

func InitializeConfigs() global.YmlConfigInterf {

	// 设置配置文件路径
	config := "config.yaml"
	// 生产环境可以通过设置环境变量来改变配置文件路径
	if configEnv := os.Getenv("VIPER_CONFIG"); configEnv != "" {
		config = configEnv
	}

	// 初始化 viper
	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("read config failed: %s \n", err))
	}
	fmt.Println("配置文件读取成功")
	return &Config{
		viper: v,
		mu:    new(sync.Mutex),
	}
}

//ConfigFileChangeListen 监听文件变化
func (y *Config) ConfigFileChangeListen() {
	// y.viper.OnConfigChange(func(changeEvent fsnotify.Event) {
	// 	if time.Now().Sub(lastChangeTime).Seconds() >= 1 {
	// 		if changeEvent.Op.String() == "WRITE" {
	// 			y.clearCache()
	// 			lastChangeTime = time.Now()
	// 		}
	// 	}
	// })
	// y.viper.WatchConfig()
}

// Get 一个原始值
func (y *Config) Get(keyName string) interface{} {
	value := y.viper.Get(keyName)
	// if y.keyIsCache(keyName) {
	// 	return y.getValueFromCache(keyName)
	// } else {
	// 	value := y.viper.Get(keyName)
	// 	y.cache(keyName, value)
	// 	return value
	// }
	return value
}

// GetString 字符串格式返回值
func (y *Config) GetString(keyName string) string {
	value := y.viper.GetString(keyName)
	return value
	// if y.keyIsCache(keyName) {
	// 	return y.getValueFromCache(keyName).(string)
	// } else {
	// 	value := y.viper.GetString(keyName)
	// 	y.cache(keyName, value)
	// 	return value
	// }
}

// GetBool 布尔格式返回值
func (y *Config) GetBool(keyName string) bool {
	value := y.viper.GetBool(keyName)
	return value
}

// GetInt 整数格式返回值
func (y *Config) GetInt(keyName string) int {
	value := y.viper.GetInt(keyName)
	return value
}

// GetInt32 整数格式返回值
func (y *Config) GetInt32(keyName string) int32 {
	value := y.viper.GetInt32(keyName)
	return value
}

// GetInt64 整数格式返回值
func (y *Config) GetInt64(keyName string) int64 {
	value := y.viper.GetInt64(keyName)
	return value
}

// GetFloat64 小数格式返回值
func (y *Config) GetFloat64(keyName string) float64 {
	value := y.viper.GetFloat64(keyName)
	return value
}

// func (y *Config) GetFloat32(keyName string) float64 {
// 	// value := y.viper.GetFloat32(keyName)
// 	value := y.viper.GetFloat64(keyName)
// 	return value
// }
// GetDuration 时间单位格式返回值
func (y *Config) GetDuration(keyName string) time.Duration {
	value := y.viper.GetDuration(keyName)
	return value
}

// GetStringSlice 字符串切片数格式返回值
func (y *Config) GetStringSlice(keyName string) []string {
	value := y.viper.GetStringSlice(keyName)
	return value
}
