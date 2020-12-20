package elasticsearch

import (
	es "gopkg.in/olivere/elastic.v5"
	array "lll/study/utils/limitarray"
	"sync"
	"time"
)

var Client *es.Client

func GetEsClient(url,user,password string) (*es.Client,error) {
	if Client != nil {
		return Client,nil
	}
	return es.NewClient(es.SetURL(url),
		es.SetBasicAuth(user,password),
		es.SetHealthcheck(false),
	)
}

type IndexRequest struct {
	Index string
	Raw string
}

type SearchRequest struct {
	start time.Time
	end time.Time
	keyword string
	appName string
}

var ArrayPool = sync.Pool{
	New: func()interface{}{
		return array.NewLimitarray(1000)
	},
}