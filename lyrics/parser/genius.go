package parser

import (
	"fmt"

	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gosimple/slug"
)

type Genius struct {
}

func (g *Genius) getName() string {
	return "Genius"
}

func (g *Genius) generateURL(artist, title string) (string, error) {
	s := slug.Make(artist + " " + title + " " + "lyrics")
	url := "https://genius.com/" + s
	return url, nil
}

func (g *Genius) parse(trackURL string) (string, error) {
	doc, err := goquery.NewDocument(trackURL)
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
