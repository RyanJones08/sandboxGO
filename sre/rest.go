package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type server struct{}

func param(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fullName := vars["fullName"]
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(fmt.Sprintf(`{"message": "hello %s"}`, fullName)))
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/{fullName}", param)
	log.Fatal(http.ListenAndServe(":8080", r))
}
