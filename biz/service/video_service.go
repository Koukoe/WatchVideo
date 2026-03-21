package service

import (
	"strconv"
	"strings"

	"WatchVideo/biz/dao/db"
	"WatchVideo/biz/model/api"
	"WatchVideo/biz/model/store"
)

func PublishVideo(currentUserID string, req *api.PublishVideoRequest, savedPlayURL string) (*api.VideoResponse, error) {
	if strings.TrimSpace(currentUserID) == "" || req == nil {
		return nil, ErrInvalidParams
	}

	title := strings.TrimSpace(req.Title)
	desc := strings.TrimSpace(req.Description)
	playURL := strings.TrimSpace(savedPlayURL)
	coverURL := strings.TrimSpace(req.CoverURL)

	if title == "" || playURL == "" {
		return nil, ErrInvalidParams
	}

	nextID, err := db.NextVideoID()
	if err != nil {
		return nil, err
	}
	videoID := strconv.FormatInt(nextID, 10)

	v := &store.Video{
		ID:          videoID,
		AuthorID:    currentUserID,
		Title:       title,
		Description: desc,
		PlayURL:     playURL,
		CoverURL:    coverURL,
	}

	if err := db.CreateVideo(v); err != nil {
		return nil, err
	}

	return &api.VideoResponse{
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
	}, nil
}

func ListVideosByUser(req *api.VideoListRequest) (*api.VideoListResponse, int64, error) {
	if req == nil || strings.TrimSpace(req.UserID) == "" {
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

	total, err := db.CountVideosByUserID(req.UserID)
	if err != nil {
		return nil, 0, err
	}

	videos, err := db.ListVideosByUserID(req.UserID, offset, pageSize)
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

	return &api.VideoListResponse{Videos: resp}, total, nil
}

func SearchVideos(req *api.SearchVideosRequest) (*api.SearchVideosResponse, int64, error) {
	if req == nil {
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

	total, err := db.CountSearchVideos(req.Keywords, req.Username, req.FromDate, req.ToDate)
	if err != nil {
		return nil, 0, err
	}

	videos, err := db.SearchVideos(req.Keywords, req.Username, req.FromDate, req.ToDate, offset, pageSize)
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

	return &api.SearchVideosResponse{Videos: resp}, total, nil
}
