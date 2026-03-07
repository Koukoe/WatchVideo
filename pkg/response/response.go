package response

import (
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
)

type Base struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
}

type Envelope struct {
	Base Base        `json:"base"`
	Data interface{} `json:"data"`
}

type PageData struct {
	Items interface{} `json:"items"`
	Total int64       `json:"total"`
}

const (
	CodeSuccess      int32 = 10000
	CodeBadRequest   int32 = 10001
	CodeUnauthorized int32 = 10002
	CodeForbidden    int32 = 10003
	CodeInternal     int32 = 10004
)

func Success(c *app.RequestContext, data interface{}) {
	c.JSON(http.StatusOK, Envelope{
		Base: Base{Code: CodeSuccess, Msg: "success"},
		Data: data,
	})
}

func Error(c *app.RequestContext, httpStatus int, code int32, msg string) {
	c.JSON(httpStatus, Envelope{
		Base: Base{Code: code, Msg: msg},
	})
}
