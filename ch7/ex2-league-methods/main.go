package main

import (
	"fmt"
	"sort"
)

// Exercise 7.2:
// Add two methods to League. MatchResult takes four parameters: the name of
// the first team, its score, the name of the second team, and its score. This
// method should update the Wins field. Add a second method called Ranking that
// returns a slice of the team names in order of wins.

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

// MatchResult records a game outcome and updates the win map.
func (l *League) MatchResult(team1, team2 string, score1, score2 int) {
	if score1 > score2 {
		l.Wins[team1]++
	} else if score2 > score1 {
		l.Wins[team2]++
	}
}

// Ranking returns team names sorted by wins (descending).
func (l *League) Ranking() []string {
	tNames := []string{}
	for team := range l.Wins {
		tNames = append(tNames, team)
	}

	sort.Slice(tNames, func(i, j int) bool {
		return l.Wins[tNames[i]] > l.Wins[tNames[j]]
	})
	return tNames
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

	// Record a match result and show updated ranking.
	league.MatchResult("Lakers", "Warriors", 110, 105)
	fmt.Println("\nAfter Lakers beat Warriors:")
	fmt.Println("Ranking:", league.Ranking())
}
