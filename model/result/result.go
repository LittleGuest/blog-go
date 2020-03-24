package result

import "time"

// StatisticResult StatisticResult
type StatisticResult struct {
	PostCount       int64 `json:"post_count"`
	CommentCount    int64 `json:"comment_count"`
	CategoryCount   int64 `json:"category_count"`
	AttachmentCount int64 `json:"attachment_count"`
	TagCount        int64 `json:"tag_count"`
	JournalCount    int64 `json:"journal_count"`
	Birthday        int64 `json:"birthday"`
	EstablishDays   int64 `json:"establish_days"`
	LinkCount       int64 `json:"link_count"`
	VisitCount      int64 `json:"visit_count"`
	LikeCount       int64 `json:"like_count"`
}

// EnvironmentResult EnvironmentResult
type EnvironmentResult struct {
	Database  string `json:"database,omitempty"`
	StartTime int64  `json:"start_time,omitempty"`
	Version   string `json:"version,omitempty"`
}

// PostMinimalResult PostMinimalResult
type PostMinimalResult struct {
	Id              uint      `json:"id,omitempty"`
	Title           string    `json:"title,omitempty"`
	Status          int       `json:"status,omitempty"`
	Slug            string    `json:"slug,omitempty"`
	EditorType      int       `json:"editor_type,omitempty"`
	CreatedAt       time.Time `json:"created_at,omitempty"`
	UpdatedAt       time.Time `json:"updated_at,omitempty"`
	DeletedAt       time.Time `json:"deleted_at,omitempty"`
	MetaKeywords    string    `json:"meta_keywords,omitempty"`
	MetaDescription string    `json:"meta_description,omitempty"`
	FullPath        string    `json:"full_path,omitempty"`
	// Thumbnail       string    `json:"thumbnail,omitempty"`
	// Url             string    `json:"url,omitempty"`
}

// PostSimpleResult PostSimpleResult
type PostSimpleResult struct {
	PostMinimalResult
	Summary         string `json:"summary,omitempty"`
	Thumbnail       string `json:"thumbnail,omitempty"`
	Visits          int64  `json:"visits,omitempty"`
	DisallowComment int    `json:"disallow_comment,omitempty"`
	Password        string `json:"password,omitempty"`
	Template        string `json:"template,omitempty"`
	TopPriority     int    `json:"top_priority,omitempty"`
	Likes           int64  `json:"likes,omitempty"`
}
