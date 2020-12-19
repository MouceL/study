package elasticsearch

import (
	"context"
	es "gopkg.in/olivere/elastic.v5"
	"lll/study/utils/limitarray"
)
// 条数 + 时间 限制

type WriteManager struct {
	array *limitarray.Limitarray
	bulk *es.BulkService
}

func NewWriteManage() *WriteManager {
	return &WriteManager{
		array: limitarray.New(1000),
		bulk: es.NewBulkService(Client),
	}
}

// 将数据先写到缓存队列中，等待缓存队列满了，再按照批量写入es中
// 这样会不会十分消耗内存，copy TODO
func (w *WriteManager)Write2Buffer(obj interface{}) error{
	if !w.array.Insert(obj){
		items := w.array.Flush()
		w.Write2Es(items)
	}
	return nil
}

func (w *WriteManager)Write2Es(items [] interface{}) error{

	for _,item := range items {
		request := es.NewBulkIndexRequest().Index("").Type("").Doc("")
		w.bulk = w.bulk.Add(request)
	}
	response , err := w.bulk.Do(context.TODO())
	if err != nil {
		return err
	}

	// 有数据没写入成功，要重新写入
	if response.Errors {
		for _, item := range response.Items{

		}
	}
	return nil
}

