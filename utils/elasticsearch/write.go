package elasticsearch

import (
	"context"
	es "gopkg.in/olivere/elastic.v5"
	"lll/study/log"
	"lll/study/utils/limitarray"
	"sync"
	"time"
)
// 条数 + 时间 限制

// 当前就用一个写线程
type Writer struct {
	array limitarray.Limitarray

	// todo 多个 goroutine 是不是会往同一个 bulk 里add ? 就一个吧
	bulk *es.BulkService
	ctx context.Context
	num int
}

func NewWriteManager(num int) *Writer {

	bulk := es.NewBulkService(Client)
	return &Writer{
		array: ArrayPool.Get().(limitarray.Limitarray),
		bulk: bulk,
		num: 1,
	}
}

// wg 要是指针传递，如果是 值传递那么done 是对其副本的操作
// add 要在 goroutine 外
func (w *Writer)Run(input chan limitarray.Limitarray) {

	wg := &sync.WaitGroup{}
	for i:=0 ; i< w.num ; i++ {
		wg.Add(1)
		go w.Do(wg,input)
	}
	wg.Wait()
}

func (w *Writer)Do(wg *sync.WaitGroup,input chan limitarray.Limitarray) {

	tick := time.NewTicker(1*time.Second)
	for{
		select {
		case requests, ok := <-input:
			if !ok {
				return
			}
			w.Write2Es(requests)
			// requests 用完后 ，放入池子里
			ArrayPool.Put(requests)
		case <- w.ctx.Done() :
			return
		case <-	tick.C :
			log.Logger.Debug("heartbeat")
		}
	}
	wg.Done()
}

// 根据数据信息，创建request
func (w *Writer)Write2Es(requests limitarray.Limitarray) error{

	var msg IndexRequest
	for _,item := range requests.GetData() {
		var ok bool
		if msg,ok = item.(IndexRequest);!ok{
			continue
		}
		request := es.NewBulkIndexRequest().Index(msg.Index).Type("data").Doc(msg.Raw)
		w.bulk = w.bulk.Add(request)
	}

	// 逻辑比较复杂 TODO
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

