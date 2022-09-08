package usecase

import (
	"context"
	"github/go-sven/sven-layout/app/data"
	"github/go-sven/sven-layout/app/server/demo/cache"
	"github/go-sven/sven-layout/app/server/demo/model"
	"github/go-sven/sven-layout/logger"
	"golang.org/x/sync/singleflight"
)

//该层负责整合mysql redis 的数据
type DemoUseCase struct {
	repo model.DemoRepo
	extraRepo model.DemoExtraRepo
	tx  data.Transaction
	dc cache.DemoCache
	barrier singleflight.Group
}

var (
	bar singleflight.Group
)




// 业务逻辑的组装层，
//类似 DDD 的 domain 层，
//data 类似 DDD 的 repo，
//而 repo 接口在这里定义，使用依赖倒置的原则。
func NewDemoUseCase(repo model.DemoRepo,extraRepo model.DemoExtraRepo,tx data.Transaction,dc cache.DemoCache ) *DemoUseCase  {
	return &DemoUseCase{
		repo:      repo,
		extraRepo: extraRepo,
		tx :        tx,
		dc :		dc,
		barrier: bar,
	}
}

func (uc *DemoUseCase) SayHello(ctx context.Context) string {
	return uc.repo.SayHello(ctx)
}

func (uc *DemoUseCase) FindDemo(ctx context.Context,id int64) (*model.DemoInfo,error)  {
	logger.InfoWithCtx(ctx,"GetDemo:",id)
	demoInfoCache ,_ := uc.dc.Get(ctx,id)
	if demoInfoCache != nil {
		return demoInfoCache,nil
	}
	v,err,_ := uc.barrier.Do("1", func() (interface{}, error) {
		demo,err := uc.repo.FindById(ctx,id)
		if err != nil {
			return nil,err
		}

		extra ,err := uc.extraRepo.FindByDemoId(ctx,demo.Id)
		if err != nil {
			return nil,err
		}
		demoInfo := &model.DemoInfo{
			Id:         demo.Id,
			Name:       demo.Name,
			ExtraField: extra.ExtraField,
		}
		uc.dc.Set(ctx,id,demoInfo,120)
		return demoInfo,nil
	})
	if err != nil {
		return nil,err
	}
	return v.(*model.DemoInfo),nil
}

func (uc *DemoUseCase) CreateDemo (ctx context.Context,demo *model.Demo,extra *model.DemoExtra) (bool,error)  {
	logger.InfoWithCtx(ctx,"AddDemo:",*demo)
	err := uc.tx.ExecTx(ctx, func(ctx context.Context) error {
		m ,err := uc.repo.Save(ctx,demo)
		if err != nil {
			return err
		}

		extra.Id = m.Id
		extra.Id = 1
		_,err = uc.extraRepo.Save(ctx,extra)
		if err !=nil {
			return err
		}
		return nil
	})
	if err != nil {
		return false,err
	}
	return true,nil
}
