package model

import "github.com/jinzhu/gorm"

// PostTag 帖子标签
type PostTag struct {
	gorm.Model
	PostID uint `json:"post_id" gorm:"type:int unsigned;not null"`
	TagID  uint `json:"tag_id" gorm:"type:int unsigned;not null"`
}

func (t PostTag) TableName() string {
	return "post_tag"
}
