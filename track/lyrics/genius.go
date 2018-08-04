package lyrics

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/sirupsen/logrus"

	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/spf13/viper"
)

type genius struct {
	accessToken     string
	qartist, qtitle string
	trackURL        string
}

func (g *genius) init(qartist, qtitle string) {
	g.accessToken = viper.GetString("genius_access_token")
	g.qartist = qartist
	g.qtitle = qtitle
}

type response struct {
	Meta    meta    `json:"meta"`
	Message message `json:"response"`
}

type meta struct {
	Status int `json:"status"`
}

type message struct {
	Hits []hit `json:"hits"`
}

type hit struct {
	Result result `json:"result"`
}

type result struct {
	FullTitle  string       `json:"full_title"`
	Title      string       `json:"title"`
	TrackURL   string       `json:"url"`
	SongArtURL string       `json:"song_art_image_thumbnail_url"`
	Artist     geniusArtist `json:"primary_artist"`
}

type geniusArtist struct {
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

var httpClient = &http.Client{Timeout: 5 * time.Second}

func (g *genius) getTrackInfo() (string, string, error) {
	apiURL := "https://api.genius.com/"
	method := "search"
	query := url.QueryEscape(g.qartist + " " + g.qtitle)
	reqURL := fmt.Sprintf("%s%s?q=%s", apiURL, method, query)

	req, _ := http.NewRequest("GET", reqURL, nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", g.accessToken))
	resp, err := httpClient.Do(req)
	if err != nil {
		logrus.Error(err)
		return "", "", err
	}
	defer resp.Body.Close()

	res := new(response)
	json.NewDecoder(resp.Body).Decode(res)

	status := res.Meta.Status
	hits := res.Message.Hits

	if status != 200 || len(hits) < 1 {
		return "", "", fmt.Errorf("cannot fetch track information")
	}
	g.trackURL = hits[0].Result.TrackURL

	artist := hits[0].Result.Artist.Name
	title := hits[0].Result.Title
	return artist, title, nil
}

func (g *genius) getLyrics() (string, error) {
	if g.trackURL == "" {
		_, _, err := g.getTrackInfo()
		if err != nil {
			return "", err
		}
	}
	doc, err := goquery.NewDocument(g.trackURL)
	if err != nil {
		return "", err
	}

	nodes := doc.Find(".lyrics")
	if nodes.Size() < 1 {
		return "", fmt.Errorf("Not found")
	}
	lyrics := strings.TrimSpace(nodes.Eq(0).Text())
	return lyrics, nil
}
