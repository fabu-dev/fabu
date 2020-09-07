package db

import (
	"fmt"
	"go-web-frame/pkg/config"
	"net/url"

	"github.com/sirupsen/logrus"

	"github.com/jinzhu/gorm"
)

type MysqlConnector struct {
	Config config.Database
}

func NewMysqlConnector() *MysqlConnector {
	return &MysqlConnector{
		Config: config.Conf.Database,
	}
}

// 创建MySQL连接客户端
// 参考资料：https://gorm.io/zh_CN/
func (c *MysqlConnector) Connect() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=%s", c.Config.User, c.Config.Password, c.Config.Host, c.Config.Port, c.Config.Name, url.QueryEscape("Asia/Shanghai"))
	Mysql, err = gorm.Open("mysql", dsn)

	if err != nil {
		logrus.Error("connect mysql failed!", err)
	} else {
		Mysql.LogMode(true) // 打印mysql执行记录

		logrus.Info("连接mysql成功!")
	}

}

func (c *MysqlConnector) Close() {
	if err := Mysql.Close(); err != nil {
		logrus.Error(err)
	} else {
		logrus.Info("MySQL close!")
	}
}
