package metric

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"lll/study/log"
	"net/http"
)

var (
	ConsumeTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name:"kafka_consume_total_count",
		Help: "kafka consume total count",
})
	ConsumeCount = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name:"kafka_consume_count",
		Help: "kafka consume count per source",
	},[]string{"source"})

	ProduceCount = prometheus.NewCounter(prometheus.CounterOpts{
		Name:"kafka_produce_count",
		Help: "kafka produce total count",
	})
	ProduceFailCount = prometheus.NewCounter(prometheus.CounterOpts{
		Name:"kafka_produce_failed_count",
		Help: "kafka produce failed count",
	})
	ProduceSuccessCount = prometheus.NewCounter(prometheus.CounterOpts{
		Name:"kafka_produce_success_count",
		Help: "kafka produce success count",
	})
)

func init(){
	prometheus.MustRegister(ConsumeTotal)
	prometheus.MustRegister(ConsumeCount)
	prometheus.MustRegister(ProduceCount)
	prometheus.MustRegister(ProduceFailCount)
	prometheus.MustRegister(ProduceSuccessCount)
}

func ExposeMetric(){
	go func() {
		http.Handle("/metric",promhttp.Handler())
		err:=http.ListenAndServe(":8888",nil)
		if err!=nil{
			log.Logger.Fatalf("start prometheus port err %s",err.Error())
		}
	}()
}