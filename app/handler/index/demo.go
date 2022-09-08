package index

import (
	"context"
	"github.com/gin-gonic/gin"
	"github/go-sven/sven-layout/app/common"
)

//依赖倒置 handler 依赖 service，但是这里依赖的是接口 实际逻辑的处理是实现了该接口的service结构体
//定义该handler 提供的接口
type DemoService interface {
	GetDemo(context.Context, *GetItemDemoRequest) *common.Reply
	SayHello(context.Context) *common.Reply
	AddDemo (context.Context ,*AddDemoRequest) *common.Reply
}


type DemoHandler struct {
	service DemoService
}

func NewDemoHandler(demoService DemoService) *DemoHandler  {
	return &DemoHandler{
		service: demoService,
	}
}

type GetItemDemoRequest struct {
	Id int64
}

type AddDemoRequest struct {
	Name string
	DemoExtra string
}

func (this *DemoHandler) GetItem(ctx *gin.Context)   {
	var req GetItemDemoRequest
	_ =ctx.ShouldBind(&req)
	out := this.service.GetDemo(ctx,&req)
	ctx.JSON(200,out)
}

func (this *DemoHandler) Hello(ctx *gin.Context)   {
	var req GetItemDemoRequest
	_ =ctx.ShouldBind(&req)
	out := this.service.SayHello(ctx)
	ctx.JSON(200,out)
}

func (this *DemoHandler) Add(ctx *gin.Context)   {
	var req AddDemoRequest
	_ =ctx.ShouldBind(&req)
	out := this.service.AddDemo(ctx,&req)
	ctx.JSON(200,out)
}


