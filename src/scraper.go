package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func loadStats(entry int, playerName string, mapName string) {
	// grab player stat page
	res, err := http.Get("https://dashleague.games/players/" + playerName + "/")
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// write their name
	err = filedump(entry, "name", playerName)
	if err != nil {
		log.Fatal(err)
	}

	// get their team name and write it
	var team string
	doc.Find(".player__tag--team").Each(func(i int, s *goquery.Selection) {
		team = s.Find("span").Text()

		err := filedump(entry, "team", team)
		if err != nil {
			log.Fatal(err)
		}
	})

	// write the map name
	err = filedump(entry, "mapname", mapName)
	if err != nil {
		log.Fatal(err)
	}

	// write their map play count
	matches := doc.Find(".matches").Find("tbody").Children().Length() - 1
	err = filedump(entry, "match_count", strconv.Itoa(matches))
	if err != nil {
		log.Fatal(err)
	}

	// write their total kills
	kills := doc.Find(".stats__stat").First().Clone().Children().Remove().End().Text() // HACK: remove children and get the kill number
	kills = strings.TrimSpace(kills)                                                   // HACK: remove the newlines
	err = filedump(entry, "total_kills", kills)
	if err != nil {
		log.Fatal(err)
	}

	// TODO: write their personal best kills

	// TODO: write their average kills
}

func filedump(entry int, label string, value string) error {
	data := []byte(value)

	// print to console
	fmt.Println(label + ": " + value)

	// write to file
	folderName := strconv.Itoa(entry)
	os.MkdirAll("data/"+folderName+"/", 0777)
	return ioutil.WriteFile("data/"+folderName+"/"+label+".txt", data, 0644)
}
