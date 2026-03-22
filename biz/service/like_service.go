package service

import (
	"strings"

	"WatchVideo/biz/dao/db"
	"WatchVideo/biz/model/api"
	"WatchVideo/biz/model/store"
)

func LikeAction(currentUserID string, req *api.LikeActionRequest) error {
	if req == nil || strings.TrimSpace(currentUserID) == "" {
		return ErrInvalidParams
	}

	videoID := strings.TrimSpace(req.VideoID)
	if videoID == "" || (req.Action != 1 && req.Action != 2) {
		return ErrInvalidParams
	}

	// 增加 visit_count
	if err := db.IncVideoVisitCount(videoID); err != nil {
		return err
	}

	exist, err := db.GetLikeByUserAndVideo(currentUserID, videoID)
	if err != nil {
		return err
	}

	if req.Action == 1 {
		if exist != nil {
			return nil
		}
		created, err := db.CreateLike(&store.Like{
			UserID:  currentUserID,
			VideoID: videoID,
		})
		if err != nil {
			return err
		}
		if !created {
			return nil
		}
		return db.IncVideoLikeCount(videoID)
	}

	if exist == nil {
		return nil
	}
	deleted, err := db.DeleteLike(currentUserID, videoID)
	if err != nil {
		return err
	}
	if deleted == 0 {
		return nil
	}
	return db.DecVideoLikeCount(videoID)
}

func LikeList(currentUserID string, req *api.LikeListRequest) (*api.LikeListResponse, int64, error) {
	if req == nil || strings.TrimSpace(currentUserID) == "" {
		return nil, 0, ErrInvalidParams
	}

	pageNum := req.PageNum
	pageSize := req.PageSize
	if pageNum <= 0 {
		pageNum = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	offset := (pageNum - 1) * pageSize

	total, err := db.CountLikesByUserID(currentUserID)
	if err != nil {
		return nil, 0, err
	}

	videos, err := db.ListLikedVideosByUserID(currentUserID, offset, pageSize)
	if err != nil {
		return nil, 0, err
	}

	resp := make([]*api.VideoResponse, 0, len(videos))
	for _, v := range videos {
		resp = append(resp, &api.VideoResponse{
			ID:           v.ID,
			UserID:       v.AuthorID,
			Title:        v.Title,
			Description:  v.Description,
			PlayURL:      v.PlayURL,
			CoverURL:     v.CoverURL,
			LikeCount:    v.LikeCount,
			CommentCount: v.CommentCount,
			VisitCount:   v.VisitCount,
			CreatedAt:    v.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:    v.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return &api.LikeListResponse{Videos: resp}, total, nil
}
