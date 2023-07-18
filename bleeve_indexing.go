//
//Index any go data structure (including JSON)
//Intelligent defaults backed up by powerful configuration
//Supported field types:
//Text, Numeric, Datetime, Boolean
//Supported query types:
//Term, Phrase, Match, Match Phrase, Prefix, Fuzzy
//Conjunction, Disjunction, Boolean (must/should/must_not)
//Term Range, Numeric Range, Date Range
//Geo Spatial
//Simple query string syntax for human entry
//tf-idf Scoring
//Query time boosting
//Search result match highlighting with document fragments
//Aggregations/faceting support:
//Terms Facet
//Numeric Range Facet
//Date Range Facet

package main

import (
	"fmt"
	"github.com/blevesearch/bleve"
)

func main() {
	// Indexing
	//message := struct {
	//	Id   string
	//	From string
	//	Body string
	//}{
	//	Id:   "example",
	//	From: "marty.schoch@gmail.com",
	//	Body: "bleve indexing is easy",
	//}
	//
	//mapping := bleve.NewIndexMapping()
	//index, err := bleve.New("example.bleve", mapping)
	//if err != nil {
	//	panic(err)
	//}
	//index.Index(message.Id, message)

	// Querying
	index, _ := bleve.Open("example.bleve")
	query := bleve.NewQueryStringQuery("bleve")
	searchRequest := bleve.NewSearchRequest(query)
	searchResult, _ := index.Search(searchRequest)
	fmt.Println(searchResult)
	//query := bleve.NewMatchQuery("text")
	//search := bleve.NewSearchRequest(query)
	//searchResults, _ := index.Search(search)
	//fmt.Println(searchResults)
}
