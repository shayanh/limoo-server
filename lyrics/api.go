package lyrics

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shayanh/limoo-server/lyrics/parser"
)

type request struct {
	Artist string
	Title  string
}

type response struct {
	Lyrics string `json:"lyrics"`
}

func getLyrics(artist string, title string) (response, error) {
	// TODO normalize stringss

	resp := response{}

	t, found := searchTrack(artist, title)
	if found {
		resp.Lyrics = t.Lyrics
		return resp, nil
	}

	parsers := parser.GetParsers()
	for _, p := range parsers {
		lyrics, err := parser.GetLyrics(p, artist, title)
		if err == nil {
			resp.Lyrics = lyrics

			t := track{
				Artist: artist,
				Title:  title,
				Lyrics: lyrics,
			}
			go insertTrack(&t)

			return resp, nil
		}
	}
	return resp, fmt.Errorf("lyrics not found")
}

func HandleFuncs(router *mux.Router) {
	router.HandleFunc("", func(w http.ResponseWriter, r *http.Request) {
		req := request{
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
