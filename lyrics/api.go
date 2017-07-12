package lyrics

import "github.com/gorilla/mux"
import "net/http"
import "encoding/json"
import "fmt"

type Request struct {
	Artist string
	Title  string
}

type Response struct {
	Lyric string `json:"lyrics"`
}

func getLyrics(artist string, title string) (Response, error) {
	// TODO normalize strings

	resp := Response{
		Lyric: "Hello!",
	}
	return resp, nil
}

func HandleFuncs(router *mux.Router) {
	router.HandleFunc("", func(w http.ResponseWriter, r *http.Request) {
		req := Request{
			Artist: r.URL.Query().Get("artist"),
			Title:  r.URL.Query().Get("title"),
		}

		fmt.Println("artist =", req.Artist, "title =", req.Title)

		resp, err := getLyrics(req.Artist, req.Title)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
	})
}
