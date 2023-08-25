package routers

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"go/Model"
	"go/Utils"
	"go/global"
	"go/middleware"
	"go/service/My_Captcha"
	"go/service/My_File"
	"go/service/My_token"
	"go/service/response"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"sync"
	"time"

	"github.com/dchest/captcha"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/sashabaranov/go-openai"
)

type Loginjson struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Captcha   string `json:"captcha"`
	CaptchaID string `json:"captchaID"`
}
type Users struct {
}

func Jsoa(data map[string]string) map[string]string {
	fmt.Println(data)
	return data
}

//上传图片
func Img(ctx *gin.Context) {
	var u Model.UserInfo
	ctx.ShouldBind(&u)
	fmt.Println("文件大小为：")
	fmt.Println(len(u.Portrait))

	re, _ := regexp.Compile(`^data:\s*image\/(\w+);base64,`)
	allData := re.FindAllSubmatch([]byte(u.Portrait), 2)
	fileType := string(allData[0][1]) //png ，jpeg 后缀获取
	fmt.Println(fileType)
	base64Str := re.ReplaceAllString(u.Portrait, "")
	byte, _ := base64.StdEncoding.DecodeString(base64Str)
	// fmt.Println(byte)
	date := time.Now().Format("20060102")

	if ok := My_File.IsFileExist(global.ConfigYml.GetString("FileUploadSetting.UploadFileSavePath") + date); !ok {
		os.Mkdir(global.ConfigYml.GetString("FileUploadSetting.UploadFileSavePath")+date, 0666)
	}
	FileName := date + "/" + strconv.FormatInt(time.Now().Unix(), 10) + strconv.Itoa(rand.Intn(999999-100000)+100000) + "." + fileType
	path := global.ConfigYml.GetString("FileUploadSetting.UploadFileSavePath") + FileName
	path2 := global.ConfigYml.GetString("app.app_url") + ":" + global.ConfigYml.GetString("app.port") + global.ConfigYml.GetString("FileUploadSetting.UploadFileReturnPath") + FileName
	err := ioutil.WriteFile(path, byte, 0666)
	if err != nil {
		fmt.Println("写入err")
	}
	fmt.Println("写入response.Success")
	response.Success(ctx, gin.H{
		"url": path2,
	}, "")
}

func IsFileExist(s string) {
	panic("unimplemented")
}

//登录 4
func Login(ctx *gin.Context) {

	// fmt.Println(session)
	var u Loginjson
	err := ctx.ShouldBind(&u)
	if err != nil {
		fmt.Println("账号未登录=eer====")
	}
	var Isopen = global.ConfigYml.GetString("jwt.isopen")
	Isopenbool, _ := strconv.ParseBool(Isopen)
	var UserInfo *Model.UserInfo
	// 检查用户是否已登录
	session := sessions.Default(ctx)
	// if auth, ok := session.Get(UserInfo.Id).(bool); !ok || !auth {
	// 	// 如果用户未通过身份验证，执行相关逻辑，如重定向到登录页
	// 	fmt.Println("账号已经登录了")
	// } else {
	// 	fmt.Println("账号未登录=====")
	// }

	fmt.Println(UserInfo.Id, "-----------ids -----")
	session.Set("name", true)
	session.Save()
	UserInfo = Model.Login(u.Username, u.Password)
	fmt.Println(UserInfo)
	fmt.Println("验证码状态：", Isopenbool)
	if Isopenbool {
		if captcha.VerifyString(u.CaptchaID, u.Captcha) {
			fmt.Println("验证成功")
			if UserInfo.Name == "" {
				data := gin.H{}
				response.Error(ctx, data, "账号或密码错误，请重试")
			} else {
				token, _ := My_token.GenerateToken(UserInfo.Id, UserInfo.Name, UserInfo.Phone, 28800)
				data := gin.H{
					"Code":  200,
					"rule":  UserInfo.Role,
					"token": token,
				}
				response.Success(ctx, data, "")
			}
		} else {
			response.Error(ctx, gin.H{}, "验证码失败")
		}
	} else {
		if UserInfo.Name == "" {
			data := gin.H{}
			response.Error(ctx, data, "账号或密码错误，请重试")
		} else {
			token, _ := My_token.GenerateToken(UserInfo.Id, UserInfo.Name, UserInfo.Phone, 28800)
			data := gin.H{
				"Code":  200,
				"rule":  UserInfo.Role,
				"token": token,
			}
			response.Success(ctx, data, "")
		}
	}

}

//获取路由
func GetRoute(ctx *gin.Context) {
	jsonFile, err := os.Open("data/Route.json")
	if err != nil {
		fmt.Println("error opening json file")
		return
	}
	defer jsonFile.Close()

	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("error reading json file")
		return
	}
	formData := make(map[string]interface{})
	json.Unmarshal(jsonData, &formData)
	response.Success(ctx, gin.H{"data": formData["Data"]}, "注册成功")
	// ctx.JSON(200, gin.H{"Code": 200, "Data": formData["Data"]})
}

