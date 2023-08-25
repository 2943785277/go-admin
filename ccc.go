package main

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1"
)

func da() {
	cfg, err = ini.Load("./config/config.ini")
	// fmt.Println(cfg)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("userName ==>", cfg.Section("pas").Key("pwd").String())
}

func main() {

	// 初始化Gin框架
	r := gin.Default()

	// 初始化session中间件
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("sessionName", store))

	// 登录路由
	r.POST("/api/User/login", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")

		// 假设在登录验证成功后，设置session状态为已登录
		session := sessions.Default(c)
		session.Set("isLoggedIn", true)
		session.Save()

		c.JSON(http.StatusOK, gin.H{
			"message":  "登录成功",
			"username": username,
			"password": password,
		})
	})

	// 需要验证登录状态的路由
	r.GET("/api/auth/captcha", func(c *gin.Context) {
		session := sessions.Default(c)
		isLoggedIn := session.Get("isLoggedIn")

		// 检查session中的isLoggedIn状态
		if isLoggedIn == true {
			c.JSON(http.StatusOK, gin.H{
				"message": "已经登录",
			})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "未登录",
			})
		}
	})

	// 启动服务器
	r.Run(":8090")
}
