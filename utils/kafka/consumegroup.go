package kafka

import (
	"context"
	"github.com/Shopify/sarama"
	"lll/study/log"
	"lll/study/utils/metric"
	"sync"
	"time"
)

type group struct {
	nums int
	topic string
	groupId string
	output chan string
	cg []sarama.ConsumerGroup
}

func NewGroup(addrs []string,topic,groupId string ,nums int)(*group,error){

	config := sarama.NewConfig()
	config.Version = sarama.V2_5_0_0

	var group = &group{
		nums:    nums,
		topic:   topic,
		groupId: groupId,
		output:  make(chan string),
		cg:      make([]sarama.ConsumerGroup,0,nums),
	}

	for i:=0;i<nums;i++{
		consumerGroup,err := sarama.NewConsumerGroup(addrs,groupId,config)
		if err!=nil{
			return nil,err
		}
		log.Logger.Infof("NewConsumerGroup")
		group.cg = append(group.cg,consumerGroup)
	}
	return group,nil
}

func (g*group) Output()chan string{
	return g.output
}


func (g *group)Consume(ctx context.Context){
	go func() {
		var wg sync.WaitGroup
		for _,item := range g.cg{
			consumer := item
			wg.Add(1)
			go func() {
				defer wg.Done()
				for{
					handle := GroupHandler{
						output: g.output,
					}
					// `Consume` should be called inside an infinite loop, when a
					// server-side rebalance happens, the consumer session will need to be
					// recreated to get the new claims
					err := consumer.Consume(ctx,[]string{g.topic},handle)
					if err!=nil{
						log.Logger.Errorf("consume topic %s err :%s",g.topic,err)
					}
					log.Logger.Infof("consume topic %s",g.topic)
				}
			}()
		}
		wg.Wait()
	}()
}



type GroupHandler struct {
	output chan string
}

// Setup is run at the beginning of a new session, before ConsumeClaim.
func (g GroupHandler)Setup(sarama.ConsumerGroupSession) error{
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
// but before the offsets are committed for the very last time.
func (g GroupHandler) Cleanup(sarama.ConsumerGroupSession) error{
	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
// Once the Messages() channel is closed, the Handler must finish its processing
// loop and exit.
func (g GroupHandler)ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error{

	log.Logger.Infof(" member %s start consume topic %s , partition %d",session.MemberID(),claim.Topic(),claim.Partition())
	tick := time.NewTicker(1*time.Minute)
	var offset int64
	for{
		select {
		case msg,ok := <- claim.Messages():
			if !ok {
				return nil
			}
			g.output <- string(msg.Value)
			metric.ConsumeTotal.Inc()
			offset = msg.Offset
			session.MarkMessage(msg,"")
			log.Logger.Infof("get message :%s",string(msg.Value))
		case <- tick.C:
			log.Logger.Infof("member %s ,topic %s , partition %d, offset %d",session.MemberID(),claim.Topic(),claim.Partition(),offset)
		}
	}
}