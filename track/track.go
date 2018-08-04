package track

import "gopkg.in/mgo.v2/bson"

type track struct {
	Artist string `json:"artist"`
	Title  string `json:"title"`
	Lyrics string `json:"lyrics"`
	Source string `json:"source"`
}

func searchTrack(artist, title string) (*track, bool) {
	t := new(track)
	found := false

	session, err := newSession()
	if err != nil {
		return t, false
	}
	defer session.Close()

	c := session.DB("limoo").C("tracks")
	err = c.Find(bson.M{"artist": artist, "title": title}).One(t)
	if err == nil {
		found = true
	}
	return t, found
}

func insertTrack(t *track) error {
	session, err := newSession()
	if err != nil {
		return err
	}
	defer session.Close()

	c := session.DB("limoo").C("tracks")
	err = c.Insert(t)
	return err
}
