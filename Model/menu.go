package Model

import (
	"fmt"
	"go/global"
)

//获取菜单Getmenulist
func Routes() ([]global.Routes, bool) {
	list := []global.Menus{}
	err := Db.Find(&list).Error
	a1 := []global.Routes{}
	for _, x := range list {
		if x.ParentId == 0 {
			a1 = append(a1, global.Routes{
				Menus: x,
				Meta: global.Meta{
					Title: x.Title,
					Icon:  x.Icon,
				},
			})
		}
	}
	for _, item := range list {
		for i, it := range a1 {
			if item.ParentId == it.Id {
				a1[i].Children = append(a1[i].Children, global.Routes{
					Menus: item,
					Meta: global.Meta{
						Title: item.Title,
						Icon:  item.Icon,
					},
				})
			}
		}
	}
	if err != nil {
		return a1, false
	}
	return a1, true
}
func Getmenulist() ([]global.Menus, bool) {
	list := []global.Menus{}
	err := Db.Find(&list).Error
	if err != nil {
		return list, false
	}
	return list, true
}

//新增菜单
func Addmenu(Data global.Menus) bool {
	err := Db.Create(&Data).Error
	if err != nil {
		return false
	}
	return true
}

//修改菜单
func Editmenu(Data global.Menus) bool {
	err := Db.Save(&Data).Error
	if err != nil {
		return false
	}
	return true
}

//删除菜单
func Delmenu(ID int64) bool {
	var u global.Menus
	err := Db.Where("id = ?", ID).Delete(&u).Error
	if err != nil {
		return false
	}
	return true
}
func Getmenu(ID int64) (global.Menus, bool) {
	var u global.Menus
	err := Db.First(&u, ID).Error
	if err != nil {
		fmt.Println(u.Title)
		return u, false
	}
	fmt.Println(err)
	return u, true
}
