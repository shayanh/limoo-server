package parser

type Parser interface {
	getName() string
	generateURL(artist, title string) (string, error)
	parse(url string) (string, error)
}

var (
	_parsers []Parser
)

func GetParsers() []Parser {
	if _parsers != nil {
		return _parsers
	}

	genius := &Genius{}

	parsers := []Parser{genius}
	_parsers = parsers
	return _parsers
}

func GetLyrics(p Parser, artist, title string) (string, error) {
	url, err := p.generateURL(artist, title)
	if err != nil {
		return "", err
	}
	lyrics, err := p.parse(url)
	return lyrics, err
}
