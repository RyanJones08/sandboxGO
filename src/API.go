package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
)

func main(){
	handleRequest()
}

func handleRequest(){
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8080",nil))
}

func homePage(w http.ResponseWriter, r *http.Request){
	json.NewEncoder(w).Encode("Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}