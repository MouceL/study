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
	Start time.Time
	End time.Time
	Keyword string
	AppName string
	Match []Match
	Sort bool
	From int
	Size int
}

type Match struct {
	Eq map[string][]interface{}
}

var ArrayPool = sync.Pool{
	New: func()interface{}{
		return array.NewLimitarray(1000)
	},
}