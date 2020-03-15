package middleware

import (
	"blog/response"
	"blog/tool/strtool"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// CorsMiddleware 跨域中间件
func CorsMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Headers", "*")
		ctx.Header("Access-Control-Allow-Methods", "*")
		if ctx.Request.Method == http.MethodOptions {
			ctx.AbortWithStatus(http.StatusNoContent)
		}
		ctx.Next()
	}
}

// AuthMiddleware 认证中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenStr := ctx.Request.FormValue("token")
		if strtool.IsBlank(tokenStr) {
			ctx.JSON(http.StatusNonAuthoritativeInfo, response.Resp{
				Code: http.StatusNonAuthoritativeInfo,
				Msg:  "token为空",
				Data: nil,
			})
			ctx.Abort()
			return
		}
		log.Printf("token>>>%v", tokenStr)
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return []byte("AllYourBase"), nil
		})
		if token == nil {
			ctx.JSON(http.StatusNonAuthoritativeInfo, response.Resp{
				Code: http.StatusNonAuthoritativeInfo,
				Msg:  "token解析失败",
				Data: nil,
			})
			ctx.Abort()
			return
		}

		if token.Valid {
			fmt.Println("You look nice today")
		} else if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				fmt.Println("That's not even a token")
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				// Token is either expired or not active yet
				fmt.Println("Timing is everything")
			} else {
				fmt.Println("Couldn't handle this token:", err)
			}
		} else {
			fmt.Println("Couldn't handle this token:", err)
		}

		ctx.Next()
	}
}
