package main

import (
	"fmt"
	"math/rand"
)

// What all players should implement
type Player interface {
	Tiles() *[]Tile
	MakeMove(board Board) (bool, Move)
	OpponentMove(current_player, n_drawn int, tile Tile)
}

// Player methods

// Appends a tile to player's tiles.
func AddTile(p Player, tile Tile) {
	new := append(*p.Tiles(), tile)
	*p.Tiles() = new
}

// Removes tile (by checking matching sides) from player's tiles.
func RemoveTile(p Player, tile Tile) error {

	// First find the tile
	tiles := *p.Tiles()
	n := len(tiles)
	i := 0
	for ; i < n && (tiles[i].a != tile.a || tiles[i].b != tile.b); i++ {
	}

	// Remove if we found it
	if i < n {
		new := make([]Tile, n-1)
		copy(new, tiles[:i])
		copy(new[i:], tiles[i+1:])
		*p.Tiles() = new
		return nil

	} else {
		return fmt.Errorf("Tile to remove not in player tiles.")
	}
}

// Sum of pips in player's tiles
func PlayerScore(p Player) int {
	score := 0
	for _, t := range *p.Tiles() {
		score += t.a + t.b
	}
	return score
}

// Players

// An illustratory player that plays randomly
type RandomPlayer struct {
	tiles []Tile
}

// "Getter" for the player's tiles (is there are better way to do this?).
func (p *RandomPlayer) Tiles() *[]Tile {
	return &p.tiles
}

// Returns pointer to a new random player.
func NewRandomPlayer() *RandomPlayer {
	return &RandomPlayer{tiles: make([]Tile, 0)}
}

// Tries to make a move. Returns (canMove, Move) if it can make a move
// and which one.
func (p *RandomPlayer) MakeMove(board Board) (bool, Move) {

	// Get all legal moves
	validMoves := board.ValidMoves(p.tiles)

	if len(validMoves) == 0 {
		return false, Move{}
	}

	i := rand.Intn(len(validMoves))
	return true, validMoves[i]
}

// "Registers" an opponent's move and with how many tiles it drew, if any.
func (p *RandomPlayer) OpponentMove(currentPlayer, nDrawn int, tile Tile) {
	// The random player does not care
}
