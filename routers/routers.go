package routers

import (
	"blog/handler"
	"blog/middleware"
	"net/http"
	"os"

	_ "blog/docs"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// InitRouter 路由
func InitRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.LoggerWithWriter(os.Stdout))
	router.Use(gin.Recovery())
	router.Use(middleware.CorsMiddleware())
	router.Use(gzip.Gzip(gzip.DefaultCompression))

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/api/v1/admin/is_installed", handler.IsInstall)
	router.POST("/api/v1/admin/login", handler.Login)

	// router.Use(middleware.AuthMiddleware())

	router.POST("/api/v1/admin/logout", handler.Logout)
	v1 := router.Group("/api/v1")
	{
		admin := v1.Group("/admin")
		{
			admin.POST("/password/code", handler.SendResetPasswordCode)
			admin.PUT("/password/reset")
			// admin.POST("refresh/{refreshToken}")
			admin.GET("/counts", handler.GetCount)
			admin.GET("/environments", handler.GetEnvironments)
			admin.PUT("/blog-admin")
			admin.POST("/blog/restart")
			admin.GET("/blog/logfile")

			admin.POST("/installations", handler.InstallBlog)

			users := admin.Group("/users")
			{
				users.GET("/profiles", handler.GetProfile)
			}

			posts := admin.Group("/posts")
			{
				posts.GET("/", handler.PagePost)
				posts.POST("/", handler.CreatePost)
				posts.DELETE("/")
				posts.GET("/latest", handler.PageLatestPost)
				posts.GET("/status/:status", handler.PagePostByStatus)
				posts.PUT("/status/:status")
				// posts.PUT("/:postId")
				posts.DELETE("/:postId")
				// posts.GET("/:postId", handler.GetPostById)
				// posts.GET("/:postId/preview")
				// posts.GET("/:postId/likes", handler.IncreaseLike)
				// posts.PUT("/:postId/status/:status")
				// posts.PUT("/:postId/status/draft/content")
			}

			logs := admin.Group("/logs")
			{
				logs.GET("/", handler.PageLatestLog)
				logs.GET("/latest", handler.PageLog)
				logs.GET("/clear", handler.ClearLog)
			}
		}

	}

	// 静态文件服务，访问上传的文件
	router.StaticFS("/upload", http.Dir("upload"))

	return router
}
