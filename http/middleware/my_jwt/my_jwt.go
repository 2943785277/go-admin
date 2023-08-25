package my_jwt

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

// 定义一个 JWT验签 结构体
type JwtSign struct {
	SigningKey []byte
}

var Jwtkey = []byte("")

// CreateToken 生成一个token
func CreateToken(claims CustomClaims) (string, error) {
	// 生成jwt格式的header、claims 部分
	tokenPartA := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 继续添加秘钥值，生成最后一部分

	tokendata, _ := tokenPartA.SignedString(Jwtkey)
	fmt.Println(tokendata)
	return tokenPartA.SignedString(Jwtkey)
}
