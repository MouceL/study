package redis

import (
	"context"
	"lll/study/log"
	"time"
)

func Subscribe(topic string){
	ctx := context.TODO()
	sub := rdb.Subscribe(ctx,topic)

	_, err := sub.Receive(ctx)
	if err != nil {
		panic(err)
	}

	ch := sub.Channel()

	time.AfterFunc(time.Second, func() {
		// When pubsub is closed channel is closed too.
		_ = sub.Close()
	})

	for msg := range ch {
		log.Logger.Info(msg.Channel, msg.Payload)
	}
}
