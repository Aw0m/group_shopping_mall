package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"

	"group_shopping_mall/utils/constant"
)

type MyCustomClaims struct {
	OpenID     string `json:"openID"`
	UserID     int64  `json:"userID"`
	UserName   string `json:"userName"`
	SessionKey string `json:"sessionKey"`
	jwt.StandardClaims
}

func Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		//log.Println("Authorization Token: ", tokenString)

		token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return constant.MySigningKey, nil
		})
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "访问未授权"})
			fmt.Println("\n没有访问token，拦截", err)
			c.Abort()
			return
		}
		if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
			fmt.Printf("\nUsername: %v, ExpiresAt %v;   ", claims.UserName, claims.StandardClaims.ExpiresAt)
			// 验证通过，会继续访问下一个中间件
			//c.Next()
			return
		}
		fmt.Println("\n出现错误：", err)
		// 验证不通过，不再调用后续的函数处理
		c.JSON(http.StatusUnauthorized, gin.H{"message": "访问未授权"})
		fmt.Println("\ntoken错误，拦截")
		c.Abort()
	}
}

func GetID(ctx *gin.Context) (openid string, userId int64, retErr error) {
	tokenString := ctx.GetHeader("Authorization")
	if len(tokenString) == 0 {
		return "", 0, errors.New("Authorization is nil")
	}

	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return constant.MySigningKey, nil
	})
	if err != nil {
		return "", 0, errors.Wrap(err, "parse token error")
	}
	claims := token.Claims.(*MyCustomClaims)
	openid = claims.OpenID
	userId = claims.UserID
	return openid, userId, nil
}
