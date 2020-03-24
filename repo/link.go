package repo

import (
	"blog/model"
	"blog/repo/db"
)

// CountLink count link
func CountLink() int64 {
	var count int64
	db.GetMDB().Model(&model.Link{}).Count(&count)
	return count
}
