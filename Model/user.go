package Model

import (
	"fmt"
	"go/global"
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	*gorm.DB `gorm:"-" json:"-"`
	Id       int64 `gorm:"primaryKey" json:"id"`
	// CreatedAt int64 `gorm:"column:created_at;type:datetime"` //日期时间字段统一设置为字符串即可  time.Time
	// UpdatedAt int64 `gorm:"column:updated_at;type:datetime"`
	// DeletedAt gorm.DeletedAt `json:"deleted_at"` // 如果开发者需要使用软删除功能，打开本行注释掉的代码即可，同时需要在数据库的所有表增加字段deleted_at 类型为 datetime
}

// UserInfo 用户信息
type UserInfo struct {
	BaseModel
	Name         string //名称
	Account      string `gorm:"unique;not null"` //账户
	Email        string //邮箱
	Age          string //年龄
	Password     string //密码
	Portrait     string //头像
	Role         string //权限
	Phone        string `gorm:"unique;not null"` //手机号
	State        int    //状态，0-未激活，1-审核中，2-审核未通过，3-已审核
	OrganizatiId int    `gorm:"comment:组织Id" json:"organizatiId"`
}

func UseDbConn(sqlType string) *gorm.DB {
	var db *gorm.DB
	return db
}
func CreateUserFactory(sqlType string) *UserInfo {
	return &UserInfo{BaseModel: BaseModel{DB: UseDbConn(sqlType)}}
}

type Data struct {
	code  int
	rule  string
	token string
}

//修改用户信息 name,password,phone,email
func PostUser(ID int64, name string, password string, phone string, email string, portrait string) bool {
	// var u = new(UserInfo)
	// Db.Update(   )
	// Db.Where("id = ?",ID)
	sql := "UPDATE user_info SET name = ?,password = ?,phone = ?,email = ?,portrait = ? WHERE id = ?"
	result := Db.Exec(sql, name, password, phone, email, portrait, ID)
	fmt.Println("修改状态是")
	fmt.Println(result.RowsAffected)
	if result.RowsAffected > 0 {
		return true
	} else {
		return false
	}
}

//UPDATE user_info SET name = 23,password = 3,phone = 3,email = 3 WHERE id = 3

// func GetUser(id int64) *UserInfo {
// 	var u = new(UserInfo)
// 	sql := "select * from user_info where  id=?"
// 	Db.Raw(sql, id).Scan(u)

// 	var uu UserInfo
// 	Db.First(&uu)

// 	return u
// }

func GetUser(id int64) *UserInfo {
	var u UserInfo
	sql := "select * from user_info where  id=?"
	Db.Raw(sql, id).Scan(&u)

	// var uu UserInfo
	// Db.First(&uu)

	return &u
}

type WorkFlowResp struct {
	Items []*UserInfo `json:"items"`
	Total int64       `json:"total"`
}

func GetUsers(UserQuery *global.UserQuery) ([]*UserInfo, int64) {
	var list []*UserInfo
	var total int64
	fmt.Println("id是")
	// var OrganizatiId interface{}
	query := Db.Where("1 = 1")
	if UserQuery.OrganizatiId != 0 {
		query = query.Where("organizatiId = ?", UserQuery.OrganizatiId)
	}
	if UserQuery.Keywords != "" {
		query = query.Where("name LIKE ?", "%"+UserQuery.Keywords+"%")
	}
	query.Offset((UserQuery.PageNum - 1) * 10).Limit(UserQuery.PageSize).Find(&list)
	query.Model(&UserInfo{}).Offset((UserQuery.PageNum - 1) * 10).Limit(UserQuery.PageSize).Count(&total)
	// Db.Where("name LIKE ? AND organizatiId = ?", "%"+UserQuery.Keywords+"%", OrganizatiId).Offset((UserQuery.PageNum - 1) * 10).Limit(UserQuery.PageSize).Find(&list)
	// Db.Where("name LIKE ?", "%"+UserQuery.Keywords+"%").Offset((UserQuery.PageNum - 1) * 10).Limit(UserQuery.PageSize).Count(&total)
	// Db.Find(&list)
	// Db.Model(&UserInfo{}).Where("name LIKE ?", "%"+UserQuery.Keywords+"%").Offset((UserQuery.PageNum - 1) * 10).Limit(UserQuery.PageSize).Count(&total)
	// err := Db.Model(&UserInfo{}).Count(&total)

	// fmt.Print(err)

	return list, int64(total)
}

