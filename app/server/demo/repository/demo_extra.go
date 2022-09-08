package repository

import (
	"context"
	"github/go-sven/sven-layout/app/data"
	"github/go-sven/sven-layout/app/server/demo/model"
	"github/go-sven/sven-layout/logger"
	"gorm.io/gorm"
)

//该结构体实现了 model 下的 DemoRepo interface
type demoExtraRepo struct {
	db *gorm.DB
	trans *data.Transact
}

func NewDemoExtraRepo(db *gorm.DB,tx *data.Transact) model.DemoExtraRepo {
	return &demoExtraRepo{db: db,trans: tx}
}

func (r *demoExtraRepo) Save( ctx context.Context,m *model.DemoExtra) (*model.DemoExtra, error) {
	err := r.trans.DB(ctx).Create(&m).Error
	if err != nil {
		logger.ErrorWithCtx(ctx,"demo_extra save err:",err)
		return nil,err
	}
	return m, nil
}

func (r *demoExtraRepo) Update(ctx context.Context, m *model.DemoExtra) (*model.DemoExtra, error) {

	return m, nil
}

func (r *demoExtraRepo) FindById(ctx context.Context, id int64) (model *model.DemoExtra,err  error) {
	err = r.db.First(&model,id).Error
	if err != nil {
		return nil,err
	}
	return model, nil
}

func (r *demoExtraRepo) FindByDemoId(ctx context.Context,demoId int64) (model *model.DemoExtra, err error) {
	err = r.db.Where("demo_id = ?",demoId).First(&model).Error
	if err != nil {
		logger.ErrorWithCtx(ctx,"demo_extra FindByDemoId err:",err)
		return nil,err
	}
	return model,nil
}
