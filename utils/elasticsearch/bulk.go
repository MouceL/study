package elasticsearch

import (
	"context"
	"lll/study/utils/limitarray"
	"time"
)

// 来的所有消息都先写入 array , 当达到一定个数后，一次性写入一个 array

type BulkRequest struct {
	array limitarray.Limitarray
	output chan limitarray.Limitarray
	ctx context.Context
}

func NewBulkRequest(ctx context.Context) *BulkRequest{
	return &BulkRequest{
		array: ArrayPool.Get().(limitarray.Limitarray),
		output: make(chan limitarray.Limitarray),
		ctx: ctx,
	}
}

func (b *BulkRequest)Run( input chan interface{}) {

	go func() {
		tick := time.NewTicker(5*time.Second)
		for {
			select {
			case  request := <- input :
				b.Add2Array(request)
			case <- b.ctx.Done():
				return
			case <- tick.C :
				b.flush()
			}
		}
	}()
}

// TODO 要不要 + lock , 不要 消息是一个一个处理的 不会出现一下子处理两个消息
func (b *BulkRequest) Add2Array(obj interface{}) {

	if b.array.IsFull() {
		b.flush()
	} else {
		b.array.Insert(obj)
	}
}

func (b *BulkRequest) flush() {
	b.output <- b.array
	b.array = ArrayPool.Get().(limitarray.Limitarray)
}