// Model.UserInfo (data interface{}, err error)
func Login(name string, Password string) *UserInfo {
	fmt.Println("进入查询")
	var u = new(UserInfo)
	sql := "select * from user_info where  account=? and password=?"
	Db.Raw(sql, name, Password).Scan(u)
	if u.Name == "" {
		fmt.Println("返回空数据")
		return u
	} else {
		if u.Password == Password {
			fmt.Println("密码正确")
			// token, _ := My_token.GenerateToken(u.Id, u.Name, u.Phone, 28800)
			// data = gin.H{
			// 	"code":  200,
			// 	"rule":  u.Role,
			// 	"token": token,
			// }
			return u
		} else {
			fmt.Println("密码错误")
		}
		fmt.Println("有数据")
	}
	fmt.Println(u)
	return u
	//
	// data := db.Raw(sql, name).First(u)
}

//用户注册
func Register(name string, password string, phone string, email string, account string) bool {
	// u := new(UserInfo)
	user := &UserInfo{Name: name, Email: email, Password: password, Account: account, Phone: phone}
	err1 := Db.Create(user).Error // 成功创建
	if err1 != nil {
		return false
	}
	return true
	// sql := "INSERT  INTO user_info(name,password,phone,email,account) SELECT ?,?,?,? FROM DUAL   WHERE NOT EXISTS (SELECT 1  FROM user_info WHERE  name=?)"
	// result := Db.Exec(sql, name, password, phone, email, name)
	// if result.RowsAffected > 0 {
	// 	return true
	// } else {
	// 	return false
	// }
	// fmt.Println(result)
	// fmt.Println(result.RowsAffected)
}

// Problem 用户问题表
type Problem struct {
	*gorm.DB   `gorm:"-" json:"-"`
	Id         int64 `gorm:"primaryKey" json:"id"`
	Created_at int64
	Content    string
	Account_id int64
	User       *UserInfo `gorm:"foreignKey:Account_id;references:Id"`
}

func (u *Problem) BeforeCreate(tx *gorm.DB) (err error) {
	u.Created_at = time.Now().Unix()
	return
}

//获取用户问题存数据库
func Streamopai(content string, Account_id int64) bool {
	// Db.AutoMigrate(&Problem{})
	Messages := Problem{Content: content, Account_id: Account_id}
	result := Db.Save(&Messages)
	// Db.Model().Association()
	if result.RowsAffected > 0 {
		return true
	} else {
		return false
	}
}

//获取用户问题存数据库
func Getuserli() []Problem {
	// sss := []User{}
	// Db.Preload("CreditCard").Find(&sss)
	// sss := []CreditCard{}
	// Db.Preload("User").Find(&sss)
	sss := []Problem{}
	Db.Preload("User").Find(&sss)
	return sss
}

func Daddd(id int64) []Problem {
	list := []Problem{}
	Db.Where("account_id=? and created_at > ?", id, "2023-05-16").Find(&list)
	return list
}

//获取所有组织
func GetOrganizationlist() []global.Organization {
	list := []global.Organization{}
	Db.Find(&list)
	return list
}

// User 有多张 CreditCard，UserID 是外键
type User struct {
	BaseModel
	CreditCardId int
	// CreditCard   CreditCard
}

type CreditCard struct {
	BaseModel
	Number string
	Title  string
	User   []User `gorm:"foreignKey:CreditCardId;references:Id"`
}

//初始化用户表
func InitUsers() {
	// Db.AutoMigrate(&UserInfo{})
	// fmt.Println("创建数据库报错为")
	// Db.Create(&global.Organization{Name: "技术部"})
	// Db.Create(&global.Organization{Name: "人事部"})
	// Db.Create(&global.Organization{Name: "后勤部"})
	// Db.Create(&UserInfo{Name: "史尿多", Account: "admin", Password: "123"})
	// Db.AutoMigrate(&UserInfo{})
	// var err = Db.AutoMigrate(&User{}, &CreditCard{})
	// // var err = Db.Model(&SkuProperty{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT") //Db.AutoMigrate(&UserInfo{}, &Sku_property{})
	// if err != nil {
	// 	fmt.Println(err.Error())
	// } else {
	// 	fmt.Println("2")
	// }
}
