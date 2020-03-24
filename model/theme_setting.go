package model

import "github.com/jinzhu/gorm"

// ThemeSetting 主题设置
type ThemeSetting struct {
	gorm.Model
	ThemeID      string `json:"theme_id" gorm:"type:varchar(255);not null"`
	SettingKey   string `json:"setting_key" gorm:"type:varchar(255);not null"`
	SettingValue string `json:"setting_value" gorm:"type:varchar(10239);not null"`
}
