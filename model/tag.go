package model

import "github.com/jinzhu/gorm"

// Tag 标签
type Tag struct {
	gorm.Model
	Name      string `json:"name" gorm:"type:varchar(255);not null"`
	Slug      string `json:"slug" gorm:"type:varchar(255);null;unique"`
	SlugName  string `json:"slug_name" gorm:"type:varchar(255);not null;unique"`
	Thumbnail string `json:"thumbnail" gorm:"type:varchar(1023);default:'';null"`
}

func (t Tag) TableName() string {
	return "tag"
}
