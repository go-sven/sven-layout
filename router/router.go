package router

import (
	"github.com/gin-gonic/gin"
	"github/go-sven/sven-layout/app/handler/index"
)


var _ IRouter = (*Router)(nil)


type IRouter interface {
	Register (route *gin.Engine) error
}

type Router struct {
	IndexDemo *index.DemoHandler
	//

}


func (router *Router) Register (route *gin.Engine) error {
	//index 模块路由注册
	router.RegisterIndex(route)

	//admin 模块路由注册
	//router.RegisterAdmin(route)

	//对外openApi 模块路由设计
	//todo ...


	return nil
}