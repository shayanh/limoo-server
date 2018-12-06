package main

import (
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/gorilla/mux"
	psh "github.com/platformsh/gohelper"
	"github.com/shayanh/limoo-server/application"
	"github.com/shayanh/limoo-server/track"
)

func main() {
	p, err := psh.NewPlatformInfo()
	if err != nil {
		log.Fatal("Not in a Platform.sh Environemnt")
	}

	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter()
	router.StrictSlash(true)

	track.HandleFuncs(router.PathPrefix("/lyrics").Subrouter())

	err = http.ListenAndServe(":"+p.Port, application.Logger(router))
	if err != nil {
		log.Fatal(err)
	}
}
