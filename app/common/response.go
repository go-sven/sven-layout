package common

import (
	"context"
)

type Reply struct {
	Msg string `json:"msg"`
	Code int `json:"code"`
	Data interface{} `json:"data"`
	TraceId interface{} `json:"trace_id"`
}

func NewReply(ctx context.Context) *Reply  {
	return &Reply{
		Msg:     "success",
		Code:    200,
		TraceId: ctx.Value("traceId"),
	}
}

func (r *Reply) SetMsg (msg string) *Reply {
	r.Msg = msg
	return r
}

func (r *Reply) SetCode (code int) *Reply {
	r.Code = code
	return r
}

func (r *Reply) SetData(data interface{}) *Reply {
	r.Data = data
	return r
}