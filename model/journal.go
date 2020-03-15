package model

import "github.com/jinzhu/gorm"

// Journal 日记
type Journal struct {
	gorm.Model
	Content       string `json:"content" gorm:"not null"`
	Likes         uint   `json:"likes" gorm:"type:bigint;default:0;null"`
	SourceContent string `json:"source_content" gorm:"type:varchar(1023);default:'';not null"`
	Type          int    `json:"type" gorm:"type:int;default:1;null"`
}

func (t Journal) TableName() string {
	return "journal"
}
