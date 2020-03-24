package repo

import (
	"blog/model"
	"blog/repo/db"
	"blog/response"
)

// PageLog page log
func PageLog(limit int64, offset int64) *response.PageInfo {
	if limit < 0 {
		limit = 0
	}
	if offset < 0 {
		offset = 0
	}
	page := &response.PageInfo{
		Offset: offset,
		Limit:  limit,
	}
	logs := []model.Log{}
	db.GetMDB().Order("created_at desc").Limit(limit).Offset(offset).Find(&logs).Count(&page.Total)
	page.Data = logs
	return page
}

// DeleteAllLogs delete all logs
func DeleteAllLogs() {
	db.GetMDB().Delete(&model.Log{})
}
