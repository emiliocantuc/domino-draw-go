package main

import "testing"

func TestCanPlace(t *testing.T) {

	b := NewBoard()

	// Tile sides, end to place at and whether we should be allowed to
	tests := []struct {
		a, b, end int
		want      bool
	}{
		{0, 0, -1, true}, // board is empty
		{0, 0, 0, false}, // already placed
		{0, 1, 2, false}, // end does not match tile
		{0, 1, 1, false}, // end does not match board ends
		{0, 1, 0, true},  // end matches a side
		{1, 1, 0, false}, // end does not match tile
		{1, 1, 1, true},  // end matches
		{0, 3, 0, true},  // an end should still be 0
		{1, 3, 1, true},  // other end should be 1 (now both ends are 3,3)
		{1, 3, 1, false}, // already placed
		{1, 3, 3, false}, // already placed
		{3, 6, 3, true},
	}

	for _, ithTest := range tests {

		// Does not use t.Run as they run in parrallel
		got := b.CanPlace(NewTile(ithTest.a, ithTest.b), ithTest.end)

		if got != ithTest.want {
			t.Errorf("Placed %d, %d. Got %t want %t", ithTest.a, ithTest.b, got, ithTest.want)
		}

		// If we are supposed to place it we do
		if ithTest.want {
			b.Place(NewTile(ithTest.a, ithTest.b), ithTest.end)
		}
	}
}
