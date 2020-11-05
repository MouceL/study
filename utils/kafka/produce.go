package kafka

import (
	"github.com/Shopify/sarama"
	"lll/study/log"
	"lll/study/utils/metric"
	"time"
)

type producer struct {
	topic string
	addrs []string
	p sarama.AsyncProducer
	input chan string
	p2 sarama.SyncProducer
}


func NewProducer(addrs []string,topic string)(*producer,error){
	config := sarama.NewConfig()
	config.Version = sarama.V2_5_0_0
	config.Producer.Return.Successes = true
	config.Producer.Flush.Frequency = time.Second
	p,err := sarama.NewAsyncProducer(addrs,config)
	if err!=nil{
		return nil,err
	}
	return &producer{
		topic: topic,
		addrs: addrs,
		p:     p,
		input: make(chan string),
	},nil
}

func (p *producer)SendMessageAsync(input chan string){
	for message := range input{
		msg := &sarama.ProducerMessage{Topic: p.topic,Key:nil,Value: sarama.StringEncoder(message)}
		select {
		case p.p.Input()<-msg:
			metric.ProduceCount.Inc()
		case err :=<- p.p.Errors():
			metric.ProduceFailCount.Inc()
			log.Logger.Errorf("produce msg err,%s",err)
		case _,ok := <- p.p.Successes():
			if !ok{
				break
			}
			metric.ProduceSuccessCount.Inc()
		}
	}
}



func (p *producer)Close(){
	if p.p !=nil{
		err := p.p.Close()
		if err!=nil{
			log.Logger.Fatalf("close async producer err,%s",err)
		}
	}
	if p.p2 !=nil{
		err := p.p2.Close()
		if err!=nil{
			log.Logger.Fatalf("close sync producer err,%s",err)
		}
	}
}



// 同步方式，适合发送一些频率不高的信息

func NewSyncProducer(addrs []string,topic string)(*producer,error){
	config := sarama.NewConfig()
	config.Version = sarama.V2_5_0_0

	p,err := sarama.NewSyncProducer(addrs,config)
	if err!=nil{
		return nil,err
	}
	return &producer{
		topic: topic,
		addrs: addrs,
		p2:     p,
		input: make(chan string),
	},nil
}


func (p *producer)SendMessageSync(message string)error{

	msg := &sarama.ProducerMessage{Topic: p.topic,Key:nil,Value: sarama.StringEncoder(message)}
	partition,offset,err := p.p2.SendMessage(msg)
	if err!=nil{
		log.Logger.Errorf("sync send message err,%s",err)
		return err
	}
	log.Logger.Errorf("sync send message to partition %d,offset %d",partition,offset)
	return nil
}