package lyrics

type backend interface {
	init(qartist, qtitle string)
	getTrackInfo() (TrackInfo, error)
	getLyrics() (string, error)
}
