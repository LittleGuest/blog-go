package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

// CommentBlackList 评论黑名单
type CommentBlackList struct {
	gorm.Model
	BanTime   time.Time `json:"ban_time" gorm:"type:datetime(6);null"`
	IPAddress string    `json:"ip_address" gorm:"type:varchar(127);not null"`
}
