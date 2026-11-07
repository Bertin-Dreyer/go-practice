package main

import (
	"io"
	"os"
)

// Exercise 7.3:
// Define an interface called Ranker that has a single method called Ranking
// that returns a slice of strings. Write a function called RankPrinter with two
// parameters, the first of type Ranker and the second of type io.Writer.
// Use the io.Write
// String function to write the values returned by Ranker to the io.Writer, with a
// newline separating each result. Call this function from main.
type Ranker interface {
    Ranking() []string
}

func RankPrinter(r Ranker, i io.Writer){
	rSlice := r.Ranking()

	for _, we := range rSlice {
		io.WriteString(i, we + "\n")
	}
}
func main() {
	league := &basketballLeague{
		teams: []string{"Lakers", "Warriors", "Celtics", "Bulls"},
	}
	RankPrinter(league, os.Stdout)
}

type basketballLeague struct {
	teams []string
}

func (l *basketballLeague) Ranking() []string {
	return l.teams
}
