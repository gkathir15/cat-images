package main

import (
	"net/http"
	"strings"
)

// simple http server with a handler that greets the user

func main() {

	http.ListenAndServe(":9000", http.HandlerFunc(defaultHandler))
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path == "/greet" {
		greet(w)
		return
	}
	if strings.HasPrefix(r.URL.Path, "/r/") {
		getParsedSubredditData(w, r.URL.Path[3:])
		return
	}
	if r.URL.Path == "/" {
		w.Write([]byte("Hello, pass your subreddit in the url like this: localhost:9000/r/news"))
		return

	}
}
func greet(w http.ResponseWriter) {
	w.Write([]byte("Hello, pass your subreddit in the url like this: localhost:9000/r/news"))
}
