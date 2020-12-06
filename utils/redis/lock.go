package redis

import (
	"context"
	"time"
)

func SetNX(key,value string)(done bool,err error){
	done,err = rdb.SetNX(context.TODO(),key,value,time.Second).Result()
	return done,err
}

