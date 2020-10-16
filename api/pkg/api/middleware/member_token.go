package middleware

import (
	"net/http"

	"fabu.dev/api/model"
	"fabu.dev/api/pkg/api"
	"fabu.dev/api/pkg/api/code"
	"github.com/gin-gonic/gin"
)

// 做权限验证，除了auth控制器之外，其他控制器都需要有有效的token
func VerifyToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Access-Token")

		if len(token) < 1 {
			api.SetResponse(c, http.StatusForbidden, code.ErrorSign, "")
			return
		}

		memberInfo, err := model.NewMember().GetDetailByToken(token)
		if err != nil {
			api.SetResponse(c, http.StatusForbidden, code.ErrorSign, err.Message)
			return
		}

		//将会员信息放到gin框架的context中
		c.Set("member_id", int64(memberInfo.Id)) // gin框架 没有c.GetUint64()方法，这里转成int64
		c.Set("account", memberInfo.Account)

		c.Next()
	}
}
