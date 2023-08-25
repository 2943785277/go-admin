package routers

import (
	"fmt"
	"go/Model"
	"go/service/response"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

//添加商品
func Addshop(ctx *gin.Context) {
	var FormData Model.Sku_property
	ctx.ShouldBind(&FormData)
	time.Sleep(3 * time.Second)
	if Model.Addshop(FormData) {
		response.Success(ctx, gin.H{}, "添加成功")
	} else {
		response.Success(ctx, gin.H{}, "添加失败，请重试")
	}
}

//获取商品列表
func GetshopList(ctx *gin.Context) {
	list, Total := Model.GetshopList()
	data := gin.H{
		"item":  list,
		"total": Total,
	}
	response.Success(ctx, data, "")
}

//获取商品列表
func GetShop(ctx *gin.Context) {
	Id := ctx.Param("ID")
	sID, _ := strconv.Atoi(Id)
	fmt.Println(Id)
	fmt.Println("id是  -----------------")
	fmt.Println(sID)
	list := Model.GetShop(sID)
	data := gin.H{
		"data": list,
	}
	response.Success(ctx, data, "")
}

//修改商品
func PutShop(ctx *gin.Context) {
	var FormData Model.Sku_property
	ctx.ShouldBind(&FormData)
	fmt.Print(FormData.Id)
	if Model.PutShop(FormData) {
		response.Success(ctx, gin.H{}, "修改成功")
	} else {
		response.Success(ctx, gin.H{}, "修改失败，请重试")
	}
}

func LoadShop(e *gin.Engine) {
	e.POST("/shop/Addshop", Addshop)
	e.GET("/shop/GetshopList", GetshopList)
	e.GET("/shop/GetShop/:ID", GetShop)
	e.PUT("/shop/PutShop", PutShop)
}
