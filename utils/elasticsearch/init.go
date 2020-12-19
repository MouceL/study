package elasticsearch

import (
	es "gopkg.in/olivere/elastic.v5"
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
	index string // 要写到哪个索引
	raw string
}

type SearchRequest struct {
	start time.Time
	end time.Time
	keyword string
	appName string
}