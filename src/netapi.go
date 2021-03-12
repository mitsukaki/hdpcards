package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type ScrapeRequest struct {
	MapName string
	Players []string
}

func scrape(w http.ResponseWriter, req *http.Request) {
	var body ScrapeRequest

	// decode the request
	err := json.NewDecoder(req.Body).Decode(&body)
	if err != nil {
		// throw error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}

	fmt.Printf("%+v\n", body)
	// load stats for all the Players
	for i, playerName := range body.Players {
		log.Println("Loading stats for " + playerName)
		loadStats(i, playerName, body.MapName)
	}

	// respond with response code 200
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(""))
}
