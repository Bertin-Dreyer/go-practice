package main

import (
	"io"
	"os"
)

// Exercise 7.3:
// Define an interface called Ranker that has a single method called Ranking
// that returns a slice of strings. Write a function called RankPrinter with two
// parameters, the first of type Ranker and the second of type io.Writer.
// Use the io.WriteString function to write the values returned by Ranker to the
// io.Writer, with a newline separating each result. Call this function from main.

// Ranker is an interface with one method: Ranking() []string.
// Any type that has a Ranking() method returning []string automatically
// satisfies Ranker — Go interfaces are satisfied implicitly (no "implements"
// keyword needed).
type Ranker interface {
	Ranking() []string
}

// RankPrinter takes a Ranker and an io.Writer.
// It calls Ranking() on the Ranker, then writes each string to the Writer,
// separated by newlines.
//
// Parameters:
//   r - anything that satisfies the Ranker interface
//   i - destination for the output (e.g., os.Stdout, a file, a buffer)
func RankPrinter(r Ranker, i io.Writer) {
	rSlice := r.Ranking()

	for _, we := range rSlice {
		io.WriteString(i, we+"\n")
	}
}

func main() {
	// Create a basketballLeague value. Its pointer (*basketballLeague) has a
	// Ranking() method, so it satisfies the Ranker interface.
	league := &basketballLeague{
		teams: []string{"Lakers", "Warriors", "Celtics", "Bulls"},
	}

	// Pass the league (as a Ranker) and os.Stdout to RankPrinter.
	// Prints each team name on its own line.
	RankPrinter(league, os.Stdout)
}

// basketballLeague is a concrete type that holds a list of team names.
// It implements Ranker because *basketballLeague has a Ranking() method.
type basketballLeague struct {
	teams []string
}

// Ranking returns the list of team names.
// This is the method that makes *basketballLeague satisfy the Ranker interface.
func (l *basketballLeague) Ranking() []string {
	return l.teams
}
