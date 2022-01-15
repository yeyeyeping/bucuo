package middleware

import (
	"bucuo/model"
	"bucuo/util"
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type JwtClaim struct {
	UserId string `json:"user_id"`
	jwt.StandardClaims
}

var signkey = []byte(util.JwtKey)

// GenerateJwt 生成jwt签名
func GenerateJwt(userid uint) (string, error) {
	expire, err := strconv.ParseInt(util.ExpiresAt, 10, 64)
	if err != nil {
		return "", err
	}
	claims := JwtClaim{
		strconv.FormatUint(uint64(userid), 10),
		jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(time.Duration(expire) * time.Minute)),
			Issuer:    util.Issuer,
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
			context.JSON(
				http.StatusUnauthorized,
				model.BaseResponse{
					Code: http.StatusUnauthorized,
				},
			)
			context.Abort()
			return
		}
		userid, err := ParseJwt(token)
		if err != nil {
			context.JSON(
				http.StatusBadRequest,
				model.BaseResponse{
					Code: http.StatusBadRequest,
					Msg:  err.Error(),
				},
			)
			context.Abort()
			return
		}

		context.Set("UserId", userid)
		context.Next()
	}
}
