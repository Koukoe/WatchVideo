package api

type PublishVideoRequest struct {
	UserID      string `form:"user_id" json:"user_id"`
	Title       string `form:"title" json:"title"`
	Description string `form:"description" json:"description"`
	PlayURL     string `form:"play_url" json:"play_url"`
	CoverURL    string `form:"cover_url" json:"cover_url"`
}

type VideoResponse struct {
	ID           string `json:"id"`
	UserID       string `json:"user_id"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	PlayURL      string `json:"play_url"`
	CoverURL     string `json:"cover_url"`
	LikeCount    int64  `json:"like_count"`
	CommentCount int64  `json:"comment_count"`
	VisitCount   int64  `json:"visit_count"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

type VideoListRequest struct {
	UserID   string `query:"user_id"`
	PageNum  int    `query:"page_num"`
	PageSize int    `query:"page_size"`
}

type VideoListResponse struct {
	Videos []*VideoResponse `json:"videos"`
}
