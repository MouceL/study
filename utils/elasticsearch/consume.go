package elasticsearch

import (
	"context"
	"lll/study/utils/kafka"
)
// 产生待写入的 es 消息 ，TODO

type producer struct {

}

func produce() error {
	g ,err := kafka.NewGroup([]string{""},"","",1)
	if err!= nil{
		return err
	}
	g.Consume(context.TODO())

	for {
		select {
		case <- g.Output():

		}
	}


	return nil
}

