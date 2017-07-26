package parser

type backend interface {
	init(qartist, qtitle string)
	getTrackInfo() (string, string, error)
	getLyrics() (string, error)
}
