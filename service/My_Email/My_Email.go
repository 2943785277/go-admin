package My_Email

import (
	"go/global"
	"log"
	"net/smtp"

	"github.com/jordan-wright/email"
)

func SendEmail(From string, To []string, Subject string, Text []byte) bool {
	e := email.NewEmail()
	//设置发送方的邮箱
	e.From = From //"dj <2943785277@qq.com>"
	// 设置接收方的邮箱
	e.To = To //[]string{"2943785277@qq.com"}
	//设置主题
	e.Subject = Subject //"这是主题"
	//设置文件发送的内容
	e.Text = Text //[]byte("www.topgoer.com是个不错的go语言中文文档")
	//设置服务器相关的配置
	err := e.Send(global.ConfigYml.GetString("email.address"), smtp.PlainAuth(global.ConfigYml.GetString("email.identity"), global.ConfigYml.GetString("email.username"), global.ConfigYml.GetString("email.password"), global.ConfigYml.GetString("email.host")))
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}
