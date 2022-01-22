package middleware

import (
	"bucuo/constant/errormsg"
	"bucuo/model/response"
	"bucuo/util/setting"
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type JwtClaim struct {
	UserId string `json:"user_id"`
	jwt.StandardClaims
}

var signkey = []byte(setting.JwtKey)

// GenerateJwt 生成jwt签名
func GenerateJwt(userid string) (string, error) {
	expire, err := strconv.ParseInt(setting.ExpiresAt, 10, 64)
	if err != nil {
		return "", err
	}
	claims := JwtClaim{
		userid,
		jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(time.Duration(expire) * time.Minute)),
			Issuer:    setting.Issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(signkey)
}

// ParseJwt 解析jwt签名
func ParseJwt(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JwtClaim{}, func(token *jwt.Token) (interface{}, error) {
		return signkey, nil
	})

	if claims, ok := token.Claims.(*JwtClaim); ok && token.Valid {
		return claims.UserId, nil
	} else {
		return "", err
	}
}

func JwtMiddleWare() gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.Request.Header.Get("Token")
		if token == "" {
			context.AbortWithStatusJSON(errormsg.Unauthorized,
				response.RespModel(errormsg.JwtError, "", nil))
			return
		}
		userid, err := ParseJwt(token)
		if err != nil {
			context.AbortWithStatusJSON(errormsg.BadRequest,
				response.RespModel(errormsg.JwtError, err.Error(), nil))
			context.Abort()
			return
		}

		context.Set("UserId", userid)
		context.Next()
	}
}