//注册
func Register(ctx *gin.Context) {
	var u Model.UserInfo
	ctx.ShouldBind(&u)
	if Model.Register(u.Name, u.Password, u.Phone, u.Email, u.Account) {
		response.Success(ctx, gin.H{}, "注册成功")
	} else {
		response.Error(ctx, gin.H{}, "注册失败")
	}
}

//注销登录

func Exit(ctx *gin.Context) {
	tokenString := ctx.GetHeader("Authorization")
	_, claims, _ := My_token.ParseToken(tokenString)
	global.Redis.Set(ctx, tokenString, claims.StandardClaims.ExpiresAt, 10000*time.Second).Err()
	response.Success(ctx, gin.H{}, "退出成功")
}

func PostUser(ctx *gin.Context) {
	var u Model.UserInfo
	ctx.ShouldBind(&u)
	if Model.PostUser(u.Id, u.Name, u.Password, u.Phone, u.Email, u.Portrait) {
		response.Success(ctx, gin.H{"Code": 200}, "修改成功")
	} else {
		response.Error(ctx, gin.H{"Code": 400}, "修改失败，请重新修改")
	}
}

//获取用户信息
func GetUser(ctx *gin.Context) {
	tokenString := ctx.GetHeader("Authorization")
	fmt.Println("进入GetUser")
	fmt.Println(tokenString)
	// if tokenString == ""{
	// 	response.Error(ctx,gin.H{"code": 401, "msg": ""},"")
	// }
	// dd := Model.PostUser(3, "花咋啊的", "3", "180", "花2@qq.com")
	// fmt.Println("返回修改数据")
	// fmt.Println(dd)
	var UserInfo *Model.UserInfo
	_, claims, err := My_token.ParseToken(tokenString)
	UserInfo = Model.GetUser(claims.Id)
	fmt.Println(claims)
	response.Success(ctx, UserInfo, "")
	if err != nil {
		// ctx.JSON(200, gin.H{"code": 401, "msg": "权限不足1"}) //http.StatusUnauthorized
		// ctx.Abort()
		// return
	}
}

//获取所有用户
func GetUsers(ctx *gin.Context) {
	var UserQuery global.UserQuery
	ctx.ShouldBind(&UserQuery)

	list, Total := Model.GetUsers(&UserQuery)
	data := gin.H{
		"list":  list,
		"total": Total,
	}
	response.Success(ctx, data, "")
}

//发送短信
func SendEmail(ctx *gin.Context) {

}

//验证码
func Captcha(ctx *gin.Context) {
	My_Captcha.NewCaptcha(ctx)
	// response.Success(ctx, gin.H{"Code": 200, "data": captcha.New()}, "")
}

//获取验证码图片
func CaptchaImg(c *gin.Context) {
	captchaId := c.Param("captchaId")
	w, _ := strconv.Atoi(c.Param("width"))
	h, _ := strconv.Atoi(c.Param("height"))
	_ = Serve(c.Writer, c.Request, captchaId, ".png", "zh", false, w, h)
}

func Serve(w http.ResponseWriter, r *http.Request, id, ext, lang string, download bool, width, height int) error {
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")

	var content bytes.Buffer
	switch ext {
	case ".png":
		w.Header().Set("Content-Type", "image/png")
		_ = captcha.WriteImage(&content, id, width, height)
	case ".wav":
		w.Header().Set("Content-Type", "audio/x-wav")
		_ = captcha.WriteAudio(&content, id, lang)
	default:
		return captcha.ErrNotFound
	}

	if download {
		w.Header().Set("Content-Type", "application/octet-stream")
	}
	http.ServeContent(w, r, id+ext, time.Time{}, bytes.NewReader(content.Bytes()))
	return nil
}

type Openaidata struct {
	Key      string
	Text     string
	Model    string
	Messages []openai.ChatCompletionMessage
}

func Streamopai(ctx *gin.Context) {
	var u Openaidata
	erri := ctx.ShouldBind(&u)
	fmt.Printf("Completion error: %v\n", erri)
	fmt.Println(u.Messages[len(u.Messages)-1])
	fmt.Println(global.ConfigYml.GetString("openai.Key"))
	fmt.Println(global.ConfigYml.GetFloat64("openai.Temperature"))
	Authorization := ctx.GetHeader("Authorization")
	_, claims, err := My_token.ParseToken(Authorization)
	if err != nil {
		ctx.JSON(401, gin.H{"Code": 401, "msg": "权限过期"}) //http.StatusUnauthorized
		ctx.Abort()
		return
	}
	fmt.Println(claims)
	fmt.Println(claims.Id)
	Model.Streamopai(u.Messages[len(u.Messages)-1].Content, claims.Id)
	// var Temperature = global.ConfigYml.GetFloat64("openai.Temperature").(float32)
	// Temperature, _ := global.ConfigYml.GetInt64("openai.Temperature")
	// ctx := context.Background()
	// req := openai.ChatCompletionRequest{
	// 	Model:       Model,
	// 	MaxTokens:   global.ConfigYml.GetInt("openai.MaxTokens"),
	// 	Temperature: float32(global.ConfigYml.GetFloat64("openai.Temperature")), //断言，转换数据格式,
	// 	Messages:    u.Messages,
	// 	Stream:      true,
	// }
	// flusher, ok := c.Writer.(http.Flusher)
}

