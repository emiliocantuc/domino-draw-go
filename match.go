package main

import (
	"fmt"
	"math/rand"
)

// Keeps track of the board, players, whose turn it is, and
// the number of consecutive turns that have resulted in a pass.
type Match struct {
	board             Board
	players           []Player
	turn              int
	consecutivePasses int
	verbose           bool
}

// Returns copy of a new match. The boneyard and player hands are
// randomly initialized. A player is chosen randomly to start.
// TODO the player with highest mule should start.
func NewMatch(players []Player, maxPip, initialHandSize int, verbose bool) (Match, error) {

	new := Match{verbose: verbose}
	new.board = NewBoard()
	new.players = players
	new.turn = rand.Intn(len(players)) // Random player starts

	// Place tiles in boneyard
	for a := 0; a <= maxPip; a++ {
		for b := a; b <= maxPip; b++ {
			new.board.boneyard = append(new.board.boneyard, NewTile(a, b))
		}
	}

	// Distribute cards to players randomly
	for p := range new.players {
		for i := 0; i < initialHandSize; i++ {
			drawn, err := new.board.drawTile()
			if err != nil {
				return new, fmt.Errorf("Something aint right")
			}
			AddTile(new.players[p], drawn)
		}
	}

	return new, nil
}

// Returns slice with player scores.
func (match *Match) Scores() []int {
	scores := make([]int, len(match.players))
	for i, p := range match.players {
		scores[i] = PlayerScore(p)
	}
	return scores
}

// True if no one can move or someone has won.
func (match *Match) IsOver() bool {

	// If all players have passed and no one can move
	if match.consecutivePasses > len(match.players) {
		return true
	}

	// If someone has score 0 and no tiles left
	for _, p := range match.players {
		if PlayerScore(p) == 0 && len(*p.Tiles()) == 0 {
			return true
		}
	}
	return false
}

// Plays out the match and returns final player scores.
func (match *Match) Play() []int {

	nDrawn := 0
	match.consecutivePasses = 0

	for !match.IsOver() {

		// Player w/ turn moves
		currentPlayer := match.players[match.turn]
		canMove, move := currentPlayer.MakeMove(match.board)

		if canMove {

			if match.verbose {
				fmt.Println(match.turn, "moves", move)
			}

			// Make the move
			RemoveTile(currentPlayer, move.tile)
			match.board.Place(move.tile, move.boardEnd)
			nDrawn = 0

		} else {

			// Draw tiles and add to player
			drawn := match.board.Draw()
			for _, drawnTile := range drawn {
				AddTile(currentPlayer, drawnTile)
			}
			nDrawn = len(drawn)

			// Try to make move w/ drawn tiles
			canMove, move = currentPlayer.MakeMove(match.board)

			if canMove {
				// Make the move
				RemoveTile(currentPlayer, move.tile)
				match.board.Place(move.tile, move.boardEnd)
			} else {
				// Pass
				match.consecutivePasses++
			}

			if match.verbose {
				fmt.Println(match.turn, "cannot move. drawed", drawn, "move", canMove, move)
			}
		}

		// Other players are informed of the move
		for _, p := range match.players {
			p.OpponentMove(match.turn, nDrawn, move.tile)
		}

		// The turn is passed
		match.turn = (match.turn + 1) % len(match.players)
	}
	return match.Scores()
}

// String representation. TODO pretty up
func (match Match) String() string {
	playersStr := "\n"
	for i, p := range match.players {
		playersStr += fmt.Sprintln("\tPlayer ", i, p)
	}
	return fmt.Sprintln("Board:\n", match.board, "\nPlayers:", playersStr, "\nScores:", match.Scores(), "Is Over:", match.IsOver())
}
