package kafka

import (
	"context"
	"github.com/Shopify/sarama"
	"lll/study/log"
	"lll/study/utils/metric"
	_ "lll/study/utils/metric"
)

// 一个消费者消费全部的partition
type KafkaConsumer struct {
	consumer sarama.Consumer
	topic string
	output chan string
}

func NewConsumer(addrs []string)(*KafkaConsumer,error){

	config := sarama.NewConfig()
	config.Version = sarama.V2_5_0_0
	consumer,err :=sarama.NewConsumer(addrs,config)
	if err!=nil{
		return nil,err
	}
	return &KafkaConsumer{consumer: consumer,topic: "test",output: make(chan string)},nil
}


func (c*KafkaConsumer) Output()chan string{
	return c.output
}

// 每个partition 都起一个 goroutine 消费
func (c *KafkaConsumer) ConsumePartitions(ctx context.Context){

	partitions,err:= c.consumer.Partitions(c.topic)
	log.Logger.Infof("get partitions %v",partitions)
	if err!=nil{
		log.Logger.Error(err.Error())
		return
	}
	for _,partition := range partitions{
		pc,err := c.consumer.ConsumePartition(c.topic,partition,sarama.OffsetNewest)
		if err!=nil{
			return
		}
		log.Logger.Infof("start go routine to consume partition[%d]",partition)
		go func(ctx context.Context) {
			defer pc.AsyncClose()
			for {
				select {
				case message,ok := <-pc.Messages():
					if !ok{
						log.Logger.Infof("chanel clonsed")
						return
					}
					log.Logger.Infof("get message :%s",string(message.Value))
					c.output <- string(message.Value)
					metric.ConsumeTotal.Inc()
				case <-ctx.Done():
					log.Logger.Info("ctx done")
				}
			}
		}(ctx)
	}
}