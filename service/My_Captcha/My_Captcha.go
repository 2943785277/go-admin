package My_Captcha

import (
	"go/service/response"

	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
)

//验证码
func NewCaptcha(ctx *gin.Context) {
	response.Success(ctx, gin.H{"data": captcha.New()}, "")
}
