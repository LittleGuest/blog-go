package param

import (
	"blog/base"
)

// LoginParam LoginParam
type LoginParam struct {
	base.Query
	Account  string `form:"account" validate:"required|maxLen:255"`
	Password string `form:"password" validate:"required|maxLen:100"`
}

// RestPasswordParam RestPasswordParam
type RestPasswordParam struct {
	Account  string `form:"account" validate:"required"`
	Password string `form:"password" validate:"required"`
	Email    string `form:"email" validate:"email"`
	Code     string `form:"code"`
}

// UserParam UserParam
type UserParam struct {
	Account     string `form:"account" validate:"required|maxLen:50"`
	Nickname    string `form:"nickname" validate:"required|maxLen:255"`
	Email       string `form:"email" validate:"required|email|maxLen:127"`
	Password    string `form:"password" validate:"minLen:8|maxLen:100"`
	Avatar      string `form:"avatar" valieate:"maxLen:1023"`
	Description string `form:"description" valieate:"maxLen:1023"`
}

// InstallParam InstallParam
type InstallParam struct {
	UserParam
	Locale string `form:"locale"`
	Title  string `form:"title" validate:"required"`
	Url    string `form:"url"`
}

// PostParam PostParam
type PostParam struct {
	Type            int    `form:"type" validate:""`
	DisallowComment int    `form:"disallow_comment" validate:""`
	EditorType      int    `form:"editor_type" validate:""`
	FormatContent   string `form:"format_content" validate:""`
	Likes           int64  `form:"likes" validate:""`
	MetaDescription string `form:"meta_description" validate:""`
	MetaKeywords    string `form:"meta_keywords" validate:""`
	OriginalContent string `form:"original_content" validate:""`
	Password        string `form:"password" validate:"maxLen:255"`
	Slug            string `form:"slug" validate:"maxLen:255"`
	Status          int    `form:"status" validate:""`
	Summary         string `form:"summary" validate:""`
	Template        string `form:"template" validate:"maxLen:255"`
	Thumbnail       string `form:"thumbnail" validate:""`
	Title           string `form:"title" validate:"required|maxLen:100"`
	TopPriority     int    `form:"top_priority" validate:"min:0"`
	Url             string `form:"url" validate:"maxLen:1023"`
	Visits          int64  `form:"visits" validate:""`

	TagIds      []uint      `form:"tag_ids"`
	CategoryIds []uint      `form:"category_id"`
	Metas       []MetaParam `form:"metas"`
}

// MetaParam MetaParam
type MetaParam struct {
	PostId uint   `form:"post_id" valieate:"required"`
	Key    string `form:"key" valieate:"required|maxLen:1023"`
	Value  string `form:"value" valieate:"required|maxLen:1023"`
}
