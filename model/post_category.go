package model

import "github.com/jinzhu/gorm"

// PostCategory 帖子分类
type PostCategory struct {
	gorm.Model
	PostID     uint `json:"post_id" gorm:"type:bigint;null"`
	CategoryID uint `json:"category_id" gorm:"type:bigint;null"`
}
