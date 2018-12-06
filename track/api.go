package track

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
	"github.com/shayanh/limoo-server/track/lyrics"
)

type request struct {
	Artist string
	Title  string
}

func getTrack(qartist, qtitle string) (track, error) {
	log.WithFields(log.Fields{
		"qartist": qartist,
		"qtitle":  qtitle,
	}).Info()

	lyricsHandler := new(lyrics.Handler)
	lyricsHandler.Init(qartist, qtitle)
	trackInfo, err := lyricsHandler.GetTrackInfo()
	if err != nil {
		return track{}, err
	}
	log.WithFields(log.Fields{
		"artist": trackInfo.Artist,
		"title":  trackInfo.Title,
	}).Info()

	// t, found := searchTrack(trackInfo.Artist, trackInfo.Title)
	// log.WithFields(log.Fields{"found in db": found}).Info()
	// if found {
	// return *t, nil
	// }

	lyrics, err := lyricsHandler.GetLyrics()
	if err != nil {
		return track{}, err
	}
	newTrack := &track{
		Artist:     trackInfo.Artist,
		Title:      trackInfo.Title,
		SongArtURL: trackInfo.SongArtURL,
		Lyrics:     lyrics,
	}
	// go insertTrack(newTrack)

	return *newTrack, nil
}

// HandleFuncs for /lyrics path
func HandleFuncs(router *mux.Router) {
	router.HandleFunc("", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		req := request{
			Artist: r.URL.Query().Get("artist"),
			Title:  r.URL.Query().Get("title"),
			// Artist: r.PostForm.Get("artist"),
			// Title:  r.PostForm.Get("title"),
		}

		resp, err := getTrack(req.Artist, req.Title)
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
