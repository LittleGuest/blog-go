package repo

import (
	"blog/model"
	"blog/repo/db"
)

// CountAttachment count attachment
func CountAttachment() int64 {
	var count int64
	db.GetMDB().Model(&model.Attachment{}).Count(&count)
	return count
}
