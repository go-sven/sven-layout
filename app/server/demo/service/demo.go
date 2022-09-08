package service

import (
	"context"
	"github/go-sven/sven-layout/app/common"
	"github/go-sven/sven-layout/app/handler/index"
	"github/go-sven/sven-layout/app/server/demo/model"
	"github/go-sven/sven-layout/app/server/demo/usecase"
)

//实现了 api 定义的服务层，
//类似 DDD 的 application 层，
//处理 DTO 到 biz 领域实体的转换(DTO -> DO)，
//同时协同各类 useCase 交互，但是不应处理复杂逻辑
//该层的理解为 DTO层
//useCase 对应 DO层



type demoService struct {
	DemoUseCase *usecase.DemoUseCase
}






func NewDemoService(demoUC *usecase.DemoUseCase) index.DemoService  {

	return &demoService{DemoUseCase: demoUC}
}

func (s *demoService) GetDemo(ctx context.Context,in *index.GetItemDemoRequest) *common.Reply  {
	reply := common.NewReply(ctx)
	data,err := s.DemoUseCase.FindDemo(ctx,in.Id)
	if err !=nil {
		reply.SetCode(0).SetMsg("数据不存在")
		return reply
	}
	reply.SetData(data)
	return reply
}

func (s *demoService) SayHello(ctx context.Context) *common.Reply  {
	reply := common.NewReply(ctx)
	reply.SetData("我是hello world")
	return reply
}

func (s *demoService) AddDemo (ctx context.Context,in *index.AddDemoRequest) *common.Reply   {
	reply := common.NewReply(ctx)

	if in.Name == "" {
		reply.SetCode(100).SetMsg("param name required")
		return reply
	}

	m := &model.Demo{
		Name: in.Name,
	}
	extra := &model.DemoExtra{
		ExtraField: in.DemoExtra,
	}

	_,err :=s.DemoUseCase.CreateDemo(ctx,m,extra)

	if err != nil {
		reply.SetCode(500).SetMsg(err.Error())
		return reply
	}

	reply.SetData(true)

	return reply
}


