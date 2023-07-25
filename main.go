package main

import "fmt"

func main() {

	// A simple match between 2 random-playing players,
	// double 6 domino set, and an initial hand of 7 tiles.

	players := []Player{NewRandomPlayer(), NewRandomPlayer()}
	match, _ := NewMatch(players, 6, 7, true)
	fmt.Println(match)
	scores := match.Play()
	fmt.Println(scores)
}
