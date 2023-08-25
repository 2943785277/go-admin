package My_token

import (
	"fmt"
	"go/http/middleware/my_jwt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

//GenerateToken 生成jwt
func GenerateToken(userid int64, username string, phone string, expireAt int64) (tokens string, err error) {

	// 根据实际业务自定义token需要包含的参数，生成token，注意：用户密码请勿包含在token
	customClaims := my_jwt.CustomClaims{
		Id:    userid,
		Name:  username,
		Phone: phone,
		// 特别注意，针对前文的匿名结构体，初始化的时候必须指定键名，并且不带 jwt. 否则报错：Mixture of field: value and value initializers
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 10,       // 生效开始时间
			ExpiresAt: time.Now().Unix() + expireAt, // 失效截止时间
		},
	}

	return my_jwt.CreateToken(customClaims)
}

//解析jwt
func ParseToken(Token string) (*jwt.Token, *my_jwt.CustomClaims, error) {
	Claims := &my_jwt.CustomClaims{}
	token, err := jwt.ParseWithClaims(Token, Claims, func(token *jwt.Token) (i interface{}, err error) {
		return my_jwt.Jwtkey, nil
	})
	fmt.Println("Token是")
	fmt.Println(token)
	fmt.Println(Claims.StandardClaims.ExpiresAt)
	return token, Claims, err
}
