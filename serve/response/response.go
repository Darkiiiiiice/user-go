/**
 * @Author: mariomang
 * @Date: 2023-12-05 15:08:12
 * @File: response/response.go
**/

package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"user-go/serve/status"
)

const (
	defaultLang = "zh-cn"

	serviceErrCode = http.StatusBadRequest
	authnErrCode   = http.StatusUnauthorized
	paramErrCode   = http.StatusUnprocessableEntity
)

var errMap = map[uint32]string{
	status.SuccessCode: "Success",
	status.UnknownCode: "Unknown",
	status.FailedCode:  "Failed",
}

type CommonBody struct {
	Code    uint32 `json:"code"`           // 错误代码
	Message string `json:"message"`        // 提示信息
	Data    any    `json:"data,omitempty"` // 返回数据
}

func SuccessResponse(ctx *gin.Context, data any) {

	res := CommonBody{
		Data: data,
	}
	if msg, ok := errMap[status.SuccessCode]; ok {
		res.Code = status.SuccessCode
		res.Message = msg
	}

	ctx.JSON(http.StatusOK, res)
}

func NewResponse(ctx *gin.Context, code uint32, data any) {

	res := CommonBody{
		Data: data,
	}
	if msg, ok := errMap[code]; ok {
		res.Code = code
		res.Message = msg
	} else {
		res.Code = status.UnknownCode
		res.Message = errMap[res.Code]
	}

	ctx.JSON(http.StatusOK, res)
}
