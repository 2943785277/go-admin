package routers

import (
	"fmt"
	"go/global"
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// type Loginjson struct {
// 	Username  string `json:"username"`
// 	Password  string `json:"password"`
// 	Captcha   string `json:"cap tcha"`
// 	CaptchaID string `json:"captchaID"`
// }
// type Users struct {
// }

// func Jsoa(data map[string]string) map[string]string {
// 	fmt.Println(data)
// 	return data
// }

// //上传图片
// func Img(ctx *gin.Context) {
// 	var u Model.UserInfo
// 	ctx.ShouldBind(&u)
// 	fmt.Println("文件大小为：")
// 	fmt.Println(len(u.Portrait))

// 	re, _ := regexp.Compile(`^data:\s*image\/(\w+);base64,`)
// 	allData := re.FindAllSubmatch([]byte(u.Portrait), 2)
// 	fileType := string(allData[0][1]) //png ，jpeg 后缀获取
// 	fmt.Println(fileType)
// 	base64Str := re.ReplaceAllString(u.Portrait, "")
// 	byte, _ := base64.StdEncoding.DecodeString(base64Str)
// 	// fmt.Println(byte)
// 	date := time.Now().Format("20060102")

// 	if ok := My_File.IsFileExist(global.ConfigYml.GetString("FileUploadSetting.UploadFileSavePath") + date); !ok {
// 		os.Mkdir(global.ConfigYml.GetString("FileUploadSetting.UploadFileSavePath")+date, 0666)
// 	}
// 	FileName := date + "/" + strconv.FormatInt(time.Now().Unix(), 10) + strconv.Itoa(rand.Intn(999999-100000)+100000) + "." + fileType
// 	path := global.ConfigYml.GetString("FileUploadSetting.UploadFileSavePath") + FileName
// 	path2 := global.ConfigYml.GetString("app.app_url") + ":" + global.ConfigYml.GetString("app.port") + global.ConfigYml.GetString("FileUploadSetting.UploadFileReturnPath") + FileName
// 	err := ioutil.WriteFile(path, byte, 0666)
// 	if err != nil {
// 		fmt.Println("写入err")
// 	}
// 	fmt.Println("写入response.Success")
// 	response.Success(ctx, gin.H{
// 		"url": path2,
// 	}, "")
// }

// func IsFileExist(s string) {
// 	panic("unimplemented")
// }

// //登录 4
// func Login(ctx *gin.Context) {
// 	var u Loginjson
// 	err := ctx.ShouldBind(&u)
// 	if err != nil {

// 	}
// 	var Isopen = global.ConfigYml.GetString("jwt.isopen")
// 	Isopenbool, _ := strconv.ParseBool(Isopen)
// 	var UserInfo *Model.UserInfo
// 	UserInfo = Model.Login(u.Username, u.Password)
// 	fmt.Println(UserInfo)
// 	fmt.Println("验证码状态：", Isopenbool)
// 	if Isopenbool {
// 		if captcha.VerifyString(u.CaptchaID, u.Captcha) {
// 			fmt.Println("验证成功")
// 			if UserInfo.Name == "" {
// 				data := gin.H{}
// 				response.Error(ctx, data, "账号或密码错误，请重试")
// 			} else {
// 				token, _ := My_token.GenerateToken(UserInfo.Id, UserInfo.Name, UserInfo.Phone, 28800)
// 				data := gin.H{
// 					"Code":  200,
// 					"rule":  UserInfo.Role,
// 					"token": token,
// 				}
// 				response.Success(ctx, data, "")
// 			}
// 		} else {
// 			response.Error(ctx, gin.H{}, "验证码失败")
// 		}
// 	} else {
// 		if UserInfo.Name == "" {
// 			data := gin.H{}
// 			response.Error(ctx, data, "账号或密码错误，请重试")
// 		} else {
// 			token, _ := My_token.GenerateToken(UserInfo.Id, UserInfo.Name, UserInfo.Phone, 28800)
// 			data := gin.H{
// 				"Code":  200,
// 				"rule":  UserInfo.Role,
// 				"token": token,
// 			}
// 			response.Success(ctx, data, "")
// 		}
// 	}

// }

// //获取路由
// func GetRoute(ctx *gin.Context) {
// 	jsonFile, err := os.Open("data/Route.json")
// 	if err != nil {
// 		fmt.Println("error opening json file")
// 		return
// 	}
// 	defer jsonFile.Close()

// 	jsonData, err := ioutil.ReadAll(jsonFile)
// 	if err != nil {
// 		fmt.Println("error reading json file")
// 		return
// 	}
// 	formData := make(map[string]interface{})
// 	json.Unmarshal(jsonData, &formData)
// 	ctx.JSON(200, gin.H{"Code": 200, "Data": formData["Data"]})
// }

// //注册
// func Register(ctx *gin.Context) {
// 	var u Model.UserInfo
// 	ctx.ShouldBind(&u)
// 	if Model.Register(u.Name, u.Password, u.Phone, u.Email) {
// 		ctx.JSON(200, gin.H{"Code": 200, "Msg": "注册成功"})
// 	} else {
// 		ctx.JSON(200, gin.H{"Code": 401, "Msg": "注册失败"})
// 	}
// }

// func PostUser(ctx *gin.Context) {
// 	var u Model.UserInfo
// 	ctx.ShouldBind(&u)
// 	if Model.PostUser(u.Id, u.Name, u.Password, u.Phone, u.Email, u.Portrait) {
// 		response.Success(ctx, gin.H{"Code": 200}, "修改成功")
// 	} else {
// 		response.Error(ctx, gin.H{"Code": 400}, "修改失败，请重新修改")
// 	}
// }

// //获取用户信息
// func GetUser(ctx *gin.Context) {
// 	tokenString := ctx.GetHeader("Authorization")
// 	fmt.Println("进入GetUser")
// 	fmt.Println(tokenString)
// 	// dd := Model.PostUser(3, "花咋啊的", "3", "180", "花2@qq.com")
// 	// fmt.Println("返回修改数据")
// 	// fmt.Println(dd)
// 	var UserInfo *Model.UserInfo
// 	_, claims, err := My_token.ParseToken(tokenString)
// 	UserInfo = Model.GetUser(claims.Id)
// 	fmt.Println(claims)
// 	response.Success(ctx, UserInfo, "")
// 	if err != nil {
// 		// ctx.JSON(200, gin.H{"code": 401, "msg": "权限不足1"}) //http.StatusUnauthorized
// 		// ctx.Abort()
// 		// return
// 	}
// }

// //获取所有用户
// func GetUsers(ctx *gin.Context) {
// 	// var UserInfo *Model.UserInfo

// 	list, Total := Model.GetUsers()
// 	data := gin.H{
// 		"item":  list,
// 		"total": Total,
// 	}
// 	response.Success(ctx, data, "")
// }

// //发送短信
// func SendEmail(ctx *gin.Context) {

// }

// //验证码
// func Captcha(ctx *gin.Context) {
// 	My_Captcha.NewCaptcha(ctx)
// 	// response.Success(ctx, gin.H{"Code": 200, "data": captcha.New()}, "")
// }

// //获取验证码图片
// func CaptchaImg(c *gin.Context) {
// 	captchaId := c.Param("captchaId")
// 	w, _ := strconv.Atoi(c.Param("width"))
// 	h, _ := strconv.Atoi(c.Param("height"))
// 	_ = Serve(c.Writer, c.Request, captchaId, ".png", "zh", false, w, h)
// }

// func Serve(w http.ResponseWriter, r *http.Request, id, ext, lang string, download bool, width, height int) error {
// 	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
// 	w.Header().Set("Pragma", "no-cache")
// 	w.Header().Set("Expires", "0")

// 	var content bytes.Buffer
// 	switch ext {
// 	case ".png":
// 		w.Header().Set("Content-Type", "image/png")
// 		_ = captcha.WriteImage(&content, id, width, height)
// 	case ".wav":
// 		w.Header().Set("Content-Type", "audio/x-wav")
// 		_ = captcha.WriteAudio(&content, id, lang)
// 	default:
// 		return captcha.ErrNotFound
// 	}

// 	if download {
// 		w.Header().Set("Content-Type", "application/octet-stream")
// 	}
// 	http.ServeContent(w, r, id+ext, time.Time{}, bytes.NewReader(content.Bytes()))
// 	return nil
// }

// //PostJson 获取post json参数
// func Setting(req *http.Request, obj interface{}) error {
// 	body, err := ioutil.ReadAll(req.Body)
// 	if err != nil {
// 		return err
// 	}
// 	err = json.Unmarshal(body, obj)
// 	if err != nil {
// 		return err
// 	}
// 	fmt.Println(body)
// 	return nil
// }

//post传参
// func PostJson(req *http.Request, obj interface{}) error {
// 	body, err := ioutil.ReadAll(req.Body)
// 	if err != nil {
// 		return err
// 	}
// 	fmt.Println(body)
// 	return ({"code":"200"})
// }

//关闭服务
func ServerClose(c *gin.Context) {
	fmt.Println("进入退出服务")
	log.Fatalf("进入退出服务")
	srv := &http.Server{
		Addr: ":" + global.ConfigYml.GetString("app.port"),
		// Handler: r,
	}
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
	log.Fatalf("close:success")

}

func Init() {
	r := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("sessionName", store))
	// 其他静态资源
	// r.Static("/public", "./static")
	r.Static("/public/storage/uploaded", "./storage/uploaded")
	r.Use(Cors())
	LoadShop(r)
	LoadUser(r)
	LoadMenu(r)
	// authRouter := r.Group("").Use(middleware.JWTAuth("a"))
	// {
	// 	authRouter.GET("/api/User/getUser", GetUser)
	// 	authRouter.POST("/api/User/postUser", PostUser)
	// 	authRouter.POST("/api/User/email", SendEmail)
	// 	authRouter.GET("/api/admin/users", GetUsers)
	// }
	// r.GET("/api/User/login", Login)
	// r.POST("/api/User/login", Login)
	// r.POST("/api/User/getRoute", GetRoute)
	// r.POST("/api/User/Register", Register)
	// r.POST("/api/User/Img", Img)
	// // r.GET("/api/User/captcha", Captcha)
	// r.GET("/captcha", Captcha)
	// r.GET("/captcha/:captchaId/:width/:height", CaptchaImg)
	// r.GET("/serverclose", ServerClose)
	// r.GET("/api/User/getUser", GetUser):captchaId/:value

	//监听端口默认为8080
	r.Run(":" + global.ConfigYml.GetString("app.port"))
}
