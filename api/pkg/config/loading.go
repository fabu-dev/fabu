package config

import (
	"go-web-frame/pkg/utils"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Loader struct {
	ConfigFile string
	Parser     *viper.Viper
}

func NewLoader(file string) *Loader {
	return &Loader{
		ConfigFile: file,
		Parser:     viper.New(),
	}
}

// 初始化配置文件
func (l *Loader) InitConfig() {

	// 检查文件是否存在
	if exist := utils.PathExists(l.ConfigFile); !exist {
		logrus.WithField("config-path", l.ConfigFile).Error("找不到配置文件")
		os.Exit(0)
	}

	// 读取配置文件
	l.Parser.SetConfigType("yaml")
	l.Parser.AddConfigPath(".")
	l.Parser.SetConfigFile(l.ConfigFile)

	l.LoadConfig()
	l.WatchConfig()
}

// 加载配置文件
func (l *Loader) LoadConfig() {
	err := l.Parser.ReadInConfig()
	if err != nil {
		logrus.WithField("config-path", l.ConfigFile).Error(err)
		os.Exit(0)
	}

	err = l.Parser.Unmarshal(Conf)
	if err != nil {
		logrus.WithField("config-path", l.ConfigFile).Error(err)
	} else {
		logrus.WithField("config-path", l.ConfigFile).Info("成功读取配置文件")
	}
}

// 监听配置文件是否变更
func (l *Loader) WatchConfig() {
	l.Parser.WatchConfig()
	l.Parser.OnConfigChange(func(e fsnotify.Event) {
		logrus.WithField("config-path", l.ConfigFile).Info("配置文件变更，重新生效")
		l.LoadConfig()
	})
}
