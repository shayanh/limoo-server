package main

import (
	"log"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello!"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", sayHello)

	log.Fatalln(http.ListenAndServe(":80", mux))
}
