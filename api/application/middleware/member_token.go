package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"fabu.dev/api/pkg/constant"
	"fabu.dev/api/pkg/utils"

	"fabu.dev/api/pkg/db"

	"fabu.dev/api/application/model"
	"fabu.dev/api/pkg/api"
	"fabu.dev/api/pkg/api/code"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// 做权限验证，除了auth控制器之外，其他控制器都需要有有效的token
func VerifyToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Access-Token")
		if len(token) < 1 {
			api.SetResponse(c, http.StatusForbidden, code.ErrorSign, "")
			return
		}

		j := utils.NewJWT()
		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == utils.TokenExpired {
				c.JSON(http.StatusOK, gin.H{
					"status": -1,
					"msg":    "授权已过期",
				})
				c.Abort()
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"status": -1,
				"msg":    err.Error(),
			})
			c.Abort()
			return
		}

		// 先从redis内读取缓存
		memberInfo, errDb := GetMemberInfo(claims.ID)
		if errDb != nil {
			api.SetResponse(c, http.StatusForbidden, code.ErrorSign, errDb.Message)
			return
		}

		//将会员信息放到gin框架的context中
		c.Set("member_id", int64(memberInfo.Id)) // gin框架 没有c.GetUint64()方法，这里转成int64
		c.Set("account", memberInfo.Account)

		c.Next()
	}
}

// 根据token获取member信息
func GetMemberInfo(id uint64) (*model.MemberInfo, *api.Error) {
	// 先从redis内读取缓存
	memberInfo, _ := GetMemberInfoFromRedis(id)
	if memberInfo != nil {
		return memberInfo, nil
	}

	memberInfo, err := model.NewMember().GetDetailByID(id)

	memberData, errJson := json.Marshal(memberInfo)
	if errJson != nil {
		logrus.Error(err)
	}

	errJson = db.Redis.Set(fmt.Sprintf("%s%d", constant.MemberCacheKey, id), memberData, time.Minute*60).Err()
	if errJson != nil {
		logrus.Error(err)
	}

	return memberInfo, err
}

// 从缓存中获取会员信息
func GetMemberInfoFromRedis(id uint64) (*model.MemberInfo, *api.Error) {
	memberData, err := db.Redis.Get(fmt.Sprintf("%s%d", constant.MemberCacheKey, id)).Bytes()
	if err != nil {
		return nil, api.NewError(code.ErrorSyntax, code.GetMessage(code.ErrorSyntax))
	}
	if len(memberData) == 0 {
		return nil, nil
	}

	memberInfo := &model.MemberInfo{}

	if err := json.Unmarshal(memberData, memberInfo); err != nil {
		return nil, api.NewError(code.ErrorSyntax, code.GetMessage(code.ErrorSyntax))
	}

	return memberInfo, nil
}
