package model

import (
	"context"
)

//PO层
type Demo struct {
	Id int64
	Name string
	//....
}


//定义dto 结构体
type DemoDto struct {

}

type DemoInfo struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	ExtraField string `json:"extra_field"`
}

//定义操作数据库接口
type DemoRepo interface {
	SayHello(context.Context) string
	Save(context.Context, *Demo) (*Demo,error)
	Update (context.Context, *Demo) (*Demo,error)
	FindById(context.Context, int64) (*Demo,error)
	ListAll(context.Context) ([]*Demo,error)
}