var lastAccessTime time.Time
var mutex sync.Mutex

func Getuserli(ctx *gin.Context) {
	mutex.Lock()
	defer mutex.Unlock()
	if time.Since(lastAccessTime) < 10*time.Second {
		response.Error(ctx, gin.H{"Code": 400}, "接口访问限制，请稍后重试")
		return
	}
	lastAccessTime = time.Now()
	list := Model.Getuserli()
	response.Success(ctx, list, "")
}

//获取 gpt 问题生成uuid,存入redis返回uuid
func PustOpenai(ctx *gin.Context) {
	var u Openaidata
	ctx.ShouldBind(&u)
	// fmt.Println(u.Messages)
	id := Utils.Setuuid()
	Jsondata, _ := json.Marshal(u)
	// jsonStr := `{"name": "John", "age": 30}`
	err := global.Redis.Set(ctx, id, Jsondata, 50*time.Second).Err()
	if err != nil {
		fmt.Println(err)
	}
	Authorization := ctx.GetHeader("Authorization")
	_, claims, _ := My_token.ParseToken(Authorization)
	Model.Streamopai(u.Messages[len(u.Messages)-1].Content, claims.Id)
	response.Success(ctx, id, "")
}

//获取uuid拿到redis并执行openai方法去问问题
func GetOpenai(ctx *gin.Context) {
	Id := ctx.Param("Id")
	var u Openaidata
	value, _ := global.Redis.Get(ctx, Id).Result()
	json.Unmarshal([]byte(value), &u)

	// Model := ctx.Param("Model")
	opai := openai.NewClient(global.ConfigYml.GetString("openai.Key"))
	req := openai.ChatCompletionRequest{
		Model:       u.Model,
		MaxTokens:   global.ConfigYml.GetInt("openai.MaxTokens"),
		Temperature: float32(global.ConfigYml.GetFloat64("openai.Temperature")), //断言，转换数据格式,
		Messages:    u.Messages,
		Stream:      true,
	}
	flusher, ok := ctx.Writer.(http.Flusher)
	if !ok {
		ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("Streaming unsupported"))
		return
	}
	stream, err := opai.CreateChatCompletionStream(ctx, req)
	if err != nil {
		response.Error(ctx, gin.H{"Code": err}, "接口访问限制，请稍后重试")
		fmt.Printf("ChatCompletionStream error: %v\n", err)
		return
	}
	defer stream.Close()
	ctx.Header("Transfer-Encoding", "chunked")
	// ctx.SSEvent("message", "开始")
	fmt.Printf("Stream response: ")
	for {
		response, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			fmt.Println("\nStream finished")
			ctx.SSEvent("stop", "finish")
			flusher.Flush() // 手动刷新响应流
			return
		}
		if err != nil {
			fmt.Printf("\nStream error: %v\n", err)
			ctx.SSEvent("error", "error")
			flusher.Flush() // 手动刷新响应流
			return
		}
		ctx.SSEvent("message", response.Choices[0].Delta.Content)
		fmt.Printf(response.Choices[0].Delta.Content)
		flusher.Flush() // 手动刷新响应流
	}

	// response.Success(ctx, u, "")
}

