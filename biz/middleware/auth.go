package middleware

import (
	"context"
	"net/http"
	"strings"

	"WatchVideo/pkg/response"
	"WatchVideo/pkg/utils"

	"github.com/cloudwego/hertz/pkg/app"
)

// 双Token认证中间件
func DualTokenAuth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		accessToken := strings.TrimSpace(string(c.GetHeader("Access-Token")))
		refreshToken := strings.TrimSpace(string(c.GetHeader("Refresh-Token")))

		if accessToken == "" || refreshToken == "" {
			response.Error(c, http.StatusUnauthorized, response.CodeUnauthorized, "missing token")
			c.Abort()
			return
		}

		accessUID, err := utils.ParseAndValidate(accessToken, "access")
		if err != nil {
			response.Error(c, http.StatusUnauthorized, response.CodeUnauthorized, "invalid access token")
			c.Abort()
			return
		}

		refreshUID, err := utils.ParseAndValidate(refreshToken, "refresh")
		if err != nil {
			response.Error(c, http.StatusUnauthorized, response.CodeUnauthorized, "invalid refresh token")
			c.Abort()
			return
		}

		if accessUID != refreshUID {
			response.Error(c, http.StatusUnauthorized, response.CodeUnauthorized, "token user mismatch")
			c.Abort()
			return
		}

		// 是蛇形不是驼峰
		c.Set("user_id", accessUID)
		c.Next(ctx)
	}
}
