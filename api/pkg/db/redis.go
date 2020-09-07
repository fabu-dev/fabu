package db

import (
	"go-web-frame/pkg/config"

	"github.com/sirupsen/logrus"

	"github.com/go-redis/redis"

	_ "github.com/go-sql-driver/mysql"
)

type RedisConnector struct {
	Config config.Redis
}

func NewRedisConnector() *RedisConnector {
	return &RedisConnector{
		Config: config.Conf.Redis,
	}
}

// 创建Redis连接客户端
// 参考资料：https://redis.uptrace.dev/
func (c *RedisConnector) Connect() {

	option := &redis.Options{
		Addr:         c.Config.Host,
		DB:           c.Config.Db,
		DialTimeout:  c.Config.IdleTimeout,
		ReadTimeout:  c.Config.IdleTimeout,
		WriteTimeout: c.Config.IdleTimeout,
		PoolSize:     c.Config.MaxIdle, //cpu*10
		MinIdleConns: c.Config.MaxActive,
		PoolTimeout:  c.Config.IdleTimeout,
		MaxRetries:   c.Config.MaxRetries,
	}
	if len(c.Config.Password) > 0 {
		option.Password = c.Config.Password
	}

	Redis = redis.NewClient(option)
	if err := Redis.Ping().Err(); err != nil {
		logrus.Error("connect redis failed! ", err)
	} else {
		logrus.Info("连接redis成功!")
	}
}

func (c *RedisConnector) Close() {
	if err := Redis.Close(); err != nil {
		logrus.Error(err)
	} else {
		logrus.Info("Redis close!")
	}
}
