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

// 用户注册
func Register(ctx context.Context, c *app.RequestContext) {
	var req api.RegisterRequest
	if err := c.BindAndValidate(&req); err != nil {
		response.Error(c, http.StatusBadRequest, response.CodeBadRequest, "invalid request")
		return
	}

	resp, err := service.Register(&req)
	if err != nil {
		if errors.Is(err, service.ErrInvalidParams) {
			response.Error(c, http.StatusBadRequest, response.CodeBadRequest, "invalid username or password")
			return
		}
		if errors.Is(err, service.ErrUserExists) {
			response.Error(c, http.StatusBadRequest, response.CodeBadRequest, "user already exists")
			return
		}
		response.Error(c, http.StatusInternalServerError, response.CodeInternal, "internal server error")
		return
	}

	response.Success(c, resp)
}

// 用户登录
func Login(ctx context.Context, c *app.RequestContext) {
	var req api.LoginRequest
	if err := c.BindAndValidate(&req); err != nil {
		response.Error(c, http.StatusBadRequest, response.CodeBadRequest, "invalid request")
		return
	}

	resp, err := service.Login(&req)
	if err != nil {
		if errors.Is(err, service.ErrInvalidParams) {
			response.Error(c, http.StatusBadRequest, response.CodeBadRequest, "invalid username or password")
			return
		}
		response.Error(c, http.StatusInternalServerError, response.CodeInternal, "internal server error")
		return
	}

	response.Success(c, resp)
}

// 获取用户信息
func GetUserInfo(ctx context.Context, c *app.RequestContext) {
	var req api.UserInfoRequest
	if err := c.BindAndValidate(&req); err != nil || req.UserID == "" {
		response.Error(c, http.StatusBadRequest, response.CodeBadRequest, "invalid request")
		return
	}

	uidAny, ok := c.Get("user_id")
	if !ok {
		response.Error(c, http.StatusUnauthorized, response.CodeUnauthorized, "unauthorized")
		return
	}
	currentUID, ok := uidAny.(string)
	if !ok || currentUID == "" {
		response.Error(c, http.StatusUnauthorized, response.CodeUnauthorized, "unauthorized")
		return
	}

	if req.UserID != currentUID {
		response.Error(c, http.StatusForbidden, response.CodeForbidden, "forbidden")
		return
	}

	resp, err := service.GetUserInfoByID(req.UserID)
	if err != nil {
		if errors.Is(err, service.ErrInvalidParams) {
			response.Error(c, http.StatusBadRequest, response.CodeBadRequest, "user not found")
			return
		}
		response.Error(c, http.StatusInternalServerError, response.CodeInternal, "internal server error")
		return
	}
	response.Success(c, resp)
}

// 上传用户头像
func UploadAvatar(ctx context.Context, c *app.RequestContext) {
	uidAny, ok := c.Get("user_id")
	if !ok {
		response.Error(c, http.StatusUnauthorized, response.CodeUnauthorized, "unauthorized")
		return
	}
	userID, ok := uidAny.(string)
	if !ok || userID == "" {
		response.Error(c, http.StatusUnauthorized, response.CodeUnauthorized, "unauthorized")
		return
	}

	fh, err := c.FormFile("data")
	if err != nil {
		response.Error(c, http.StatusBadRequest, response.CodeBadRequest, "missing file")
		return
	}

	// 简单图片校验
	ct := fh.Header.Get("Content-Type")
	if !strings.HasPrefix(ct, "image/") {
		response.Error(c, http.StatusBadRequest, response.CodeBadRequest, "file must be image")
		return
	}

	if err := os.MkdirAll("storage/avatars", 0755); err != nil {
		response.Error(c, http.StatusInternalServerError, response.CodeInternal, "mkdir failed")
		return
	}

	ext := filepath.Ext(fh.Filename)
	if ext == "" {
		ext = ".jpg"
	}
	name := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
	savePath := filepath.Join("storage/avatars", name)

	if err := c.SaveUploadedFile(fh, savePath); err != nil {
		response.Error(c, http.StatusInternalServerError, response.CodeInternal, "save file failed")
		return
	}

	avatarURL := "/storage/avatars/" + name
	resp, err := service.UploadAvatar(userID, avatarURL)
	if err != nil {
		if errors.Is(err, service.ErrUnauthorized) {
			response.Error(c, http.StatusUnauthorized, response.CodeUnauthorized, "invalid token")
			return
		}
		response.Error(c, http.StatusInternalServerError, response.CodeInternal, "internal server error")
		return
	}

	response.Success(c, resp)
}
