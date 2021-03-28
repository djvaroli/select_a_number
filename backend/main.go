package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)


func handleRequests(port string) {
	muxRouter := mux.NewRouter().StrictSlash(true)
	muxRouter.HandleFunc("/", Home)
	muxRouter.HandleFunc("/health", HealthCheck)
	muxRouter.HandleFunc("/es_health", EsCheck)
	muxRouter.HandleFunc("/submit_data", SubmitData).Methods("POST")
	log.Fatal(http.ListenAndServe(port, muxRouter))
}


func main() {
	port := ":5050"
	fmt.Println("Starting server on port", port)
	fmt.Println("View at http://localhost" + port)
	handleRequests(port)
}



