package model

import "github.com/jinzhu/gorm"

// Log 日志
type Log struct {
	gorm.Model
	IPAddress string `json:"ip_address" gorm:"type:varchar(127);default:'';null"`
	Type      int    `json:"type" gorm:"not null"`
	LogKey    string `json:"log_key" gorm:"type:varchar(1023);default:'';null"`
	Content   string `json:"content" gorm:"varchar(1023);not null"`
}
