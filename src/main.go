package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func loadstats(playerName string, mapName string) {
	// grab dummy page
	res, err := http.Get("https://dashleague.games/players/" + playerName + "/")
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// write their name
	err = filedump("name", playerName)
	if err != nil {
		log.Fatal(err)
	}

	// get their team name and write it
	var team string
	doc.Find(".player__tag--team").Each(func(i int, s *goquery.Selection) {
		team = s.Find("span").Text()

		err := filedump("team", team)
		if err != nil {
			log.Fatal(err)
		}
	})

	// write the map name
	err = filedump("mapname", mapName)
	if err != nil {
		log.Fatal(err)
	}

	// write their map play count
	matches := doc.Find(".matches").Find("tbody").Children().Length() - 1
	err = filedump("match_count", strconv.Itoa(matches))
	if err != nil {
		log.Fatal(err)
	}

	// write their total kills
	kills := doc.Find(".stats__stat").First().Clone().Children().Remove().End().Text() // HACK: remove children and get the kill number
	kills = strings.TrimSpace(kills)                                                   // HACK: remove the newlines
	err = filedump("total_kills", kills)
	if err != nil {
		log.Fatal(err)
	}

	// TODO: write their personal best kills

	// TODO: write their average kills
}

func filedump(label string, value string) error {
	data := []byte(value)

	// print to console
	fmt.Println(label + ": " + value)

	// write to file
	return ioutil.WriteFile("data/"+label+".txt", data, 0644)
}

func main() {
	// name of the player
	var player string

	var mapName string

	for true {
		fmt.Println("------------------")

		// read in the player name
		fmt.Print("Enter player name: ")
		fmt.Scanln(&player)

		// read in map name
		fmt.Print("Enter map name: ")
		fmt.Scanln(&mapName)
		fmt.Println("------------------")

		// dump their stats to files
		loadstats(player, mapName)

		// terminate output
		fmt.Println("------------------")
		fmt.Println("# Stat files updated.")
	}
}
