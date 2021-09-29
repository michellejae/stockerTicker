package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "AYO did it!")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", homeLink)

	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
