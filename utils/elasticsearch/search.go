package elasticsearch

import (
	"context"
	es "gopkg.in/olivere/elastic.v5"
	"strings"
)

type SearchManager struct {
	svc *es.SearchService
}

func NewSearchManager() *SearchManager {
	return &SearchManager{
		svc: es.NewSearchService(Client),
	}
}

func (s *SearchManager) Search(r *SearchRequest)(*es.SearchResult,error){

	indices := s.GetIndicesByTime()

	result,err := s.svc.
		Type("data").
		Index(indices...).
		SortWithInfo(es.SortInfo{Field:TIMESTAMP, Ascending:false}).
		Query(s.GenerateQuery(r)).
		AllowNoIndices(true).
		IgnoreUnavailable(true).
		Do(context.TODO())

	if err!= nil {
		return nil,err
	}

	res := s.Wrap(result)
	return res,nil
}


// TODO  根据请求写query
func (s *SearchManager) GenerateQuery(r *SearchRequest) *es.BoolQuery {

	for _, item := range r.Match {
		query := es.NewBoolQuery()
		s.setEq(query,item.Eq)
		s.setRange(query,item.Range)
		s.setSimpleQueryStringQuery(query,item.SimpleQuery)
	}
	return nil
}

/*
{
  "query": {
    "bool": {
      "must": [
        { "match": { "title":   "Search"        }},
        { "match": { "content": "Elasticsearch" }}
      ],
      "filter": [
        { "term":  { "status": "published" }},
        { "range": { "publish_date": { "gte": "2015-01-01" }}}
      ]
    }
  }
}

{
  "query": {
    "bool" : {
      "must" : {
        "term" : { "user" : "kimchy" }
      },
      "filter": {
        "term" : { "tag" : "tech" }
      },
      "must_not" : {
        "range" : {
          "age" : { "gte" : 10, "lte" : 20 }
        }
      },
      "should" : [
        { "term" : { "tag" : "wow" } },
        { "term" : { "tag" : "elasticsearch" } }
      ],
      "minimum_should_match" : 1,
      "boost" : 1.0
    }
  }
}
*/

func (s *SearchManager) setEq (query *es.BoolQuery,eq map[string][]interface{}) {
	for field,filters := range eq{
		query.Must(es.NewTermsQuery(field,filters...))
	}
}

// 只作为filter
func (s *SearchManager) setRange (query *es.BoolQuery,r map[string]Range) {

	for field,filter := range r{

		rangeQuery := es.NewRangeQuery(field)

		if filter.Gte != nil  {
			rangeQuery = rangeQuery.Gte(filter.Gte)
		}

		if filter.Gt != nil {
			rangeQuery = rangeQuery.Gte(filter.Gt)
		}

		if filter.Lte != nil {
			rangeQuery = rangeQuery.Lte(filter.Lte)
		}

		if filter.Lt != nil {
			rangeQuery.Lt(filter.Lt)
		}

		query.Filter(rangeQuery)
	}
}


// 更复杂的语义 TODO  可在query 中 "query":"status:(500 or 501)"
func (s *SearchManager) setQueryStringQuery(query *es.BoolQuery){

	es.NewQueryStringQuery("")
}

// "originalMsg":"+Error +500 -192.168.0.1"
// 包含 Error 和 500 ,不包含192.168.0.1
func (s *SearchManager) setSimpleQueryStringQuery(query *es.BoolQuery,simple map[string]string){

	for field , rules := range simple {
		simpleQuery := es.NewSimpleQueryStringQuery(rules).
			Field(field).
			DefaultOperator("and")
		query.Must(simpleQuery)
	}
}


// "originalMsg": ["Error","500"]
// 某个字段是否包含这个词
// 可用于全文检索 _all
func (s *SearchManager) setSubString(query *es.BoolQuery,filter map[string][]string) {

	for field,rules := range filter {
		text := strings.Join(rules," ")
		simpleQuery := 	es.NewSimpleQueryStringQuery(text).
			Field(field).DefaultOperator("or")
		query.Must(simpleQuery)
	}

}

// TODO other query

func (s *SearchManager) GetIndicesByTime() []string{

	indices := make([]string,0)

	//TODO

	return indices
}



// TODO 根据需求写格式
func (s *SearchManager) Wrap(result *es.SearchResult) *es.SearchResult{

	return result
}


// 上下文
func (s *SearchManager) SearchContext(*es.SearchResult,error){

}

// scroll


func searilize() {

}