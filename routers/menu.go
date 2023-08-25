package routers

import (
	"fmt"
	"go/Model"
	"go/global"
	"go/service/response"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gin-gonic/gin"
)

func Routes(ctx *gin.Context) {
	list, err := Model.Routes()
	if err {
		response.Success(ctx, gin.H{
			"data": list,
		}, "添加成功")
	} else {
		response.Error(ctx, gin.H{}, "获取失败请重试")
	}
}

//获取
func Getmenulist(ctx *gin.Context) {
	list, err := Model.Getmenulist()
	if err {
		response.Success(ctx, gin.H{
			"data": list,
		}, "添加成功")
	} else {
		response.Error(ctx, gin.H{}, "获取失败请重试")
	}
}

//获取menu  ID
func Getmenu(ctx *gin.Context) {
	ID, _ := strconv.ParseInt(ctx.Param("ID"), 10, 64)
	var data, err = Model.Getmenu(ID)
	if err {
		response.Success(ctx, data, "获取成功")
	} else {
		response.Error(ctx, gin.H{}, "获取失败，请重试")
	}
}

//添加
func Addmenu(ctx *gin.Context) {
	var u global.Menus
	ctx.ShouldBind(&u)
	if Model.Addmenu(u) {
		response.Success(ctx, gin.H{}, "添加成功")
	} else {
		response.Error(ctx, gin.H{}, "添加失败，请重试")
	}
}

func Editmenu(ctx *gin.Context) {
	var u global.Menus
	ctx.ShouldBind(&u)
	if Model.Editmenu(u) {
		response.Success(ctx, gin.H{}, "更新成功")
	} else {
		response.Error(ctx, gin.H{}, "更新失败，请重试")
	}
}
func Delmenu(ctx *gin.Context) {
	ID, _ := strconv.ParseInt(ctx.Param("ID"), 10, 64)
	if Model.Delmenu(ID) {
		response.Success(ctx, gin.H{}, "删除成功")
	} else {
		response.Error(ctx, gin.H{}, "添加失败，请重试")
	}
}

//导出文件
func Export(ctx *gin.Context) {
	list, err := Model.Getmenulist()
	if err {

		// 创建临时文件
		file, err := ioutil.TempFile("", "temp-*.xlsx")
		f := excelize.NewFile()
		// Create a new sheet.
		Time := time.Now().Unix()
		Fname := strconv.FormatInt(Time, 10) + "Book.xlsx"
		fmt.Println(Fname)
		sheetName := "Sheet1"
		index := f.NewSheet(sheetName)
		//设置表头
		f.SetCellValue(sheetName, "A1", "ID")
		f.SetCellValue(sheetName, "B1", "标题")
		f.SetCellValue(sheetName, "C1", "路由")
		for i, v := range list {
			f.SetCellValue(sheetName, "A"+strconv.Itoa(i+2), v.Id)
			f.SetCellValue(sheetName, "B"+strconv.Itoa(i+2), v.Title)
			f.SetCellValue(sheetName, "C"+strconv.Itoa(i+2), v.Path)
		}
		f.SetActiveSheet(index)
		// Save spreadsheet by the given path.
		if err := f.SaveAs(file.Name()); err != nil { //Fname
			fmt.Println(err)
			fmt.Println("写入失败")
		}
		// 保存文件
		if err := f.SaveAs(file.Name()); err != nil {
			log.Fatal(err)
		}

		// 关闭文件
		if err := file.Close(); err != nil {
			log.Fatal(err)
		}

		// 打开临时文件
		file, err = os.Open(file.Name())
		if err != nil {
			log.Fatal(err)
		}
		fileInfo, err := os.Stat(file.Name())
		aa := fmt.Sprintf("%d", fileInfo.Size())
		// 设置http响应头
		ctx.Writer.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", "temp.xlsx"))
		ctx.Writer.Header().Set("Content-Type", "application/octet-stream")
		ctx.Writer.Header().Set("Content-Length", aa)
		fmt.Println("dao zhebu l")
		// 将文件内容写入http响应
		if _, err := io.Copy(ctx.Writer, file); err != nil {
			log.Fatal(err)
		}
		// 关闭临时文件
		if err := file.Close(); err != nil {
			log.Fatal(err)
		}

		// 删除临时文件
		if err := os.Remove(file.Name()); err != nil {
			log.Fatal(err)
		}
		// response.Success(ctx, gin.H{
		// 	"data": list,
		// }, "添加成功")
	} else {
		response.Error(ctx, gin.H{}, "获取失败请重试")
	}
}

func LoadMenu(e *gin.Engine) {
	e.GET("/api/menu/Routes", Routes)
	e.GET("/api/menu/Getmenulist", Getmenulist)
	e.GET("/api/menu/Getmenu/:ID", Getmenu)
	e.POST("/api/menu/Addmenu", Addmenu)
	e.PUT("/api/menu/Editmenu", Editmenu)
	e.DELETE("/api/menu/Delmenu/:ID", Delmenu)
	e.GET("/api/menu/Export", Export)
}
