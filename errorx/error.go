package errorx

import (
	"fmt"
	"net/http"
	"sunylive-gallery/global"

	"github.com/gin-gonic/gin"
)

const (
	ERROR_CODE_INTERNAL = 500 //内部错误
	ERROR_CODE_EXTERNAL = 400 //外部错误（默认）
	ERROR_CODE_AUTH     = 401 //外部错误（默认）
	SHOW_TYPE           = 2   //信息错误
)

type ResponseError struct {
	Success      bool   `json:"success"`      // 【必需】请求状态
	ErrorCode    int    `json:"errorCode"`    // 【选填】错误类型代码：400错误请求，401未授权，500服务器内部错误，200成功
	ErrorMessage string `json:"errorMessage"` // 【选填】向用户显示消息
	TraceMessage string `json:"traceMessage"` // 【选填】调试错误信息，请勿在生产环境下使用，可有可无
	ShowType     int    `json:"showType"`     // 【选填】错误显示类型：0.不提示错误;1.警告信息提示；2.错误信息提示；4.通知提示；9.页面跳转
	TraceId      string `json:"traceId"`      // 【选填】方便后端故障排除：唯一的请求ID
	Host         string `json:"host"`         // 【选填】方便后端故障排除：当前访问服务器的主机
}

type Errorx struct {
	Message string //错误消息
	Addr    string //错误地址
	Error   error  //具体错误
}

func New(msg string, err error) string {
	return fmt.Sprintf("【错误信息:%s】【错误原因:%s】", msg, err)
}

func InternalError(c *gin.Context, err error) {
	msg := "服务器内部错误"
	if global.ReleaseMode {
		msg = New("服务器内部错误", err)
	}
	c.JSON(http.StatusInternalServerError, gin.H{
		"success": false,
		"code":    http.StatusInternalServerError,
		"message": msg,
	})
	c.Abort()
}

func ExternalError(c *gin.Context, err error) {
	msg := "请求参数错误"
	if !global.ReleaseMode {
		msg = New("请求参数错误", err)
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"success": false,
		"code":    http.StatusBadRequest,
		"message": msg,
	})
	c.Abort()
}
