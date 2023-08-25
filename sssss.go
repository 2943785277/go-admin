package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sashabaranov/go-openai"
)

//读取配置txt
func Getconfig() {
	file, err := os.Open("config.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// 读取文件中的配置
	scanner := bufio.NewScanner(file)
	Configs = make(map[string]interface{})
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "=")
		if len(parts) != 2 {
			continue
		}
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		Configs[key] = value
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

var Configs map[string]interface{}

type Messages struct {
	Content string
	Role    string
	Name    string
}
type openaidata struct {
	Key      string `json:"key"`
	Text     string `json:"text"`
	Model    string `json:"model"`
	Messages []openai.ChatCompletionMessage
}

func main() {
	Getconfig()
	fmt.Println("加载open开始运行")
	r := gin.Default()
	r.Use(Cors())
	r.POST("/api/opai", func(c *gin.Context) {
		var u openaidata
		erri := c.ShouldBind(&u)
		fmt.Printf("Completion error: %v\n", erri)
		fmt.Println("获取信息是下列")
		fmt.Println(u)
		fmt.Println("获取的基础配置项是")
		fmt.Println(Configs)
		// opai := openai.NewClient(Configs["Key"].(string))
		// ctx := context.Background()

		MaxTokens, _ := Configs["MaxTokens"].(int)
		Temperature, _ := Configs["Temperature"].(float32) //断言，转换数据格式,
		// openai.GPT3Dot5Turbo   gpt-3.5-turbo
		client := openai.NewClient(Configs["Key"].(string))
		resp, err := client.CreateChatCompletion(
			context.Background(),
			openai.ChatCompletionRequest{
				Model:       u.Model,
				MaxTokens:   MaxTokens,
				Temperature: Temperature, //断言，转换数据格式,
				Messages: []openai.ChatCompletionMessage{
					{
						Role:    openai.ChatMessageRoleUser,
						Content: u.Text,
					},
				},
				Stream: true,
			},
		)
		if err != nil {
			c.String(http.StatusOK, "Completion error: %v\n", err)
			fmt.Printf("ChatCompletion error: %v\n", err)
			return
		}

		c.String(http.StatusOK, resp.Choices[0].Message.Content)
	})
	r.GET("/api/Streamopai/:Text/:Model", func(c *gin.Context) {
		// var u openaidata
		// erri := c.ShouldBind(&u)
		// fmt.Printf("Completion error: %v\n", erri)
		Text := c.Param("Text")
		Model := c.Param("Model")
		opai := openai.NewClient(Configs["Key"].(string))
		MaxTokens, _ := Configs["MaxTokens"].(int)
		Temperature, _ := Configs["Temperature"].(float32) //断言，转换数据格式,
		ctx := context.Background()
		req := openai.ChatCompletionRequest{
			Model:       Model,
			MaxTokens:   MaxTokens,
			Temperature: Temperature, //断言，转换数据格式,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: Text,
				},
			},
			Stream: true,
		}
		flusher, ok := c.Writer.(http.Flusher)
		if !ok {
			c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("Streaming unsupported"))
			return
		}
		stream, err := opai.CreateChatCompletionStream(ctx, req)
		if err != nil {
			fmt.Printf("ChatCompletionStream error: %v\n", err)
			return
		}
		defer stream.Close()
		c.Header("Transfer-Encoding", "chunked")
		c.SSEvent("message", "开始")
		fmt.Printf("Stream response: ")
		for {
			response, err := stream.Recv()
			if errors.Is(err, io.EOF) {
				fmt.Println("\nStream finished")
				c.SSEvent("stop", "finish")
				flusher.Flush() // 手动刷新响应流
				return
			}
			if err != nil {
				fmt.Printf("\nStream error: %v\n", err)
				c.SSEvent("error", "error")
				flusher.Flush() // 手动刷新响应流
				return
			}
			c.SSEvent("message", response.Choices[0].Delta.Content)
			fmt.Printf(response.Choices[0].Delta.Content)
			flusher.Flush() // 手动刷新响应流
		}

	})

	r.POST("/api/Streamopai", func(c *gin.Context) {
		data := []byte("example data") // 这里替换为实际的数据

		c.Header("Content-Description", "File Transfer")
		c.Header("Content-Disposition", "attachment; filename=example.dat")
		c.Data(http.StatusOK, "application/octet-stream", data)
		// http.ResponseWriter.WriteHeader(http.StatusOK)
		// var u openaidata
		// erri := c.ShouldBind(&u)
		// fmt.Println(u)
		// // var num = len(u.Messages)
		// // fmt.Println(u.Messages[num-1])
		// fmt.Printf("Completion error: %v\n", erri)
		// c.String(http.StatusOK, u.Text)
		// return
		// // Text := c.Param("Text")
		// Model := u.Model
		// opai := openai.NewClient(Configs["Key"].(string))
		// MaxTokens, _ := Configs["MaxTokens"].(int)
		// Temperature, _ := Configs["Temperature"].(float32) //断言，转换数据格式,
		// ctx := context.Background()
		// req := openai.ChatCompletionRequest{
		// 	Model:       Model,
		// 	MaxTokens:   MaxTokens,
		// 	Temperature: Temperature, //断言，转换数据格式,
		// 	Messages:    u.Messages,
		// 	Stream:      true,
		// }
		// flusher, ok := c.Writer.(http.Flusher)
		// if !ok {
		// 	c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("Streaming unsupported"))
		// 	return
		// }
		// stream, err := opai.CreateChatCompletionStream(ctx, req)
		// if err != nil {
		// 	fmt.Printf("ChatCompletionStream error: %v\n", err)
		// 	return
		// }
		// defer stream.Close()
		// c.Header("Transfer-Encoding", "chunked")
		// c.SSEvent("message", "开始")
		// fmt.Printf("Stream response: ")

		//创建channel，用于生成的文本返回前端
		// 	chuncks:=make(chanllbyte)
		// 	//创建goroutine，不断检查任务状态
		// 	go func() {
		// 		for {
		// 				task, err = client.Completions.Get(context.Background(), task.ID)
		// 				if err != nil {
		// 						http.Error(w, err.Error(), http.StatusInternalServerError)
		// 						return
		// 				}
		// 				if task.Status == openai.CompletionStatusComplete {
		// 						// 任务完成，发送生成的文本
		// 						chuncks <- []byte(task.Choices[0].Text)
		// 						close(chuncks)
		// 						return
		// 				}
		// 				time.Sleep(1 * time.Second)
		// 		}
		// }()

		// for {
		// 	response, err := stream.Recv()
		// 	if errors.Is(err, io.EOF) {
		// 		fmt.Println("\nStream finished")
		// 		c.SSEvent("stop", "finish")
		// 		flusher.Flush() // 手动刷新响应流
		// 		return
		// 	}
		// 	if err != nil {
		// 		fmt.Printf("\nStream error: %v\n", err)
		// 		c.SSEvent("error", "error")
		// 		flusher.Flush() // 手动刷新响应流
		// 		return
		// 	}
		// 	c.SSEvent("message", response.Choices[0].Delta.Content)
		// 	fmt.Printf(response.Choices[0].Delta.Content)
		// 	flusher.Flush() // 手动刷新响应流
		// }

	})
	r.GET("/api/a/:data", func(c *gin.Context) {
		captchaId := c.Param("data")
		fmt.Println(captchaId)
		c.String(http.StatusOK, captchaId)
	})
	r.POST("/api/post", func(c *gin.Context) {
		c.Header("Transfer-Encoding", "chunked")
		c.Header("Content-Type", "text/plain")
		//写入分块数据
		//循环发送chunked数据
		for i := 0; i < 10; i++ {
			message := fmt.Sprintf("[%d]ChunkedData\n", i+1)
			chunk := fmt.Sprintf("%x\r\n%s\r\n", len(message), message)
			c.Writer.Write([]byte(chunk))
			c.Writer.Flush()
			time.Sleep(1 * time.Second)
		}
		//发送结束协议
		message := "0\r\n\r\n"
		c.Writer.Write([]byte(message))
	})
	//监听端口默认为8080
	r.Run(":8090")
}
