package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func Home(w http.ResponseWriter, r *http.Request) {
	response := homeResponse{"Hello World", time.Now()}
	_ = json.NewEncoder(w).Encode(response)
}


func HealthCheck(w http.ResponseWriter, r *http.Request) {
	response := applicationHealthResponse{"Good", 1000, 3}
	_ = json.NewEncoder(w).Encode(response)
}


func SubmitData(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	_, _ = fmt.Fprintf(w, "%+v", string(reqBody))
}