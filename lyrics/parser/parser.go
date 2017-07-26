package parser

import (
	"fmt"
)

type Parser struct {
	backends []backend
}

func (p *Parser) Init(QArtist, QTitle string) {
	gns := new(genius)
	gns.init(QArtist, QTitle)

	p.backends = append(p.backends, gns)
}

func (p *Parser) GetTrackInfo() (string, string, error) {
	for _, b := range p.backends {
		artist, title, err := b.getTrackInfo()
		if err != nil {
			continue
		}
		return artist, title, nil
	}
	return "", "", fmt.Errorf("cannot fetch track information")
}

func (p *Parser) GetLyrics() (string, error) {
	for _, b := range p.backends {
		lyrics, err := b.getLyrics()
		if err != nil {
			continue
		}
		return lyrics, nil
	}
	return "", fmt.Errorf("cannot find lyrics")
}
