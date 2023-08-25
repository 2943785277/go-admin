package global

import (
	"go/config"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type BaseModel struct {
	*gorm.DB  `gorm:"-" json:"-"`
	Id        int64 `gorm:"primaryKey" json:"id"`
	CreatedAt int64 `gorm:"column:created_at"` //日期时间字段统一设置为字符串即可  time.Time
	UpdatedAt int64 `gorm:"column:updated_at"`
	// DeletedAt gorm.DeletedAt `json:"deleted_at"` // 如果开发者需要使用软删除功能，打开本行注释掉的代码即可，同时需要在数据库的所有表增加字段deleted_at 类型为 datetime
}

type YmlConfigInterf interface {
	Get(keyName string) interface{}
	GetString(keyName string) string
	GetBool(keyName string) bool
	GetInt(keyName string) int
	GetInt32(keyName string) int32
	GetInt64(keyName string) int64
	GetFloat64(keyName string) float64
	GetDuration(keyName string) time.Duration
	GetStringSlice(keyName string) []string
}

type Application struct {
	ConfigViper *viper.Viper
	Config      config.Configuration
}

var App = new(Application)

// 配置
var (
	ConfigYml YmlConfigInterf
	Redis     *redis.Client
)

// //路由权限
// type Routes struct {
// 	ParentId  int
// 	Id        int
// 	Name      string
// 	Path      string
// 	Redirect  string
// 	Meta      Meta
// 	Component string
// 	Children  []Routes
// }

// //
// type Meta struct {
// 	Icon  string
// 	Title string
// }
