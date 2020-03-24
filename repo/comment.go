package repo

import (
	"blog/model"
	"blog/repo/db"
)

// CountCommentByStatus count comment by status
func CountCommentByStatus(status int) int64 {
	var count int64
	db.GetMDB().Model(&model.Comment{}).Where("status = ?", status).Count(&count)
	return count
}
