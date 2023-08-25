package Utils

import (
	"fmt"
	"go/global"
	"io"
	"mime/multipart"
	"os"
	"time"
)

/*
	上传文件
	file   传入图片数据
	Name  文件夹名
	FileName  文件名
	Type  类型

	返回值  静态资源地址
*/
func Upload(file multipart.File, Folder string, FileName string, Type string) (string, bool) {
	Time := time.Now().Format("20060102")
	FolderName := Folder + Time
	path := FolderName + "/" + FileName + Type
	newpath := Time + "/" + FileName + Type
	if _, err := os.Stat(FolderName); os.IsNotExist(err) {
		err = os.MkdirAll(FolderName, 0755)
		if err != nil {
			fmt.Println("Error creating directory:", err)
			return "", false
		}
	}

	// 文件夹已经存在或者新建成功了，现在可以写入文件了。
	newfile, err := os.Create(path)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return "", false
	}
	defer file.Close()
	_, err = io.Copy(newfile, file)
	if err != nil {
		fmt.Println("打开报错了")
		fmt.Println(err)
	}
	var path2 = global.ConfigYml.GetString("app.app_url") + ":" + global.ConfigYml.GetString("app.port") + global.ConfigYml.GetString("FileUploadSetting.UploadFileReturnPath") + newpath
	fmt.Println("图片路径返回值")
	fmt.Println(path)
	fmt.Println(path2)
	return path2, true
}
