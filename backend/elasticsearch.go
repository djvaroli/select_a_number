package main

import (
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
)

func GetEsClient() *elasticsearch.Client {
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		fmt.Println(err)
	}
	return es
}
