package repo

import (
	"blog/model"
	"blog/model/query"
	"blog/repo/db"
	"blog/response"
	"blog/tool"
)

// CountPostByStatus count post by status
func CountPostByStatus(status int) int64 {
	var count int64
	db.GetMDB().Model(&model.Post{}).Where("status = ?", status).Count(&count)
	return count
}

// CountPostVisit count post visit
func CountPostVisit() int64 {
	var count int64
	db.GetMDB().Model(&model.Post{}).Select("sum(visits)").Find(&count)
	return count
}

// CountPostLike count post like
func CountPostLike() int64 {
	var count int64
	db.GetMDB().Model(&model.Post{}).Select("sum(likes)").Find(&count)
	return count
}

// PagePost page post
func PagePost(q query.PostQuery) *response.PageInfo {
	if q.Limit < 0 {
		q.Limit = 0
	}
	if q.Offset < 0 {
		q.Offset = 0
	}
	page := &response.PageInfo{
		Offset: q.Offset,
		Limit:  q.Limit,
	}
	qs := ""
	if tool.IsNotBlank(q.Keyword) {
		qs += "title = " + q.Keyword
	}

	posts := []model.Post{}
	db.GetMDB().Where(qs).Order("created_at desc").Limit(q.Limit).Offset(q.Offset).Find(&posts).Count(&page.Total)
	page.Data = posts
	return page
}

// PagePostByStatus get page of post by status
func PagePostByStatus(q query.PostQuery) *response.PageInfo {
	if q.Limit < 0 {
		q.Limit = 0
	}
	if q.Offset < 0 {
		q.Offset = 0
	}
	page := &response.PageInfo{
		Offset: q.Offset,
		Limit:  q.Limit,
	}

	posts := []model.Post{}
	db.GetMDB().Where("status = ?", q.Status).Order("created_at desc").Limit(q.Limit).Offset(q.Offset).Find(&posts).Count(&page.Total)
	page.Data = posts
	return page
}

// GetPostById get post by id
func GetPostById(id uint) model.Post {
	p := model.Post{}
	p.ID = id
	db.GetMDB().First(&p)
	return p
}

// IncreaseLike like post
func IncreaseLike(postId uint) {
	post := &model.Post{}
	mdb := db.GetMDB().Model(post)
	mdb.First(post)
	if post != nil {
		mdb.Update("likes = ?", post.Likes+1)
	}
}

// CreatePost create a post
func CreatePost(post model.Post) model.Post {
	db.GetMDB().Create(&post)
	return post
}
