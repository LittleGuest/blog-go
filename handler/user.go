package handler

import (
	"blog/consts"
	"blog/model"
	"blog/model/param"
	"blog/repo"
	"blog/response"
	"crypto/sha1"
	"fmt"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/gookit/validate"
)

// Login login
// @title login
// @router /login [post]
func Login(ctx *gin.Context) {
	loginParam := &param.LoginParam{}
	_ = ctx.Bind(loginParam)
	valid := validate.Struct(loginParam)
	if !valid.Validate() {
		response.Return(ctx, http.StatusBadRequest, http.StatusBadRequest, "", valid.Errors)
		return
	}

	user := repo.GetUserByAccount(loginParam.Password)
	if user == nil {
		response.Return(ctx, http.StatusOK, consts.StatusUserNotFound, "用户不存在", nil)
		return
	}

	// compare password
	h1 := sha1.New()
	_, _ = h1.Write([]byte(loginParam.Password))
	if fmt.Sprintf("%x", h1.Sum(nil)) != user.Password {
		response.Return(ctx, http.StatusOK, consts.StatusAccountOrPasswordError, "账号或密码错误", nil)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.StandardClaims{
		ExpiresAt: 15000,
		Issuer:    loginParam.Account,
	})

	tokenStr, err := token.SignedString([]byte(consts.Salt))
	if err != nil {
		response.Return(ctx, http.StatusInternalServerError, http.StatusInternalServerError, "获取token失败", err)
		return
	}

	response.Return(ctx, http.StatusOK, consts.StatusOk, "登录成功", tokenStr)
}

// Logout logout
// @title login
// @router /login [post]
func Logout(ctx *gin.Context) {
	tokenStr := ctx.Request.FormValue("token")
	log.Println(tokenStr)
	// TODO logout
	log.Println("logout")
}

// GetProfile get user profile
// @title login
// @router /api/v1/admin/users/profiles [get]
func GetProfile(ctx *gin.Context) {
	response.Return(ctx, http.StatusOK, http.StatusOK, "", model.User{})
}
