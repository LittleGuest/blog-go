package model

import "github.com/jinzhu/gorm"

// Meta
type Meta struct {
	gorm.Model
	Type      int    `json:"type" gorm:"type:int;default:0;not null"`
	PostID    uint   `json:"post_id" gorm:"type:bigint;not null"`
	MetaKey   string `json:"meta_key" gorm:"type:varchar(100);not null"`
	MetaValue string `json:"meta_value" gorm:"type:varchar(1023);not null"`
}
