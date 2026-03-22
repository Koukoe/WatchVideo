package api

type LikeActionRequest struct {
	UserID  string `form:"user_id" json:"user_id"`
	VideoID string `form:"video_id" json:"video_id"`
	Action  int    `form:"action" json:"action"` // 1-点赞, 2-取消点赞
}

// 没有要求做点踩操作，有点失望

// 查询指定用户点赞的视频列表
type LikeListRequest struct {
	UserID   string `query:"user_id"`
	PageNum  int    `query:"page_num"`
	PageSize int    `query:"page_size"`
}

type LikeListResponse struct {
	Videos []*VideoResponse `json:"videos"`
}