//获取uuid拿到redis并执行openai方法去问问题
func GetOpenai2(ctx *gin.Context) {
	Id := ctx.Param("Id")
	fmt.Println(Id)
	var u Openaidata
	value, _ := global.Redis.Get(ctx, Id).Result()
	json.Unmarshal([]byte(value), &u)
	sse := make(chan string)
	defer close(sse)
	// 逐步返回给前端
	for {

		message := <-sse
		// fmt.Fprintf(io.Writer, "data: %s\n\n", message)
		flusher, ok := ctx.Writer.(http.Flusher)
		if !ok {
			fmt.Println("Streaming unsupported")
			// http.Error(w, "Streaming unsupported", http.StatusInternalServerError)
			return
		}
		fmt.Println(message)
		flusher.Flush()
	}
	// sse <- "Hello, world!"

	// Model := ctx.Param("Model")
	// opai := openai.NewClient(global.ConfigYml.GetString("openai.Key"))
	// req := openai.ChatCompletionRequest{
	// 	Model:       u.Model,
	// 	MaxTokens:   global.ConfigYml.GetInt("openai.MaxTokens"),
	// 	Temperature: float32(global.ConfigYml.GetFloat64("openai.Temperature")), //断言，转换数据格式,
	// 	Messages:    u.Messages,
	// 	Stream:      true,
	// }
	// ctx.Header("Transfer-Encoding", "chunked")
	// flusher, ok := ctx.Writer.(http.Flusher)
	// if !ok {
	// 	ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("Streaming unsupported"))
	// 	return
	// }
	// stream, err := opai.CreateChatCompletionStream(ctx, req)
	// if err != nil {
	// 	fmt.Printf("ChatCompletionStream error: %v\n", err)
	// 	return
	// }
	// defer stream.Close()
	// for {
	// 	response, err := stream.Recv()
	// 	if errors.Is(err, io.EOF) {
	// 		fmt.Println("\nStream finished")
	// 		ctx.SSEvent("stop", "finish")
	// 		flusher.Flush() // 手动刷新响应流
	// 		return
	// 	}
	// 	if err != nil {
	// 		fmt.Printf("\nStream error: %v\n", err)
	// 		ctx.SSEvent("error", "error")
	// 		flusher.Flush() // 手动刷新响应流
	// 		return
	// 	}
	// 	ctx.SSEvent("message", response.Choices[0].Delta.Content)
	// 	fmt.Printf(response.Choices[0].Delta.Content)
	// 	flusher.Flush() // 手动刷新响应流
	// 	time.Sleep(time.Second)
	// }
}

//移动端上传图片  blob格式的
func UploadImg(ctx *gin.Context) {
	file, handler, err := ctx.Request.FormFile("file")
	fmt.Println("上传图片开始")
	// fmt.Println(file)
	fmt.Println(handler.Filename)

	if err != nil {
		fmt.Println(err.Error())
	}
	defer file.Close()

	File := global.ConfigYml.GetString("FileUploadSetting.UploadFileSavePath")
	// 获取当前时间戳，转成字符串，再随机生成字符串拼接一起，防止重复
	FileName := strconv.FormatInt(time.Now().Unix(), 10) + strconv.Itoa(rand.Intn(999999-100000)+100000) // + fileType
	path, Is := Utils.Upload(file, File, FileName, ".png")
	if Is {
		// path2 := global.ConfigYml.GetString("app.app_url") + ":" + global.ConfigYml.GetString("app.port") + global.ConfigYml.GetString("FileUploadSetting.UploadFileReturnPath") + FileName
		response.Success(ctx, gin.H{
			"url": path,
		}, "上传成功")
	} else {
		response.Error(ctx, gin.H{}, "文件上传失败，请重试")
	}
}

func asd(ctx *gin.Context) {
	Authorization := ctx.GetHeader("Authorization")
	_, claims, _ := My_token.ParseToken(Authorization)
	list := Model.Daddd(claims.Id)
	response.Success(ctx, list, "上传成功")
}

func GetOrganization(ctx *gin.Context) {
	list := Model.GetOrganizationlist()
	response.Success(ctx, gin.H{
		"data": list,
	}, "获取成功")
}

func LoadUser(e *gin.Engine) {
	// 其他静态资源
	// r.Static("/public", "./static")
	authRouter := e.Group("").Use(middleware.JWTAuth("a"))
	{
		authRouter.POST("/api/User/PostUser", PostUser)
		authRouter.POST("/api/User/email", SendEmail)
		authRouter.POST("/api/admin/users", GetUsers)
		authRouter.GET("/api/User/getUser", GetUser)
		authRouter.POST("/api/Streamopai", Streamopai)
		authRouter.GET("/api/User/Exit", Exit)
		authRouter.POST("/api/UploadImg", UploadImg)
		authRouter.GET("/api/User/asd", asd)

	}
	// authRouter.GET("/api/User/login", Login)
	e.POST("/api/User/login", Login)
	e.POST("/api/User/getRoute", GetRoute)
	e.POST("/api/User/Register", Register)
	e.POST("/api/User/Img", Img)
	e.POST("/api/User/PustOpenai", PustOpenai)
	e.GET("/api/User/GetOpenai/:Id", GetOpenai)
	e.GET("/api/User/GetOpenai2/:Id", GetOpenai2)
	e.GET("/api/User/Getuserli", Getuserli)
	// r.GET("/api/User/captcha", Captcha)
	e.GET("/api/auth/captcha", Captcha)
	e.GET("/api/auth/captcha/:captchaId/:width/:height", CaptchaImg)
	e.GET("/api/serverclose", ServerClose)
	e.GET("/api/User/GetOrganization", GetOrganization)
	// r.GET("/api/User/getUser", GetUser):captchaId/:value
}
