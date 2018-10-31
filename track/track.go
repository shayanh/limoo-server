package track

import (
	"google.golang.org/appengine"

	"google.golang.org/appengine/datastore"
)

type track struct {
	Artist     string `json:"artist"`
	Title      string `json:"title"`
	Lyrics     string `json:"lyrics"`
	SongArtURL string `json:"song_art_url"`
	Source     string `json:"source"`
}

func searchTrack(artist, title string) (*track, bool) {
	ctx := appengine.BackgroundContext()
	q := datastore.NewQuery("tracks").Filter("Artist =", artist).Filter("Title = ", title)
	t := q.Run(ctx)
	for {
		var tck track
		_, err := t.Next(&tck)
		if err == datastore.Done || err != nil {
			return nil, false
		}
		return &tck, true
	}
}

func insertTrack(t *track) error {
	ctx := appengine.BackgroundContext()
	_, err := datastore.Put(ctx, datastore.NewIncompleteKey(ctx, "tracks", nil), t)
	return err
}
