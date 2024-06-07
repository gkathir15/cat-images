package main

import (
	"embed"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

// simple http server with a handler that greets the user

//go:embed templates/*
var resources embed.FS

var t = template.Must(template.ParseFS(resources, "templates/*"))

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "9090"
	}

	log.Println("listening on", port)
	log.Fatal(http.ListenAndServe(":"+port, http.HandlerFunc(defaultHandler)))
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

	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		data := map[string]string{
			"Region": os.Getenv("FLY_REGION"),
		}

		t.ExecuteTemplate(w, "index.html.tmpl", data)
	})
}
func greet(w http.ResponseWriter) {
	w.Write([]byte("Hello, pass your subreddit in the url like this: localhost:9000/r/news"))
}
