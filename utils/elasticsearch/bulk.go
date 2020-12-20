package elasticsearch

import (
	"context"
	"lll/study/utils/limitarray"
	"time"
)

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
		tick := time.NewTicker(time.Second)
		for false {
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