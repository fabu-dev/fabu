package filter

import (
	"fabu.dev/api/model"
	"github.com/gin-gonic/gin"
)

type BaseFilter struct {
}

func (f *BaseFilter) GetOperator(c *gin.Context) *model.Operator {
	return &model.Operator{
		Id:      c.GetInt64("member_id"),
		Account: c.GetString("account"),
	}
}
