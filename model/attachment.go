package model

import "github.com/jinzhu/gorm"

type Attachment struct {
	gorm.Model
	FileKey   string `json:"file_key" gorm:"type:varchar(2047);default:'';null"`
	Height    string `json:"height" gorm:"type:int;default:0;null"`
	MediaType string `json:"media_type" gorm:"type:varchar(127);not null"`
	Name      string `json:"name" gorm:"type:varchar(255);not null"`
	Path      string `json:"path" gorm:"type:varchar(1023);not null"`
	Size      uint   `json:"size" gorm:"type:bigint;not null"`
	Suffix    string `json:"suffix" gorm:"type:varchar(50);default:'';null"`
	ThumbPath string `json:"thumb_path" gorm:"type:varchar(1023);default:'';null"`
	Type      int    `json:"type" gorm:"type:int;default:0;null"`
	Width     int    `json:"width" gorm:"type:int;default:0;null"`
}
