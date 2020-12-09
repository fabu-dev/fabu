package short

import (
	"fabu.dev/api/pkg/constant"
	"fabu.dev/api/pkg/db"
	"fabu.dev/api/pkg/utils"
	"github.com/go-redis/redis"
)

// 短链接 生成，解析
// 参考 https://github.com/wyh267/shortService
/**
需要实现：
1. 原子数数 -- 用redis的incr来实现
2. 短链接池 -- 预先生成一部分短链接key， 用redis的Set实现
3. 数字转字符串key
4. 保存短链接与实际链接的对应关系
5. 解析短链接，查询映射关系，返回原始连接
*/

type Short struct {
	Redis *redis.Client
}

func NewShort() *Short {
	return &Short{
		Redis: db.Redis,
	}
}

func (s Short) Create() string {
	counter, err := s.Redis.Incr(constant.ShortCounter).Result()
	if err != nil {

	}

	shortKey := utils.NumToString(counter)

	return shortKey
}

func (s Short) Save(key, url string) error {
	err := s.Redis.Set(s.GetKey(key), url, 0).Err()

	return err
}

func (s Short) GetKey(key string) string {
	return utils.StringSplice(constant.ShortKeyMap, key)
}

func (s Short) Parser(key string) string {
	url := s.Redis.Get(s.GetKey(key)).String()

	return url
}
