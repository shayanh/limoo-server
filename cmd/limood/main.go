package main

import (
	"net/http"

	"github.com/shayanh/limoo-server/config"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
	psh "github.com/platformsh/gohelper"
	"github.com/shayanh/limoo-server/track"
)

func main() {
	p, err := psh.NewPlatformInfo()
	if err != nil {
		log.Fatal("Not in a Platform.sh Environemnt")
	}

	if err = config.InitEnvVars(); err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter()
	router.StrictSlash(true)

	track.HandleFuncs(router.PathPrefix("/lyrics").Subrouter())

	err = http.ListenAndServe(":"+p.Port, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Info(r.Method, r.URL)
		router.ServeHTTP(w, r)
	}))

	if err != nil {
		log.Fatal(err)
	}
}
