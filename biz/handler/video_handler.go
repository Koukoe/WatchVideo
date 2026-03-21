package handler

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"WatchVideo/biz/model/api"
	"WatchVideo/biz/service"
	"WatchVideo/pkg/response"

	"github.com/cloudwego/hertz/pkg/app"
)

func PublishVideo(ctx context.Context, c *app.RequestContext) {
	uidAny, ok := c.Get("user_id")
	if !ok {
		response.Error(c, http.StatusUnauthorized, response.CodeUnauthorized, "unauthorized")
		return
	}
	userID, ok := uidAny.(string)
	if !ok || strings.TrimSpace(userID) == "" {
		response.Error(c, http.StatusUnauthorized, response.CodeUnauthorized, "unauthorized")
		return
	}

	var req api.PublishVideoRequest
	if err := c.BindAndValidate(&req); err != nil {
		response.Error(c, http.StatusBadRequest, response.CodeBadRequest, "invalid request")
		return
	}

	fh, err := c.FormFile("data")
	if err != nil {
		response.Error(c, http.StatusBadRequest, response.CodeBadRequest, "missing video file")
		return
	}

	ct := fh.Header.Get("Content-Type")
	if !strings.HasPrefix(ct, "video/") {
		response.Error(c, http.StatusBadRequest, response.CodeBadRequest, "file must be video")
		return
	}

	if err := os.MkdirAll("storage/videos", 0755); err != nil {
		response.Error(c, http.StatusInternalServerError, response.CodeInternal, "mkdir failed")
		return
	}

	ext := filepath.Ext(fh.Filename)
	if ext == "" {
		ext = ".mp4"
	}
	name := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
	savePath := filepath.Join("storage/videos", name)

	if err := c.SaveUploadedFile(fh, savePath); err != nil {
		response.Error(c, http.StatusInternalServerError, response.CodeInternal, "save file failed")
		return
	}

	playURL := "/storage/videos/" + name
	resp, err := service.PublishVideo(userID, &req, playURL)
	if err != nil {
		if errors.Is(err, service.ErrInvalidParams) {
			response.Error(c, http.StatusBadRequest, response.CodeBadRequest, "invalid params")
			return
		}
		response.Error(c, http.StatusInternalServerError, response.CodeInternal, "internal server error")
		return
	}

	response.Success(c, resp)
}

func VideoList(ctx context.Context, c *app.RequestContext) {
	var req api.VideoListRequest
	if err := c.BindAndValidate(&req); err != nil {
		response.Error(c, http.StatusBadRequest, response.CodeBadRequest, "invalid request")
		return
	}

	items, total, err := service.ListVideosByAuthor(&req)
	if err != nil {
		if errors.Is(err, service.ErrInvalidParams) {
			response.Error(c, http.StatusBadRequest, response.CodeBadRequest, "invalid params")
			return
		}
		response.Error(c, http.StatusInternalServerError, response.CodeInternal, "internal server error")
		return
	}

	resp := &response.PageData{
		Items: items,
		Total: total,
	}
	response.Success(c, resp)
}
