package main

import "fmt"

// Exercise 7.1:
// Define two types. Team has a field for the name of the team and a field for
// the player names. League has a field called Teams for the teams in the league
// and a field called Wins that maps a team's name to its number of wins.

// Team holds a basketball team's name and roster.
type Team struct {
	Name   string
	Players []string
}

// League holds the teams in the league and each team's win count.
type League struct {
	Teams []Team
	Wins  map[string]int
}

func main() {
	league := League{
		Teams: []Team{
			{Name: "Lakers", Players: []string{"LeBron", "AD", "Russell"}},
			{Name: "Warriors", Players: []string{"Steph", "Klay", "Andrew"}},
			{Name: "Celtics", Players: []string{"Jayson", "Jaylen", "Marcus"}},
			{Name: "Bulls", Players: []string{"Zach", "DeMar", "Nikola"}},
		},
		Wins: map[string]int{
			"Lakers":   42,
			"Warriors": 45,
			"Celtics":  48,
			"Bulls":    38,
		},
	}

	// Print each team with its win count.
	for _, t := range league.Teams {
		fmt.Println(t.Name, "players:", len(t.Players), "wins:", league.Wins[t.Name])
	}
}
