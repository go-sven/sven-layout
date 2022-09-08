package router

import "github.com/gin-gonic/gin"

func (router *Router) RegisterIndex(route *gin.Engine) {
	demoGroup := route.Group("/api/index")
	{
		demoGroup.GET("/demo/hello",router.IndexDemo.Hello)
		demoGroup.POST("/demo/getItem",router.IndexDemo.GetItem)
		demoGroup.POST("/demo/add",router.IndexDemo.Add)
	}

}

