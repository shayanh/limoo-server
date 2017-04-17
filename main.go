package main

import (
	"log"
	"net/http"
	"os"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello!"))
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", sayHello)

	log.Fatalln(http.ListenAndServe(":"+port, mux))
}
