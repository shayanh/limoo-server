package parser

import (
	"fmt"
	"strings"
)

type Parser struct {
	backends []backend
}

func (p *Parser) cleanName(s string) string {
	badChars := []rune{'-', '(', '[', ')', ']'}
	res := s
	for i, c := range s {
		flag := false
		for _, bad := range badChars {
			if c == bad {
				res = s[:i]
				flag = true
				break
			}
		}
		if flag {
			break
		}
	}
	res = strings.TrimSpace(res)
	return res
}

func (p *Parser) Init(qartist, qtitle string) {
	qartist = p.cleanName(qartist)
	qtitle = p.cleanName(qtitle)

	gns := new(genius)
	gns.init(qartist, qtitle)

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
