package main

import (
	"fmt"
	"math/rand"
)

// Consists of the two sides (pips) of the piece, a and b.
type Tile struct {
	a, b int
}

// Returns a copy of a new Tile with sides i, j.
func NewTile(i, j int) Tile {

	new := Tile{}
	if i <= j {
		new.a = i
		new.b = j
	} else {
		new.b = i
		new.a = j
	}
	return new
}

// Consists of a tile and an end or extreme of the
// board where to place it.
type Move struct {
	tile     Tile
	boardEnd int
}

// Manages the played tiles and boneyard.
// It also keeps track of the ends of the board.
type Board struct {

	// List of tiles
	tiles []Tile

	boneyard []Tile

	// Ends
	ends [2]int
}

// Returns copy of a new board with empty played tiles
// and boneyard. Ends start as -1.
func NewBoard() Board {
	new := Board{tiles: make([]Tile, 0), boneyard: make([]Tile, 0)}
	new.ends[0] = -1
	new.ends[1] = -1
	return new
}

// Returns the index of the board ends that
// match the passed end. Returns -1 otherwise.
func (b *Board) MatchingEnd(end int) int {

	if end == b.ends[0] {
		return 0
	}
	if end == b.ends[1] {
		return 1
	}
	return -1
}

// Given a slice of tiles, it returns a slice of
// valid Move instances that can currently be played.
func (b *Board) ValidMoves(tiles []Tile) []Move {

	valid := make([]Move, 0)

	for _, tile := range tiles {
		for _, end := range b.ends {
			if b.CanPlace(tile, end) {
				valid = append(valid, Move{tile: tile, boardEnd: end})
			}
		}
	}

	return valid
}

// Evaluates wether a tile can be placed on an end with
// the passed value.
func (b *Board) CanPlace(t Tile, end int) bool {

	// If empty, can place on either end (wich start as -1)
	if len(b.tiles) == 0 {
		return true
	}

	// End must be in tile
	if end != t.a && end != t.b {
		return false
	}

	// Check it is not already placed
	for _, tile := range b.tiles {
		if tile == t {
			return false
		}
	}

	// End must be in board ends
	if end == b.ends[0] || end == b.ends[1] {
		return true
	}

	return false
}

// Evaluates if a tile can be placed in any board end.
func (b *Board) CanPlaceAnywhere(t Tile) bool {
	return b.CanPlace(t, b.ends[0]) || b.CanPlace(t, b.ends[1])
}

// Places a tile on board end.
func (b *Board) Place(t Tile, end int) error {

	if !b.CanPlace(t, end) {
		return fmt.Errorf("Cannot place tile")
	}

	if len(b.tiles) == 0 {
		b.ends[0] = t.a
		b.ends[1] = t.b
	} else {

		// End index to replace
		endIndex := b.MatchingEnd(end)

		// Value to replace it with
		nonCommonSide := t.a
		if t.a == end {
			nonCommonSide = t.b
		}

		// Replace
		b.ends[endIndex] = nonCommonSide
	}

	// Add to tile list
	b.tiles = append(b.tiles, t)
	return nil
}

// Randomly draws a tile from the boneyard.
func (b *Board) drawTile() (Tile, error) {

	// If boneyard is empty return error
	if len(b.boneyard) == 0 {
		return Tile{}, fmt.Errorf("No tiles in boneyard to draw")
	}

	// Select tile at random
	i := rand.Intn(len(b.boneyard))
	drawn := b.boneyard[i]

	// Remove from slice
	new_boneyard := make([]Tile, len(b.boneyard)-1)
	copy(new_boneyard, b.boneyard[:i])
	copy(new_boneyard[i:], b.boneyard[i+1:])
	b.boneyard = new_boneyard[:]

	return drawn, nil
}

// Draws and returns tiles from boneyard until one can be played or empty.
func (b *Board) Draw() []Tile {

	tiles_drawn := make([]Tile, 0)

	if len(b.boneyard) == 0 {
		return tiles_drawn
	}

	drawn, err := b.drawTile()
	for !b.CanPlaceAnywhere(drawn) && err == nil {
		tiles_drawn = append(tiles_drawn, drawn)
		drawn, err = b.drawTile()
	}

	if err == nil {
		tiles_drawn = append(tiles_drawn, drawn)
	}

	return tiles_drawn
}
