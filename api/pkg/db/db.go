package db

import (
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	Mysql       *gorm.DB
	Redis       *redis.Client
	Mongo       *mongo.Client
	dbConnector []Connector
)

type Connector interface {
	Connect()
	Close()
}

func InitConnector() {
	dbConnector = append(dbConnector, NewMysqlConnector())
	dbConnector = append(dbConnector, NewRedisConnector())
	dbConnector = append(dbConnector, NewMongoConnector())
}

func init() {
	InitConnector()
	if len(dbConnector) == 0 {
		logrus.Warning("db connector is empty!")
		return
	}

	for _, connector := range dbConnector {
		connector.Connect()
	}
}

func Destroy() {
	if len(dbConnector) == 0 {
		logrus.Warning("db connector is empty!")
		return
	}

	for _, connector := range dbConnector {
		connector.Close()
	}
}
