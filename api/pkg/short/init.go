package short

import (
	"fabu.dev/api/pkg/constant"
	"fabu.dev/api/pkg/db"
)

const (
	CounterInitNumber int64 = 300
)

// TODO 应该将这个信息保存到mysql中，程序启动后，获取这个值，如果没有则使用默认值
func init() {
	_, err := db.Redis.Get(constant.ShortCounter).Int64()
	if err != nil {
		db.Redis.Set(constant.ShortCounter, CounterInitNumber, 0)
		return
	}

}
