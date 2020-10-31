package main

import (
	"fmt"
	"lll/study/utils/metric"
)

func main(){
	fmt.Println("hello world")
	metric.ExposeMetric()
	//metric.ConsumeTotal.Inc()
	stop := make(chan struct {})
	<-stop
}