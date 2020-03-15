package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

// Post 帖子
type Post struct {
	gorm.Model
	Type            int       `json:"type" gorm:"type:int;default:0;not null"`
	DisallowComment int       `json:"disallow_comment" gorm:"type:int;default:0;null"`
	EditTime        time.Time `json:"edit_time" gorm:"type:datetime;default:CURRENT_TIMESTAMP;null"`
	EditorType      int       `json:"editor_type" gorm:"type:int;default:0;null"`
	FormatContent   string    `json:"format_content" gorm:"type:longtext;not null"`
	Likes           int64     `json:"likes" gorm:"type:bigint;default:0;null"`
	MetaDescription string    `json:"meta_description" gorm:"type:varchar(1023);default:'';null"`
	MetaKeywords    string    `json:"meta_keywords" grom:"type:varchar(500);default:'';null"`
	OriginalContent string    `json:"original_content" gorm:"type:longtext;not null"`
	Password        string    `json:"-" gorm:"type:varchar(255);default:'';null"`
	Slug            string    `json:"slug" gorm:"type:varchar(255);not null;unique"`
	Status          int       `json:"status" gorm:"type:int;default:1;null"`
	Summary         string    `json:"summary" gorm:"type:varchar(255);null"`
	Template        string    `json:"template" gorm:"type:varchar(255);default:'';null"`
	Thumbnail       string    `json:"thumbnail" gorm:"type:varchar(1023);default:'';null"`
	Title           string    `json:"title" gorm:"type:varchar(100);not null"`
	TopPriority     int       `json:"top_priority" gorm:"type:int;default:0;null"`
	Url             string    `json:"url" gorm:"type:varchar(255);null"`
	Visits          int64     `json:"visits" gorm:"type:bigint;default:0;null"`
}

func (t Post) TableName() string {
	return "post"
}
