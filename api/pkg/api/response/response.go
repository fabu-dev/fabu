package response

import (
	"go-web-frame/pkg/api/code"
	"time"

	"github.com/gin-gonic/gin"
)

type HttpResponse struct {
	Status   int         `json:"status"`
	Msg      string      `json:"msg"`
	Duration float64     `json:"duration"`
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
		Status:   responseCode,
		Msg:      code.GetMessage(responseCode),
		Duration: time.Since(b.GinCtx.GetTime("startTime")).Seconds(),
		Result:   data,
	})

	return
}
