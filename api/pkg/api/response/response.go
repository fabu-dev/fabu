package response

import (
	"time"

	"fabu.dev/api/pkg/api/code"

	"github.com/gin-gonic/gin"
)

type HttpResponse struct {
	Code   int         `json:"code"`
	Message      string      `json:"message"`
	Timestamp float64     `json:"duration"`
	Result   interface{} `json:"result"`
}

type Response struct {
	GinCtx *gin.Context
}

func NewResponse(c *gin.Context) *Response {
	return &Response{
		GinCtx: c,
	}
}

// 调用 c.json 响应
func (b *Response) SetResponse(httpCode, responseCode int, data interface{}) {
	b.GinCtx.JSON(httpCode, HttpResponse{
		Code:   responseCode,
		Message:      code.GetMessage(responseCode),
		Timestamp: time.Since(b.GinCtx.GetTime("startTime")).Seconds(),
		Result:   data,
	})

	return
}
