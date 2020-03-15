package model

import "github.com/jinzhu/gorm"

// Category 分类
type Category struct {
	gorm.Model
	Description string `json:"description" gorm:"type:varchar(100);default:'';null"`
	Name        string `json:"name" gorm:"type:varchar(255);not null"`
	ParentID    uint   `json:"parent_id" gorm:"type:bigint;default:0;null"`
	Slug        string `json:"slug" gorm:"type:varchar(255);null;unique"`
	SlugName    string `json:"slug_name" gorm:"type:varchar(50);not null;unique"`
	Thumbnail   string `json:"thumbnail" gorm:"type:varchar(1023);default:'';null"`
}

func (t Category) TableName() string {
	return "category"
}
