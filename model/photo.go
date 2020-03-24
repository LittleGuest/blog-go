package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Photo 图片
type Photo struct {
	gorm.Model
	Name        string    `json:"name" gorm:"type:varchar(255);not null"`
	Description string    `json:"description" gorm:"varchar(255)default:'';null"`
	Url         string    `json:"url" gorm:"type:varchar(1023);not null"`
	Thumbnail   string    `json:"thumbnail" gorm:"type:varchar(1023);default:'';null"`
	Location    string    `json:"location" gorm:"type:varchar(255);default:'';null"`
	TakeTime    time.Time `json:"take_time" gorm:"type:datetime;default:CURRENT_TIMESTAMP;null"`
	Team        string    `json:"team" gorm:"type:varchar(255);default:'';null"`
}
