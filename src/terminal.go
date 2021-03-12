package main

import (
	"fmt"
	"strings"
)

func terminalProcess() {
	var playerList string
	var mapName string

	for true {
		fmt.Println("------------------")

		// read in map name
		fmt.Print("Enter map name: ")
		fmt.Scanln(&mapName)
		fmt.Println("------------------")

		// read in the player list
		fmt.Print("Enter player name(s), seperated by commas (no spaces): ")
		fmt.Scanln(&playerList)

		// split into list of players
		players := strings.Split(playerList, ",")

		// research each player
		for i, playerName := range players {
			fmt.Println("------------------")
			loadStats(i, playerName, mapName)
		}

		// terminate output
		fmt.Println("------------------")
		fmt.Println("# Stat files updated.")
	}
}
