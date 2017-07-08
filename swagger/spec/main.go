package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", indexHandler)
	http.Handle("/", r)

	log.Println("listening...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("could not serve on port 8080: %s", err)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "swagger.yaml")
}
