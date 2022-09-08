package repository

import (
	"context"
	"github/go-sven/sven-layout/app/data"
	"github/go-sven/sven-layout/app/server/demo/model"
	"github/go-sven/sven-layout/logger"
	"gorm.io/gorm"
)

//该结构体实现了 model 下的 DemoRepo interface
type demoRepo struct {
	db *gorm.DB
	trans *data.Transact
}

// 业务数据访问，包含 cache、db 等封装，
//实现了 useCase 的 repo 接口。
//我们可能会把 data 与 dao 混淆在一起，
//data 偏重业务的含义，它所要做的是将领域对象重新拿出来，我们去掉了 DDD 的 infra层。

func NewDemoRepo(db *gorm.DB,trans *data.Transact) model.DemoRepo {
	return &demoRepo{
		db: db,
		trans: trans,
	}
}

func (r *demoRepo) SayHello(ctx context.Context) string {
	return "hello sven"
}


func (r *demoRepo) Save(ctx context.Context,m *model.Demo) (*model.Demo, error) {
	err := r.trans.DB(ctx).Create(&m).Error
	if err != nil {
		logger.ErrorWithCtx(ctx,"demo save err:",err)
		return nil,err
	}
	return m, nil
}

func (r *demoRepo) Update(ctx context.Context, m *model.Demo) (*model.Demo, error) {

	return m, nil
}

func (r *demoRepo) FindById(ctx context.Context, id int64) (model *model.Demo,err  error) {
	err = r.db.First(&model,id).Error
	if err != nil {
		logger.ErrorWithCtx(ctx,"demo findById err:",err)
		return nil,err
	}
	return model, nil
}

func (r *demoRepo) ListAll(ctx context.Context) ([]*model.Demo, error) {
	return nil, nil
}
