package db

import (
	"context"
	"fmt"
	"go-web-frame/pkg/config"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoConnector struct {
	Config config.Mongo
}

func NewMongoConnector() *MongoConnector {
	return &MongoConnector{
		Config: config.Conf.Mongo,
	}
}

// 创建mongodb的client
// 参考资料 - mongo-driver: https://godoc.org/go.mongodb.org/mongo-driver/mongo
// 参考资料-dsn：https://docs.mongodb.com/manual/reference/connection-string
func (c *MongoConnector) Connect() {
	var err error

	dns := fmt.Sprintf("mongodb://%s", c.Config.Host)
	Mongo, err = mongo.NewClient(options.Client().ApplyURI(dns))
	if err != nil {
		logrus.Error("connect mongo failed!", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), c.Config.IdleTimeout)
	defer cancel()

	if err = Mongo.Connect(ctx); err != nil {
		logrus.Error("connect mongo failed!", err)
	} else {
		logrus.Info("连接mongo成功!")
	}
}

func (c *MongoConnector) Close() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := Mongo.Disconnect(ctx); err != nil {
		logrus.Error(err)
	} else {
		logrus.Info("MongoDB close!")
	}
}
