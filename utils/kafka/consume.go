package kafka

import (
	"github.com/Shopify/sarama"
	"lll/study/log"
	"sync"
)

// TODO  metric


type KafkaConsumer struct {
	Consumer sarama.Consumer
	Topic string
	Output chan string
}

func NewConsumer(addrs []string)(KafkaConsumer,error){

	config := sarama.NewConfig()
	config.Version = sarama.V0_10_2_0
	consumer,err :=sarama.NewConsumer(addrs,config)
	if err!=nil{
		return nil,err
	}
	return KafkaConsumer{Consumer: consumer},nil
}

// 按 partition 消费数据
func (c *KafkaConsumer) ConsumePartitions(){

	partitions,err:= c.Consumer.Partitions(c.Topic)
	if err!=nil{
		log.Logger.Error(err.Error())
		return
	}

	var wg sync.WaitGroup
	for _,partition := range partitions{
		pc,err := c.Consumer.ConsumePartition(c.Topic,partition,sarama.OffsetNewest)
		if err!=nil{
			return
		}
		defer pc.AsyncClose()

		wg.Add(1)
		go func(pc sarama.PartitionConsumer) {
/*			for{
				select {
				case message :=<- pc.Messages():
					log.Logger.Infof("get message :%s",message)
					c.Output <- string(message.Value)
				case err:= <- pc.Errors():
					log.Logger.Errorf(err.Error())
				}
			}

 */
			defer wg.Done()
			for message := range pc.Messages(){
				log.Logger.Infof("get message :%s",message)
				c.Output <- string(message.Value)
			}

		}(pc)
	}
	wg.Wait()
}

// group 形式消费分组数据
func (c *KafkaConsumer) Consume(){

}