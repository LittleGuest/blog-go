package model

import "github.com/jinzhu/gorm"

// Link 链接
type Link struct {
	gorm.Model
	Name        string `json:"name" gorm:"type:varchar(255);not null"`
	Description string `json:"description" gorm:"type:varchar(255);default:'';null"`
	Logo        string `json:"logo" gorm:"type:varchar(1023);default:'';null"`
	Priority    int    `json:"priority" gorm:"type:int;default:0;null"`
	Team        string `json:"team" gorm:"type:varchar(255);default:'';null"`
	Url         string `json:"url" gorm:"type:varchar(1023);not null"`
}
