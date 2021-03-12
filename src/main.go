package main

import (
	"net/http"
)

func main() {
	// accept and handle terminal input in another thread
	go terminalProcess()

	// serve web UI and api
	http.HandleFunc("/api/scrape", scrape)
	http.Handle("/", http.FileServer(http.Dir("./static")))

	http.ListenAndServe(":8883", nil)
}
