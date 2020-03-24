package handler

import (
	"blog/repo"
	"blog/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// PageLatestLog PageLatestLog
// @title PageLatestLog
// @router /api/v1/admin/logs/latest [get]
func PageLatestLog(ctx *gin.Context) {
	top := ctx.DefaultQuery("top", "10")
	t, _ := strconv.ParseInt(top, 10, 64)
	response.Return(ctx, http.StatusOK, http.StatusOK, "", repo.PageLog(t, 0))
}

// PageLog PageLog
// @title PageLog
// @router /api/v1/admin/logs [get]
func PageLog(ctx *gin.Context) {
	limit := ctx.DefaultQuery("limit", "10")
	offset := ctx.DefaultQuery("offset", "0")
	l, _ := strconv.ParseInt(limit, 10, 64)
	o, _ := strconv.ParseInt(offset, 10, 64)
	response.Return(ctx, http.StatusOK, http.StatusOK, "", repo.PageLog(l, o))
}

// ClearLog clear all logs
// @title clear all logs
// @router /api/v1/admin/logs/clear [get]
func ClearLog(ctx *gin.Context) {
	response.Return(ctx, http.StatusOK, http.StatusOK, "", repo.DeleteAllLogs)
}
