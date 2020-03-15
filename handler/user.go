package handler

import (
	"blog/consts"
	"blog/response"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

// LoginHandler 登录
func LoginHandler(ctx *gin.Context) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.StandardClaims{
		ExpiresAt: consts.ExpiresAt,
		Issuer:    consts.Issuer,
	})
	tokenStr, err := token.SignedString([]byte(consts.Salt))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Resp{
			Code: http.StatusInternalServerError,
			Msg:  "获取token失败",
			Data: err,
		})
		return
	}

	ctx.JSON(http.StatusOK, response.Resp{
		Code: 0,
		Msg:  "登录成功",
		Data: tokenStr,
	})
}
