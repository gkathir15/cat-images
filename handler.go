package main

import (
	"encoding/json"
	"net/http"
)

func getParsedSubredditData(w http.ResponseWriter, subReddit string) {
	println("URL:", subReddit)

	entries, err := getFeedEntries(subReddit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&entries)
}
