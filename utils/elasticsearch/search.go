package elasticsearch

import (
	"context"
	es "gopkg.in/olivere/elastic.v5"
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

	result,err := s.svc.
		Type().
		Index().
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
	}
	return nil
}

/*

"eq":{
	"name":["jack","tom"]
}

*/

func (s *SearchManager) setEq (query *es.BoolQuery,eq map[string][]interface{}) {
	for fields,filters := range eq{
		query.Must(es.NewTermsQuery(fields,filters...))
	}
}

// TODO other query


// TODO 根据需求写格式
func (s *SearchManager) Wrap(result *es.SearchResult) *es.SearchResult{

	return result
}


// 上下文
func (s *SearchManager) SearchContext(*es.SearchResult,error){

}

// scroll