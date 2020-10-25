package design

import (
	"context"
	"fmt"
	"sync"
)


// 工人的电话 ， 电话赤子
type worker struct {
	tel chan string
	telpool chan chan string
	name string

	ctx context.Context
}
// 显示的提供一个 new 函数，说明这些参数是必要的，不要遗漏
func newWorker(telpool chan chan string ,ctx context.Context) *worker{
	return &worker{
		tel:     make(chan string),
		telpool: telpool,
		name:    "name",
		ctx:     ctx,
	}
}
func (w *worker) run(){
	for{
		w.telpool <- w.tel
		select {
		case task := <- w.tel:
			// do job
			fmt.Printf("do %s",task)
		case <- w.ctx.Done():
			fmt.Println("done")
			return
		}
	}
}


type manager struct {
	workers int
	telpool chan chan string

	work chan string
	ctx context.Context
	cancel context.CancelFunc
}

func (m * manager) run(){
	for i:=0;i<m.workers;i++{
		worker := newWorker(m.telpool,m.ctx)
		go worker.run()
	}
	go m.dispatch()
}

func (m * manager)dispatch(){

	for{
		select {
		case task :=  <- m.work:
			singletel := <- m.telpool
			singletel <- task
		case <-m.ctx.Done():
			fmt.Printf("done")
			return
		}
	}
}




// -----------------------------------------------

var ch = make(chan string)
var n = 5

type worker1 struct {

}

func (w *worker1)run()  {
	wg := sync.WaitGroup{}
	for i:=0;i<5;i++{
		wg.Add(1)
		go func() {
			defer wg.Done()
			for task := range ch{
				fmt.Printf("%s",task)
			}
		}()
	}
	wg.Wait()
}




