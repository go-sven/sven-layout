package model

import "context"

//PO层
type DemoExtra struct {
	Id int64
	DemoId int64
	ExtraField string
	//....
}

//定义操作数据库接口
type DemoExtraRepo interface {
	Save(context.Context, *DemoExtra) (*DemoExtra,error)
	Update (context.Context, *DemoExtra) (*DemoExtra,error)
	FindById(context.Context, int64) (*DemoExtra,error)
	FindByDemoId(context.Context, int64) (*DemoExtra,error)
}

