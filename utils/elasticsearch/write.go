package elasticsearch

import (
	"context"
	es "gopkg.in/olivere/elastic.v5"
	"lll/study/log"
	"lll/study/utils/limitarray"
	"time"
)
// 条数 + 时间 限制

// 当前就用一个写线程
type Writer struct {
	array limitarray.Limitarray
	bulk *es.BulkService
	ctx context.Context
}

func NewWriteManager() *Writer {
	return &Writer{
		array: ArrayPool.Get().(limitarray.Limitarray),
		bulk: es.NewBulkService(Client),
	}
}

func (w *Writer)Run( input chan limitarray.Limitarray) {

	go func() {
		tick := time.NewTicker(1*time.Second)
		for{
			select {
			case requests, ok := <-input:
				if !ok {
					return
				}
				w.Write2Es(requests)
				ArrayPool.Put(requests)
			case <- w.ctx.Done() :
				return
			case <-	tick.C :
				log.Logger.Debug("heartbeat")
			}
		}
	}()
}

func (w *Writer)Write2Es(requests limitarray.Limitarray) error{

	for _,item := range requests.Arr {
		var msg IndexRequest
		var ok bool
		if msg,ok = item.(IndexRequest);!ok{
			continue
		}
		request := es.NewBulkIndexRequest().Index(msg.Index).Type("data").Doc(msg.Raw)
		w.bulk = w.bulk.Add(request)
	}


	// 逻辑比较复杂
	if w.bulk.NumberOfActions() > 0 {


		response , err := w.bulk.Do(context.TODO())
		if err != nil {
			return err
		}

		// 有数据没写入成功，要重新写入
		if response.Errors {
			for _, item := range response.Items{
				// TODO
				if item != nil {

				}






			}
		}
	}


	return nil
}

