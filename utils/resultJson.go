package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 返回的对象
type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

const (
	ERROR   = 0
	SUCCESS = 1
)

// 封装成功返回的对象
func Success(data interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, Result{
		Code: SUCCESS,
		Msg:  "success",
		Data: data,
	})
}

// 封装失败返回的对象
func Fail(msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Result{
		Code: ERROR,
		Msg:  msg,
		Data: nil,
	})
}
