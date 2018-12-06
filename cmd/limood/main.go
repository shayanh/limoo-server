package main

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
	"github.com/shayanh/limoo-server/application"
	"github.com/shayanh/limoo-server/track"
)

func main() {
	app := application.NewApplication()

	// p, err := psh.NewPlatformInfo()
	// if err != nil {
	// log.Fatal("Not in a Platform.sh Environemnt")
	// }

	router := mux.NewRouter()
	router.StrictSlash(true)

	track.HandleFuncs(router.PathPrefix("/lyrics").Subrouter())

	// if err := track.InitDB(app.MongoAddr); err != nil {
	// 	log.Fatal(err)
	// }
	// defer track.CloseDB()
	// log.Printf("Connected to MongoDB")

	log.Printf("listening on http://%s/", app.ListenAddr)
	// err = http.ListenAndServe(":"+p.Port, application.Logger(router))
	err := http.ListenAndServe(app.ListenAddr, application.Logger(router))
	if err != nil {
		log.Fatal(err)
	}
}
