package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"lll/study/log"
)

var rdb *redis.Client

func init(){
	ctx := context.TODO()
	rdb = redis.NewClient(&redis.Options{
		Addr: ":6379",
		Password: "",
	})
	pong ,err := rdb.Ping(ctx).Result()
	if err!=nil{
		panic(err)
	}
	log.Logger.Infof("success connect to redis %s",pong)
}

func Close(){
	if rdb!=nil{
		rdb.Close()
	}
}