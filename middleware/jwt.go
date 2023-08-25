package middleware

import (
	"context"
	"fmt"
	"go/global"
	"go/service/My_token"

	"github.com/gin-gonic/gin"
)

//判断JWT令牌是否在Redis黑名单中
func IsRidesToken(token string) bool {
	ctx := context.Background()
	value, err := global.Redis.Get(ctx, token).Result()
	fmt.Println("断JWT令牌是否在Redis黑名单中")
	fmt.Println(value)
	fmt.Println(err)
	if err != nil {
		return false
	}
	return true
}

func JWTAuth(GuardName string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		if IsRidesToken(tokenString) {
			ctx.JSON(401, gin.H{"Code": 401, "msg": "权限过期"})
			ctx.Abort()
			return
		}
		fmt.Println(IsRidesToken(tokenString))
		fmt.Println("进入中间件")
		if tokenString == "" {
			ctx.JSON(401, gin.H{"Code": 401, "msg": "权限不足"})
			ctx.Abort()
			return
		}
		_, claims, err := My_token.ParseToken(tokenString)
		fmt.Println(err)
		if err != nil {
			ctx.JSON(401, gin.H{"Code": 401, "msg": "权限过期"}) //http.StatusUnauthorized
			ctx.Abort()
			return
		}
		fmt.Println(claims)
	}
}
