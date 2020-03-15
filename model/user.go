package model

import (
	"time"
)

// User 用户
type User struct {
	ID          uint       `json:"id" gorm:"primary_key"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at" sql:"index"`
	Username    string     `json:"username" gorm:"type:varchar(50);unique;not null"`
	Password    string     `json:"-" gorm:"type:varchar(255);not null"`
	Nickname    string     `json:"nickname" gorm:"type:varchar(255);not null"`
	Avatar      string     `json:"avatar" gorm:"type:varchar(1023);null;default:''"`
	Email       string     `json:"email" gorm:"type:varchar(127);null;default:''"`
	Description string     `json:"description" gorm:"type:varchar(1023);null;default:''"`
	ExpireTime  time.Time  `json:"expire_time" gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
}

func (t User) TableName() string {
	return "user"
}
