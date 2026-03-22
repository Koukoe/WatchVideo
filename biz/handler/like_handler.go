package handler

import (
	"context"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"WatchVideo/biz/model/api"
	"WatchVideo/biz/service"
	"WatchVideo/pkg/response"

	"github.com/cloudwego/hertz/pkg/app"
)

func LikeAction(ctx context.Context, c *app.RequestContext) {
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

	var req api.LikeActionRequest
	if err := c.BindAndValidate(&req); err != nil {
		response.Error(c, http.StatusBadRequest, response.CodeBadRequest, "invalid request")
		return
	}

	if err := service.LikeAction(userID, &req); err != nil {
		if errors.Is(err, service.ErrInvalidParams) {
			response.Error(c, http.StatusBadRequest, response.CodeBadRequest, "invalid params")
			return
		}
		response.Error(c, http.StatusInternalServerError, response.CodeInternal, "internal server error")
		return
	}

	response.Success(c, nil)
}

func LikeList(ctx context.Context, c *app.RequestContext) {
	userID := strings.TrimSpace(string(c.Query("user_id")))
	if userID == "" {
		response.Error(c, http.StatusBadRequest, response.CodeBadRequest, "invalid params")
		return
	}

	req := api.LikeListRequest{UserID: userID, PageNum: 1, PageSize: 10}

	if pageNumStr := strings.TrimSpace(string(c.Query("page_num"))); pageNumStr != "" {
		pageNum, err := strconv.Atoi(pageNumStr)
		if err != nil {
			response.Error(c, http.StatusBadRequest, response.CodeBadRequest, "invalid page_num")
			return
		}
		req.PageNum = pageNum
	}

	pageSizeStr := strings.TrimSpace(string(c.Query("page_size")))
	if pageSizeStr == "" {
		pageSizeStr = strings.TrimSpace(string(c.Query("size")))
	}
	if pageSizeStr != "" {
		pageSize, err := strconv.Atoi(pageSizeStr)
		if err != nil {
			response.Error(c, http.StatusBadRequest, response.CodeBadRequest, "invalid page_size")
			return
		}
		req.PageSize = pageSize
	}

	items, total, err := service.LikeList(req.UserID, &req)
	if err != nil {
		if errors.Is(err, service.ErrInvalidParams) {
			response.Error(c, http.StatusBadRequest, response.CodeBadRequest, "invalid params")
			return
		}
		response.Error(c, http.StatusInternalServerError, response.CodeInternal, "internal server error")
		return
	}

	response.Success(c, &response.PageData{
		Items: items,
		Total: total,
	})
}
