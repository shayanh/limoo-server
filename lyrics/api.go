package lyrics

import (
	"encoding/json"
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

func getLyrics(qartist string, qtitle string) (response, error) {
	resp := response{}

	p := new(parser.Parser)
	p.Init(qartist, qtitle)
	artist, title, err := p.GetTrackInfo()
	if err != nil {
		return resp, err
	}

	t, found := searchTrack(artist, title)
	if found {
		resp.Lyrics = t.Lyrics
		return resp, nil
	}

	lyrics, err := p.GetLyrics()
	if err != nil {
		return resp, err
	}
	newTrack := &track{
		Artist: artist,
		Title:  title,
		Lyrics: lyrics,
	}
	go insertTrack(newTrack)

	resp.Lyrics = lyrics
	return resp, nil
}

func HandleFuncs(router *mux.Router) {
	router.HandleFunc("", func(w http.ResponseWriter, r *http.Request) {
		req := request{
			Artist: r.URL.Query().Get("artist"),
			Title:  r.URL.Query().Get("title"),
		}

		resp, err := getLyrics(req.Artist, req.Title)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
	})
}
