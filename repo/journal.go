package repo

import (
	"blog/model"
	"blog/repo/db"
)

// CountJournalByStatus count journal by status
func CountJournalByStatus(status int) int64 {
	var count int64
	db.GetMDB().Model(&model.Journal{}).Where("status = ?", status).Count(&count)
	return count
}
