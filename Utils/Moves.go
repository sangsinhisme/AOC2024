package Utils

import "slices"

// Move is a (row, col) offset. DR changes the row (up/down), DC changes the column (left/right).
type Move struct {
	DR, DC int
}

// Apply returns the new (row, col) after taking this move from (r, c).
func (m Move) Apply(r, c int) (int, int) {
	return r + m.DR, c + m.DC
}

// Slice versions: use these with `for _, mv := range Moves4 { ... }`
// Order is clockwise starting from right: right, down, left, up.
var (
	Moves4 = []Move{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	// Same 4 moves plus the diagonals.
	Moves8 = []Move{
		{0, 1}, {1, 1}, {1, 0}, {1, -1},
		{0, -1}, {-1, -1}, {-1, 0}, {-1, 1},
	}
)

// Iterator versions (Go 1.23+): `for mv := range Move4Steps { ... }`
// mv is the Move itself, not an index.
var (
	Move4Steps = slices.Values(Moves4)
	Move8Steps = slices.Values(Moves8)
)

// UDLR maps the characters used in puzzle inputs to a Move.
var UDLR = map[byte]Move{
	'U': {-1, 0},
	'D': {1, 0},
	'L': {0, -1},
	'R': {0, 1},
}

// InBounds reports whether (r, c) is inside a grid with n rows and m columns.
func InBounds(r, c, n, m int) bool {
	return r >= 0 && r < n && c >= 0 && c < m
}
