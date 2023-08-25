package Model

import (
	"fmt"

	"go/global"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// 定义db全局变量
var Db *gorm.DB

func Init() {
	var err error
	ConfigYml := global.ConfigYml
	dsn := ConfigYml.GetString("database.username") + ":" + ConfigYml.GetString("database.password") + "@(" + ConfigYml.GetString("database.host") + ":" + ConfigYml.GetString("database.port") + ")/" +
		ConfigYml.GetString("database.database") + "?charset=" + ConfigYml.GetString("database.charset") + "&parseTime=True&loc=Local"
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: false,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 禁用表名加s
		},
		Logger:                                   logger.Default.LogMode(logger.Info), // 打印sql语句
		DisableAutomaticPing:                     false,
		DisableForeignKeyConstraintWhenMigrating: true, // 禁用创建外键约束
	})
	if err != nil {
		panic("Connecting database failed: " + err.Error())
	}

	//注册Hook
	// Db.Callback().Create().Before("gorm:create").Register("set_created_at", func(db *gorm.DB) { //tx*gorm.DB
	// timeFieldsToInit := []string{"CreateTime", "UpdateTime"}
	// tx.ha
	// for_,v:= tx.Statement.Schema.Fields{
	// 	fmt.Println()
	// 		if v.DBName == "CreatedAt" {
	// 				tx.Statement.SetColumn(v.DBName,time.Now().Unix())
	// 		}
	// }
	// })
	// Getmenulist()
	InitUsers()
	fmt.Println("SQL连接成功")
}

//GetDB
func GetDB() *gorm.DB {
	return Db
}
