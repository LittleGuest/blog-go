package routers

import (
	"blog/middleware"
	"blog/response"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	// 启用日志中间件
	// router.Use(gin.LoggerWithWriter(logFile))
	// 启用 panic 恢复中间件
	// router.Use(gin.RecoveryWithWriter(logFile))
	router.Use(middleware.CorsMiddleware())
	router.Use(gzip.Gzip(gzip.DefaultCompression))

	// router.GET("/login", handler.LoginHandler)
	router.GET("/api", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, response.Resp{
			Code: http.StatusOK,
			Msg:  "api",
			Data: nil,
		})
	})

	// router.Use(middleware.AuthMiddleware())

	router.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, response.Resp{
			Code: 0,
			Msg:  "msg",
			Data: "nil",
		})
	})

	// router.POST("/oauth/token", func(ctx *gin.Context) {
	//	claims := handler.Claims{}
	//	if err := ctx.Bind(&claims); err != nil {
	//		log.Errorln(err)
	//		ctx.JSON(http.StatusBadRequest, response.Error(consts.ERROR, err.Error()))
	//		return
	//	}
	//	log.Infoln(claims)
	//	token, _ := handler.GenerateToken(claims.Account, claims.Password)
	//	ctx.JSON(http.StatusOK, response.Success(token))
	// })

	// 静态文件服务，访问上传的文件
	router.StaticFS("/upload", http.Dir("upload"))

	return router
}
