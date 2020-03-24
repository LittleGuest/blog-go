package model

import "github.com/jinzhu/gorm"

// Comment 评论
type Comment struct {
	gorm.Model
	Type              int    `json:"type" gorm:"type:int;default:0;not null"`
	AllowNotification int    `json:"allow_notification" gorm:"type:tinyint;default:1;null"`
	Author            string `json:"author" gorm:"type:varchar(50);not null"`
	AuthorUrl         string `json:"author_url" gorm:"type:varchar(512);default:'';null"`
	Content           string `json:"content" gorm:"type:varchar(1023);not null"`
	Email             string `json:"email" gorm:"type:varchar(255);not null"`
	GravatarMd5       string `json:"gravatar_md_5" gorm:"type:varchar(128);default:'';null"`
	IPAddress         string `json:"ip_address" gorm:"type:varchar(127);default:'';null"`
	IsAdmin           int    `json:"is_admin" gorm:"type:tinyint;default:0;null"`
	ParentID          uint   `json:"parent_id" gorm:"type:bigint;default:0;null"`
	PostID            uint   `json:"post_id" gorm:"type:bigint;not null"`
	Status            int    `json:"status" gorm:"type:int;default:1;null"`
	TopPriority       int    `json:"top_priority" gorm:"type:int;default:0;null"`
	UserAgent         string `json:"user_agent" gorm:"type:varchar(512);default:'';null"`
}
