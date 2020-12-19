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

func (s *SearchManager) Search()(*es.SearchResult,error){

	result,err := s.svc.Type().SortWithInfo(es.SortInfo{Field:"", Ascending:false}).
		Query(es.NewQueryStringQuery("")).
		AllowNoIndices(true).
		Do(context.TODO())
	return result,err

}

// 上下文
func (s *SearchManager) SearchContext(*es.SearchResult,error){

}

// scroll