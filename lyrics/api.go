package lyrics

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shayanh/limoo-server/lyrics/parser"
)

type Request struct {
	Artist string
	Title  string
}

type Response struct {
	Lyrics string `json:"lyrics"`
}

func getLyrics(artist string, title string) (Response, error) {
	// TODO normalize strings

	resp := Response{}

	parsers := parser.GetParsers()
	for _, p := range parsers {
		lyrics, err := parser.GetLyrics(p, artist, title)
		if err == nil {
			resp.Lyrics = lyrics
			return resp, nil
		}
	}
	return resp, fmt.Errorf("lyrics not found")
}

func HandleFuncs(router *mux.Router) {
	router.HandleFunc("", func(w http.ResponseWriter, r *http.Request) {
		req := Request{
			Artist: r.URL.Query().Get("artist"),
			Title:  r.URL.Query().Get("title"),
		}

		fmt.Println("artist =", req.Artist, "title =", req.Title)

		resp, err := getLyrics(req.Artist, req.Title)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
	})
}
