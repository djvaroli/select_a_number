package main

import (
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7/esutil"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func Home(w http.ResponseWriter, r *http.Request) {
	response := homeResponse{"Hello World", time.Now()}
	_ = json.NewEncoder(w).Encode(response)
}

func EsCheck(w http.ResponseWriter, r *http.Request) {
	es := GetEsClient()
	res, err := es.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	fmt.Println(res)
	defer res.Body.Close()
}


func HealthCheck(w http.ResponseWriter, r *http.Request) {
	response := applicationHealthResponse{"Good", 1000, 3}
	_ = json.NewEncoder(w).Encode(response)
}


func SubmitData(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	var userInfo SubmittedUserInfo
	unmarshalErr := json.Unmarshal(reqBody, &userInfo)

	if unmarshalErr != nil {
		log.Printf("error decoding response: %v", unmarshalErr)
		if e, ok := unmarshalErr.(*json.SyntaxError); ok {
			log.Printf("syntax error at byte offset %d", e.Offset)
		}
		log.Printf("request body: %q", reqBody)
	}

	// store the data in ElasticSearch
	es := GetEsClient()
	res, _ := es.Index("test-index", esutil.NewJSONReader(&userInfo))
	fmt.Println(res)
}