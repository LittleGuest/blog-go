package handler

import (
	"blog/model"
	"blog/model/param"
	"blog/model/query"
	"blog/repo"
	"blog/response"
	"blog/tool"
	"net/http"

	"github.com/gin-gonic/gin"
)

// PagePost get page of post
// @title get page of post
// @router /api/v1/admin/posts [get]
func PagePost(ctx *gin.Context) {
	q := query.PostQuery{}
	_ = ctx.Bind(&q)
	response.Return(ctx, http.StatusOK, http.StatusOK, "", repo.PagePost(q))
}

// PageLatestPost get latest page of post
// @title get latest page of post
// @router /api/v1/admin/posts/latest [get]
func PageLatestPost(ctx *gin.Context) {
	q := query.PostQuery{}
	_ = ctx.Bind(&q)
	q.Offset = 0
	response.Return(ctx, http.StatusOK, http.StatusOK, "", repo.PagePost(q))
}

// PagePostByStatus get page of post by post status
// @title get page of post by post status
// @router /api/v1/admin/posts/status/:status [get]
func PagePostByStatus(ctx *gin.Context) {
	// param := struct {
	// 	Status int `uri:"status"`
	// }{}
	// ctx.BindUri(&param)
	q := query.PostQuery{}
	_ = ctx.Bind(&q)
	response.Return(ctx, http.StatusOK, http.StatusOK, "", repo.PagePostByStatus(q))
}

// GetPostById get post by id
// @title get post by id
// @router /api/v1/admin/posts/:postId [get]
func GetPostById(ctx *gin.Context) {
	param := struct {
		PostId uint `uri:"postId"`
	}{}
	_ = ctx.BindUri(&param)
	response.Return(ctx, http.StatusOK, http.StatusOK, "", repo.GetPostById(param.PostId))
}

// IncreaseLike like post
// @title like post
// @router /api/v1/admin/posts/:postId/likes [get]
func IncreaseLike(ctx *gin.Context) {
	param := struct {
		PostId uint `uri:"postId"`
	}{}
	_ = ctx.BindUri(&param)
	repo.IncreaseLike(param.PostId)
	response.Return(ctx, http.StatusOK, http.StatusOK, "", nil)
}

// CreatePost create a post
// @title create a post
// @router /api/v1/admin/posts [post]
func CreatePost(ctx *gin.Context) {
	param := &param.PostParam{}
	_ = ctx.Bind(param)
	post := model.Post{}
	if err := tool.CopyStructProperty(param, post); err != nil {
		response.Return(ctx, http.StatusInternalServerError, http.StatusInternalServerError, "保存失败", nil)
		return
	}
	response.Return(ctx, http.StatusOK, http.StatusOK, "", repo.CreatePost(post))
}
