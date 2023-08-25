package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(Cors())
	r.GET("/stream", func(c *gin.Context) {
		//设置responseheader
		c.Header("Content-Type", "text/event-stream")

		c.Header("Cache-Control", "no-cache")
		c.Header("Connection", "keep-alive")
		c.SSEvent("message", "1")
		c.SSEvent("message", "1")
		c.SSEvent("message", "1")
		c.SSEvent("message", "1")
		c.SSEvent("message", "1")
		c.SSEvent("message", "1")
		c.SSEvent("message", "1")
		c.SSEvent("message", "1")
		c.SSEvent("message", "1")
		c.SSEvent("message", "1")
		c.Done()
		//定时器，每隔2秒发送一次数据
		// ticker := time.NewTicker(2 * time.Second)
		// defer ticker.Stop()
		// //发送数据给前端
		// for {
		// 	select {
		// 	case <-ticker.C:
		// 		//模拟数据
		// 		data := fmt.Sprintf("data:%v\n\n", time.Now().String())
		// 		//发送数据给前端
		// 		c.SSEvent("message", data)
		// 		fmt.Println(data)
		// 	case <-c.Done():
		// 		//请求中断，跳出循环
		// 		return
		// 	}
		// }
	})
	r.Run(":8083")
	// http.HandleFunc("/stream", streamHandler)
	// http.ListenAndServe(":8082", nil)
}

func streamHandler(w http.ResponseWriter, c *gin.Context) {
	//设置responseheader
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	//定时器，每隔2秒发送一次数据
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	//循环发送数据
	for {
		select {
		case <-ticker.C:
			//模拟数据
			data := fmt.Sprintf("data:%v\n\n", time.Now().String())
			//发送数据给前端
			_, err := w.Write([]byte(data))
			c.SSEvent("message", "开始")
			fmt.Println(data)
			if err != nil {
				//发送失败，跳出循环
				return
			}
		case <-c.Done():
			//请求中断，跳出循环
			return
		}
	}
}

// 定义全局的CORS中间件
func Cors() gin.HandlerFunc {
	fmt.Println("请求中间件")
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin") //请求头部
		if origin != "" {
			//接收客户端发送的origin （重要！）
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			// c.Header("Access-Control-Allow-Origin", "*")
			//服务器支持的所有跨域请求的方法
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
			//允许跨域设置可以返回其他子段，可以自定义字段
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Access-Token,session,Content-Type")
			// 允许浏览器（客户端）可以解析的头部 （重要）
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
			//设置缓存时间
			c.Header("Access-Control-Max-Age", "172800")
			//允许客户端传递校验信息比如 cookie (重要)
			c.Header("Access-Control-Allow-Credentials", "true")
		}

		//允许类型校验
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "ok!")
		}

		defer func() {
			if err := recover(); err != nil {
				// log.Printf("Panic info is: %v", err)
			}
		}()

		c.Next()
	}
}
