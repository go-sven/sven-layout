package cache

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis"
	"github/go-sven/sven-layout/app/server/demo/model"
	"github/go-sven/sven-layout/logger"
	"strconv"
	"time"
)

type DemoCache interface {
	Set (context.Context,int64,*model.DemoInfo, time.Duration) bool
	Get(context.Context,int64) (*model.DemoInfo,error)

}


type demoCache struct {
	rdb *redis.Client
}

func NewDemoCache(rdb  *redis.Client) DemoCache {
	return &demoCache{rdb: rdb}
}

func (c *demoCache) key(id int64) string  {
	return "demo:id:" + strconv.FormatInt(id, 10)
}

func (c *demoCache) Set(ctx context.Context,id int64,demo *model.DemoInfo,expire time.Duration) bool  {
	key := c.key(id)
	jsonStr ,_ := json.Marshal(demo)
	_,err := c.rdb.Set(key,jsonStr,120*time.Second).Result()
	if err != nil {
		logger.ErrorWithCtx(ctx,"demo redis set err:",err)
		return false
	}
	return true
}

func (c *demoCache) Get(ctx context.Context,id int64) (*model.DemoInfo,error)  {
	key := c.key(id)
	res,err := c.rdb.Get(key).Result()
	if err != nil {
		logger.ErrorWithCtx(ctx,"demo redis get err:",err)
		return nil,err
	}
	var data *model.DemoInfo
	_ = json.Unmarshal([]byte(res),&data)
	return data,nil
}


