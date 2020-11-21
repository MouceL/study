package redis

import "github.com/go-redis/redis"

func Subscribe(topic string){
	sub := rdb.Subscribe(topic)
	iface, err := sub.Receive()
	if err != nil {
		// handle error
	}
	// Should be *Subscription, but others are possible if other actions have been
	// taken on sub since it was created.
	switch iface.(type) {
	case *redis.Subscription:
		// subscribe succeeded
	case *redis.Message:
		// received first message
	case *redis.Pong:
		// pong received
	default:
		// handle error
	}

	ch := sub.Channel()
}
