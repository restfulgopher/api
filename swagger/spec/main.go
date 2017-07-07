package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "swagger.yaml")
	})

	log.Println("listening...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("could not serve on port 8080: %s", err)
	}
}
