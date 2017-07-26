package main

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
	"github.com/shayanh/limoo-server/application"
	"github.com/shayanh/limoo-server/lyrics"
)

func main() {
	app := application.NewApplication()

	router := mux.NewRouter()
	router.StrictSlash(true)

	lyrics.HandleFuncs(router.PathPrefix("/lyrics").Subrouter())

	if err := lyrics.InitDB(app.MongoAddr); err != nil {
		log.Fatal(err)
	}
	defer lyrics.CloseDB()
	log.Printf("Connected to MongoDB")

	log.Printf("listening on http://%s/", app.ListenAddr)
	err := http.ListenAndServe(app.ListenAddr, application.Logger(router))
	if err != nil {
		log.Fatal(err)
	}
}
