package model

import "github.com/jinzhu/gorm"

// Menu 菜单
type Menu struct {
	gorm.Model
	Name     string `json:"name" gorm:"type:varchar(50);not null"`
	Icon     string `json:"icon" gorm:"type:varchar(50);default:'';null"`
	ParentID uint   `json:"parent_id" gorm:"type:int;default:0;null"`
	Priority int    `json:"priority" gorm:"type:int;default:0;null"`
	Target   string `json:"target" gorm:"type:varchar(20);default:'_self';null"`
	Team     string `json:"team" gorm:"type:varchar(255);default:'';null"`
	Url      string `json:"url" gorm:"type:varchar(1023);not null"`
}
