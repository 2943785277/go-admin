package main

import (
	"fmt"
	"go/Model"
	"go/service/My_token"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// UserInfo 用户信息
type UserInfo struct {
	Model.UserInfo
}

func UseDbConn(sqlType string) *gorm.DB {
	var db *gorm.DB
	return db
}

// func CreateUserFactory(sqlType string) *UserInfo {
// 	return &UserInfo{BaseModel: {DB: UseDbConn(sqlType)}}
// 	// return &UserInfo{DB: UseDbConn(sqlType)}
// }

// func (u *UserInfo) Getlogin(name string) (*UserInfo, error) {
func (u *UserInfo) Getaaa(name string) *UserInfo {
	// db, err := GetOneMysqlClient()
	// if err != nil {
	// 	panic(err)
	// }
	sql := "SELECT * FROM `user_infos`"
	u.Raw(sql).Scan(u)
	// sql := "select * from user_infos where  name=?  limit 1"
	// data := db.Raw(sql, name).First(u)

	// if data.Error != nil {
	// 	panic(err)
	// }
	return u
}

//连接mysql数据库
func GetOneMysqlClient() (*gorm.DB, error) {

	db, err := gorm.Open("mysql", "root:root@(127.0.0.1:3306)/cs?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	fmt.Println("sql连接成功")
	return db, nil
}
func Getsql(sqlType string) {

	// var u = new(Model.UserInfo)
	// sql := "SELECT * FROM `user_infos`"
	// db.Raw(sql).Scan(&u)
	// // db.Raw(sql).Scan(u)
	// fmt.Printf("%#v\n", u)
}
func main() {
	tok, _ := My_token.GenerateToken(1, "花卓", "180232", 1)
	fmt.Println("jwt是")
	fmt.Println(tok)
	time.After(10 * time.Second)
	_, datas, err := My_token.ParseToken("asdsa" + tok + "cc")
	fmt.Println(err)
	fmt.Println("jwt解析后是")
	fmt.Println(datas)
	//Model.Init()
	// Model.InitUsers()
	//Model.Register("花卓3", "1233", "2", "34")
	// Model.Login("花卓", "123")
	// data := Model.GetDB()

	// var u = new(Model.UserInfo)
	// sql := "SELECT * FROM `user_infos`"
	// data.Raw(sql).Scan(&u)
	// fmt.Printf("%#v\n", u)

	// GetOneMysqlClient()
	// Getsql("")
	// var u = new(Model.UserInfo)
	// sql := "SELECT * FROM `user_infos`"
	// db.Raw(sql).Scan(&u)
	// // db.Raw(sql).Scan(u)
	// fmt.Printf("%#v\n", u)
	// fmt.Println(db)
	// userModelFact := Model.CreateUserFactory("")
	// data := userModelFact.Getaaa("枯藤1")
	// fmt.Printf("%#v\n", data)

	// ----------------------------
	// data, err := Getlogin("枯藤1")
	// u1 := UserInfo{Name: "枯藤1", Gender: "1男", Hobby: "篮球1"}
	// db, err := GetOneMysqlClient()
	// db.AutoMigrate(&UserInfo{})
	// db.Create(&u1)
	// db, err := gorm.Open("mysql", "root:root@(127.0.0.1:3306)/cs?charset=utf8mb4&parseTime=True&loc=Local")
	// if err != nil {
	// 	panic(err)
	// }
	// db.AutoMigrate(&UserInfo{})
	// var u = new(Model.UserInfo)
	// // db.Where("Id = ?", 1).Take(u)
	// // fmt.Println(u.Name)
	// // fmt.Printf("%#v\n", u)
	// sql := "SELECT * FROM `user_infos`"
	// db.Raw(sql).Scan(u)
	// fmt.Printf("%#v\n", u)
	// db.Raw(sql).RecordNotFound{
	// 	fmt.Println("查询不到数据")
	// }
	// // db.Find(u, "name = ?", "枯藤")

	// result := u.Raw(sql, userName).First(u)
	// defer db.Close()

	// // 自动迁移
	// db.AutoMigrate(&UserInfo{})

	// u1 := UserInfo{Name: "枯藤", Gender: "男", Hobby: "篮球"}
	// u2 := UserInfo{Name: "topgoer.com", Gender: "女", Hobby: "足球"}
	// // 创建记录
	// db.Create(&u1)
	// db.Create(&u2)
	// // 查询

	// db.First(u)

	// var uu UserInfo
	// db.Find(&uu, "hobby=?", "足球")
	// fmt.Printf("%#v\n", uu)
	// // 更新
	// db.Model(&u).Update("hobby", "双色球")
	// // 删除
	// db.Delete(&u)
}
