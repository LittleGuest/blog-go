package base

import (
	"context"
	"time"
)

// 查询
type Query struct {
	Id        int64     `json:"id" uri:"id"`
	StrId     string    `json:"str_id" uri:"str_id"`
	Offset    int64     `json:"offset"`
	Limit     int64     `json:"limit"`
	Deleted   int64     `json:"deleted"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// 树
type Tree struct {
	Id       int64         `json:"id"`
	ParentId int64         `json:"parent_id"`
	Children []interface{} `json:"children"`
}

// 文件信息
type FileInfo struct {
	Id         string    `json:"id"`          // 文件ID
	MediaType  string    `json:"media_type"`  // 媒体类型
	Suffix     string    `json:"suffix"`      // 后缀
	Size       string    `json:"size"`        // 文件大小
	OriginName string    `json:"origin_name"` // 原始文件名
	FileName   string    `json:"file_name"`   // 新文件名
	OriginUrl  string    `json:"origin_url"`  // 原始路径
	ThumbUrl   string    `json:"thumb_url"`   // 缩略路径
	CreatedAt  time.Time `json:"created_at"`  // 创建时间
}

// 用户上下文
type UserContext struct {
	ctx      context.Context
	Id       int64
	Account  string
	UserType string
	Status   int
}
