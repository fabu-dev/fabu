package middleware

import (
	"net/http"
	"strings"

	"fabu.dev/api/model"
	"fabu.dev/api/pkg/api"
	"fabu.dev/api/pkg/api/code"
	"github.com/gin-gonic/gin"
)

// 做权限验证，除了auth控制器之外，其他控制器都需要有有效的token
func VerifyToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		pathSlice := strings.Split(c.Request.URL.Path, "/")

		if pathSlice[2] != "auth" {
			token := c.GetHeader("Access-Token")

			if len(token) < 1 {
				api.SetResponse(c, http.StatusForbidden, code.ErrorSign, "")
				return
			}

			memberInfo, err := model.NewMember().GetDetailByToken(token)
			if err != nil {
				api.SetResponse(c, http.StatusForbidden, code.ErrorSign, err.Error())
				return
			}

			c.Set("member_id", memberInfo.Id)
		}

		c.Next()
	}
}
