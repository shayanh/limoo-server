package lyrics

import (
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
)

type Handler struct {
	backends []backend
}

type TrackInfo struct {
	Artist     string
	Title      string
	SongArtURL string
}

func (p *Handler) cleanName(s string) string {
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

func (p *Handler) Init(qartist, qtitle string) {
	qartist = p.cleanName(qartist)
	qtitle = p.cleanName(qtitle)

	gns := new(genius)
	gns.init(qartist, qtitle)

	p.backends = append(p.backends, gns)
}

func (p *Handler) GetTrackInfo() (TrackInfo, error) {
	for _, b := range p.backends {
		trackInfo, err := b.getTrackInfo()
		if err != nil {
			logrus.Error(err)
			continue
		}
		return trackInfo, nil
	}
	return TrackInfo{}, fmt.Errorf("cannot fetch track information")
}

func (p *Handler) GetLyrics() (string, error) {
	for _, b := range p.backends {
		lyrics, err := b.getLyrics()
		if err != nil {
			continue
		}
		return lyrics, nil
	}
	return "", fmt.Errorf("cannot find lyrics")
}
