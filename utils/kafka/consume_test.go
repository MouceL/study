package kafka

import (
	"context"
	"fmt"
	"lll/study/utils/metric"
	"testing"
)

func TestNewConsumer(t *testing.T) {
	consumer,err := NewConsumer([]string{"127.0.0.1:9092"})
	if err!=nil{
		fmt.Printf("get err:%s",err)
	}
	ctx,_ := context.WithCancel(context.Background())
	consumer.ConsumePartitions(ctx)
	metric.ExposeMetric()
	go func() {
		for message := range consumer.Output(){
			//log.Logger.Infof(message)
			fmt.Printf(message)
		}
	}()

	stop := make(chan string)
	<-stop
}



func TestNewGroup(t *testing.T) {

	cg,err := NewGroup([]string{"127.0.0.1:9092"},"test","study",2)
	if err!=nil{
		fmt.Printf("get err:%s",err)
	}
	ctx,_ := context.WithCancel(context.Background())
	cg.Consume(ctx)

	metric.ExposeMetric()
	go func() {
		for message := range cg.Output(){
			//log.Logger.Infof(message)
			fmt.Println(message)
		}
	}()

	stop := make(chan string)
	<-stop
}

