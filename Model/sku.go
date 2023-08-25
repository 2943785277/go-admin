package Model

import (
	"fmt"
)

// 商品信息
type Sku_property struct {
	BaseModel
	Name       string    //商品名称
	Title      string    //商品标题
	Sketch     string    //商品描述
	Production string    //生产企业
	Code       string    //商品唯一编号
	Price      float64   //商品当前价格
	Sub_title  string    //商品副标题
	Color      string    //商品颜色
	State      int       //商品状态，0-有效，1-无效，2-锁定
	User       *UserInfo `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"` //用户信息
}

//添加商品
func Addshop(Data Sku_property) bool {
	fmt.Println("进入Addshop")
	// now := time.Now() //.Format("2006-01-02 15:04:05.00")
	// Data.CreatedAt = now.Unix()
	// Data.UpdatedAt = now.Unix()
	// fmt.Println(Data.UpdatedAt)
	// fmt.Println(Data)
	// Data.Name = "asdas6666"
	// fmt.Println(&Data)
	// result := Db.Save(&Data)
	sql := "INSERT INTO `sku_property` (`created_at`,`updated_at`,`name`,`title`,`sketch`,`production`,`code`,`price`,`sub_title`,`color`,`state`) VALUES (1679284138,1679284138,'70.62660472426823','','','','',0.000000,'','',0)"
	result := Db.Exec(sql)
	if result.RowsAffected > 0 {
		return true
	} else {
		return false
	}
	// tx := Db.Begin()
	// if err := tx.Save(&Data).Error; err != nil {
	// 	// 错误处理
	// 	tx.Rollback()
	// 	fmt.Println(err)
	// }

	// sql := "UPDATE sku_property SET name = ? WHERE id = ?"
	// if err := tx.Exec(sql, "二次修改", Data.Id).Error; err != nil {
	// 	// 错误处理
	// 	tx.Rollback()
	// 	fmt.Println(err)
	// }
	// if berr := tx.Commit().Error; berr != nil {
	// 	return false
	// } else {
	// 	return true
	// }
	// result := Db.Debug().Save(Data)
	// fmt.Println(result)

	// INSERT INTO `sku_property` (`created_at`,`updated_at`,`name`,`title`,`desc`,`production`,`code`,`price`,`sub_title`,`color`,`state`) VALUES (0,0,'啊实打实','胜多负少的方式','胜多负少的','沙发上','',453.000000,'23423','水电费水电费',1)

}

//获取商品列表
func GetshopList() ([]*Sku_property, int64) {
	var List []*Sku_property
	var total int64
	Db.Find(&List)
	// sql := "select * from sku_property"
	// Db.Raw(sql).Scan(&List)
	for i, v := range List {
		fmt.Println(i)
		fmt.Println(v.Name)
	}
	return List, int64(total)
}

//
func GetShop(id int) *Sku_property {
	var data = new(Sku_property)
	sql := "select * from sku_property where  id=?"
	Db.Raw(sql, id).Scan(data)
	return data
}

// UPDATE sku_property SET name = '修改',title = '223','dketch' = '',code = '',price = 0.000000,sub_title = '23',color = '234',state = 0 WHERE id = 32
// UPDATE sku_property SET name = '商品名称',title = '223','dketch' = '胜多负少的',code = '',price = 6.000000,sub_title = '23',color = '234',state = 0 WHERE id = 32
//修改商品信息
func PutShop(Data Sku_property) bool {
	sql := "UPDATE sku_property SET name = ?,title = ?,sketch = ?,code = ?,price = ?,sub_title = ?,color = ?,state = ? WHERE id = ?"
	result := Db.Exec(sql, Data.Name, Data.Title, Data.Sketch, Data.Code, Data.Price, Data.Sub_title, Data.Color, Data.State, Data.Id)
	fmt.Println("修改状态是")
	fmt.Println(result.RowsAffected)
	if result.RowsAffected > 0 {
		return true
	} else {
		return false
	}
}
