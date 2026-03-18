package service

import (
	"strconv"
	"strings"
	"time"

	"WatchVideo/biz/dao/db"
	"WatchVideo/biz/model/api"
	"WatchVideo/biz/model/store"
)

func PublishVideo(currentUserID string, req *api.PublishVideoRequest, savedPlayURL string) (*api.PublishVideoResponse, error) {
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

	return &api.PublishVideoResponse{
		ID:           v.ID,
		AuthorID:     v.AuthorID,
		Title:        v.Title,
		Description:  v.Description,
		PlayURL:      v.PlayURL,
		CoverURL:     v.CoverURL,
		LikeCount:    v.LikeCount,
		CommentCount: v.CommentCount,
		VisitCount:   v.VisitCount,
		CreatedAt:    v.CreatedAt.Format(time.RFC3339),
		UpdatedAt:    v.UpdatedAt.Format(time.RFC3339),
	}, nil
}
