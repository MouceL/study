package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

func Get(ctx context.Context,key string)(string,error){
	val,err := rdb.Get(ctx,key).Result()
	if err == redis.Nil{
		return val,Err(fmt.Sprintf("%s not exist",key))
	}else if err!=nil{
		return val,err
	}
	return val,err
}

func Err(str string) error{
	return fmt.Errorf(str)
}