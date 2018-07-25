package lyrics

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"

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

func getLyrics(qartist, qtitle string) (response, error) {
	resp := response{}

	log.WithFields(log.Fields{
		"qartist": qartist,
		"qtitle":  qtitle,
	})

	p := new(parser.Parser)
	p.Init(qartist, qtitle)
	artist, title, err := p.GetTrackInfo()
	log.WithFields(log.Fields{
		"artist": artist,
		"title":  title,
	}).Info()
	if err != nil {
		return resp, err
	}

	t, found := searchTrack(artist, title)
	log.WithFields(log.Fields{
		"found in db": found,
	}).Info()
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

// HandleFuncs for /lyrics path
func HandleFuncs(router *mux.Router) {
	router.HandleFunc("", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		req := request{
			// Artist: r.URL.Query().Get("artist"),
			// Title:  r.URL.Query().Get("title"),
			Artist: r.PostForm.Get("artist"),
			Title:  r.PostForm.Get("title"),
		}

		resp, err := getLyrics(req.Artist, req.Title)
		if err != nil {
			log.WithFields(log.Fields{
				"artist": req.Artist,
				"title":  req.Title,
			}).Error(err)
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
	})
}
