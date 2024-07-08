package main

import (
	"log"
	"net/http"
)

func main() {
	// Start a simple HTTP server on localhost:8080
	// You can send requests using `curl localhost:8080`
	addr := ":8080"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Try changing this response with Lekko!"))
	})
	log.Println("Listening on", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
