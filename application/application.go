package application

import (
	"log"
	"net/http"

	"github.com/spf13/viper"
)

type Application struct {
	ListenAddr string
}

func NewApplication() *Application {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	app := &Application{}
	app.ListenAddr = viper.GetString("addr")
	return app
}

func Logger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.URL)
		h.ServeHTTP(w, r)
	})
}
