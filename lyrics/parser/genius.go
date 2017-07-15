package parser

import "github.com/gosimple/slug"
import "fmt"

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

func (g *Genius) parse(url string) (string, error) {
	fmt.Println(url)
	return "Genius parse works!", nil
}
