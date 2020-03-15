package model

import "github.com/jinzhu/gorm"

// Option 配置
type Option struct {
	gorm.Model
	Type        int    `json:"type" gorm:"type:int;default:0;null"`
	OptionKey   string `json:"option_key" gorm:"type:varchar(100);not null"`
	OptionValue string `json:"option_value" gorm:"type:varchar(1023);not null"`
}

func (t Option) TableName() string {
	return "option"
}
