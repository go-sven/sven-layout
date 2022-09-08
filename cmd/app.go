package cmd

import (
	"github.com/gin-gonic/gin"
	"github/go-sven/sven-layout/app/conf"
	"github/go-sven/sven-layout/app/middleware/common"
	"github/go-sven/sven-layout/logger"
	"github/go-sven/sven-layout/router"
)

func Run(config *conf.Config)  {

	//初始化日志
	logger.NewZapLogger(config.LoggerConf)

	//设置
	gin.SetMode(config.WebConf.Mode)
	//启动
	route ,err := wireApp(config)
	if err != nil {
		panic(err)
	}

	//运行
	_ = route.Run(config.WebConf.Port)
}

func initEngine(r router.IRouter) *gin.Engine  {
	//route := gin.Default()
	route := gin.New()
	////改写gin.Default() 默认的 Logger 和 Recovery 中间件
	route.Use(common.Logger(),common.Recovery(false))

	//使用中间件
	//route.Use()
	_ = r.Register(route)
	return route
}
