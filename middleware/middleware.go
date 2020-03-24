package middleware

import (
	"blog/consts"
	"blog/response"
	"blog/tool"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
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
		if tool.IsBlank(tokenStr) {
			response.Return(ctx, http.StatusNonAuthoritativeInfo, http.StatusNonAuthoritativeInfo, "token为空", nil)
			ctx.Abort()
			return
		}
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return []byte(consts.Salt), nil
		})
		if token == nil {
			response.Return(ctx, http.StatusNonAuthoritativeInfo, http.StatusNonAuthoritativeInfo, "token为空", nil)
			ctx.Abort()
			return
		}

		if token.Valid {
			// TODO 登录的用户信息
			ctx.Next()
		} else if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				response.Return(ctx, http.StatusNonAuthoritativeInfo, http.StatusNonAuthoritativeInfo,
					"That's not even a token", nil)
				ctx.Abort()
				return
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				response.Return(ctx, http.StatusNonAuthoritativeInfo, http.StatusNonAuthoritativeInfo,
					"Timing is everything", nil)
				ctx.Abort()
				return
			} else {
				response.Return(ctx, http.StatusNonAuthoritativeInfo, http.StatusNonAuthoritativeInfo,
					"Couldn't handle this token", nil)
				ctx.Abort()
				return
			}
		} else {
			response.Return(ctx, http.StatusNonAuthoritativeInfo, http.StatusNonAuthoritativeInfo,
				"Couldn't handle this token", nil)
			ctx.Abort()
			return
		}
	}
}
