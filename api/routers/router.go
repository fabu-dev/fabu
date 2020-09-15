package routers

import (
	"io"
	"os"

	"fabu.dev/api/pkg/api/middleware"
	"fabu.dev/api/pkg/config"
	"github.com/sirupsen/logrus"

	_ "fabu.dev/api/docs/swagger"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func init() {
	logrus.Println("routers init exec")
	SetLogs()

	InitRouter()
	CreateApiRouter()
}

func SetLogs() {
	// 创建日志文件并设置为 gin.DefaultWriter
	logrus.Info("log path", config.Conf.System.LogPath)
	f, _ := os.Create(config.Conf.System.LogPath)
	gin.DefaultWriter = io.MultiWriter(f)

	if gin.Mode() != gin.ReleaseMode {
		gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	}
}

func InitRouter() {
	Router = gin.New()

	if gin.Mode() != gin.ReleaseMode {
		prefix := "/pprof"
		pprof.Register(Router, prefix)
	}

	Router.Use(gin.Logger())
	Router.Use(gin.Recovery())
	Router.Use(middleware.Cors())  // 跨域
	Router.Use(middleware.Consuming())

	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
