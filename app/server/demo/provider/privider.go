package provider

import (
	"github.com/google/wire"
	"github/go-sven/sven-layout/app/server/demo/cache"
	"github/go-sven/sven-layout/app/server/demo/repository"
	"github/go-sven/sven-layout/app/server/demo/service"
	"github/go-sven/sven-layout/app/server/demo/usecase"
)

//加载wire 需要的provider set 按照以下顺序 加载
var ProviderSet = wire.NewSet(
	//repository
	repository.NewDemoRepo,
	repository.NewDemoExtraRepo,

	//redis
	cache.NewDemoCache,

	//useCase
	usecase.NewDemoUseCase,

	//service
	service.NewDemoService,
)

