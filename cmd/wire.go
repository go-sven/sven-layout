//+build wireinject

package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github/go-sven/sven-layout/app/conf"
	"github/go-sven/sven-layout/app/data"
	"github/go-sven/sven-layout/app/handler/index"
	"github/go-sven/sven-layout/app/server/demo/provider"
	"github/go-sven/sven-layout/router"
)

func wireApp(config *conf.Config) (*gin.Engine,error)  {
	panic(wire.Build(
		//配置
		conf.ProviderSet,
		//数据库 redis
		data.ProviderSet,

		//demo 模块
		provider.ProviderSet,
		//todo...

		//index handler
		index.ProviderSet,
		//todo...

		//router
		router.ProviderSet,

		initEngine,
	))
}