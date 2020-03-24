package query

// PostQuery PostQuery
type PostQuery struct {
	Keyword    string `form:"keyword"`
	Status     int    `form:"status" uri:"status"`
	CategoryId uint   `form:"category_id"`
	Offset     int64  `form:"offset"`
	Limit      int64  `form:"limit"`
}
