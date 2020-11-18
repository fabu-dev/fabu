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
func VerifyAuth() gin.HandlerFunc {
	return func(c *gin.Context) {

		var memberId uint64
		var account string
		var isOk bool

		// 先验证access token
		if memberId, account = VerifyAccessToken(c); memberId > 0 && len(account) > 0 {
			isOk = true
		}

		// 如果access token 没有验证通过，则验证api token
		if !isOk {
			if memberId, account = VerifyToken(c); memberId > 0 && len(account) > 0 {
				isOk = true
			}
		}

		if !isOk {
			api.SetResponse(c, http.StatusForbidden, code.ErrorSign, "")
			return
		}

		//将会员信息放到gin框架的context中
		c.Set("member_id", int64(memberId)) // gin框架 没有c.GetUint64()方法，这里转成int64
		c.Set("account", account)

		c.Next()
	}
}

// 验证API TOKEN
func VerifyToken(c *gin.Context) (uint64, string) {
	token := c.GetHeader("Token")

	if len(token) < 1 {
		return 0, ""
	}

	// 先从redis内读取缓存
	memberInfo, err := GetMemberInfo(token)
	if err != nil {
		return 0, ""
	}

	return memberInfo.Id, memberInfo.Account
}

// 验证ACCESS TOKEN
func VerifyAccessToken(c *gin.Context) (uint64, string) {
	accessToken := c.GetHeader("Access-Token")
	if len(accessToken) < 1 {
		return 0, ""
	}

	objJwt := utils.NewJWT()
	// parseToken 解析token包含的信息
	claims, err := objJwt.ParseToken(accessToken)
	if err != nil {
		logrus.Error("access token err:", err)
		return 0, ""
	}

	return claims.ID, claims.Account
}

// 根据token获取member信息
func GetMemberInfo(token string) (*model.MemberInfo, *api.Error) {
	// 先从redis内读取缓存
	memberInfo, _ := GetMemberInfoFromRedis(token)
	if memberInfo != nil {
		return memberInfo, nil
	}

	memberInfo, err := model.NewMember().GetDetailByToken(token)

	memberData, errJson := json.Marshal(memberInfo)
	if errJson != nil {
		logrus.Error(err)
	}

	errJson = db.Redis.Set(fmt.Sprintf("%s%s", constant.MemberCacheKey, token), memberData, time.Minute*60).Err()
	if errJson != nil {
		logrus.Error(err)
	}

	return memberInfo, err
}

// 从缓存中获取会员信息
func GetMemberInfoFromRedis(token string) (*model.MemberInfo, *api.Error) {
	memberData, err := db.Redis.Get(fmt.Sprintf("%s%s", constant.MemberCacheKey, token)).Bytes()
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
