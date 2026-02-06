package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

type PageData[T any] struct {
	Total    int `json:"total,string"` // 数据量很大使用uint64,一般场景够用
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
	Data     []T `json:"data"`
}

func newResponse(code int, message string, data any) Response {
	return Response{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

func newPageData[T any](total, page, pageSize int, data []T) PageData[T] {
	return PageData[T]{
		Total:    total,
		Page:     page,
		PageSize: pageSize,
		Data:     data,
	}
}

func Result(c *gin.Context, httpCode int, code int, message string, data any) {
	c.JSON(httpCode, newResponse(code, message, data))
}

func Success(c *gin.Context, data any) {
	Result(c, http.StatusOK, CodeSuccess, "success", data)
}

func PageSuccess[T any](c *gin.Context, data []T, total, page, pageSize int) {
	Success(c, newPageData[T](total, page, pageSize, data))
}

func Fail(c *gin.Context, msg string) {
	Result(c, http.StatusOK, CodeServerError, msg, nil)
}

func FailWithCode(c *gin.Context, code int, msg string) {
	Result(c, http.StatusOK, code, msg, nil)
}

func Unauthorized(c *gin.Context, msg string) {
	Result(c, http.StatusUnauthorized, CodeUnauthorized, msg, nil)
}

func Forbidden(c *gin.Context, msg string) {
	Result(c, http.StatusForbidden, CodeForbidden, msg, nil)
}

func NotFound(c *gin.Context, msg string) {
	Result(c, http.StatusNotFound, CodeNotFound, msg, nil)
}
