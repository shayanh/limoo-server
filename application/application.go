package application

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/viper"
)

type Application struct {
	ListenAddr string
	MongoAddr  string
}

func NewApplication() *Application {
	viper.SetConfigName("config")
	viper.AddConfigPath("../../")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	app := &Application{}
	app.ListenAddr = viper.GetString("addr")
	app.MongoAddr = viper.GetString("mongo_addr")
	return app
}

func Logger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Info(r.Method, r.URL)
		h.ServeHTTP(w, r)
	})
}
