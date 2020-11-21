package redis

import (
	"github.com/go-redis/redis"
	"lll/study/log"
)

var rdb *redis.Client

func init(){
	rdb = redis.NewClient(&redis.Options{
		Addr: ":6379",
		Password: "",
	})
	pong ,err := rdb.Ping().Result()
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