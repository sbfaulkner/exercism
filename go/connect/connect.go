package connect

const testVersion = 3

// Stone represents a playing piece
type Stone byte

// define the standard pieces
const (
	WHITE Stone = 'O'
	BLACK Stone = 'X'
)

// Hex represents a place on the board so we can track visits
type Hex struct {
	stone   Stone
	visited bool
}

// Board is the collection of hexes for the playing area
type Board [][]Hex

// NewBoard creates a Board from a slice of strings containing the current board state
func NewBoard(board []string) *Board {
	hexes := make([][]Hex, len(board))

	for r := range hexes {
		hexes[r] = make([]Hex, len(board[r]))

		for c := range hexes[r] {
			hexes[r][c].stone = Stone(board[r][c])
		}
	}

	b := Board(hexes)

	return &b
}

// ResultOf determines the winner of the provided board state
func ResultOf(board []string) (string, error) {
	return NewBoard(board).Winner()
}

// Winner determines the winner of a given Board
func (b *Board) Winner() (string, error) {
	if b.checkDown(WHITE) {
		return string(WHITE), nil
	}

	if b.checkAcross(BLACK) {
		return string(BLACK), nil
	}

	return "", nil
}

// checkAcross checks for paths of a specific colour from the leftmost hexes to a righthand hex
func (b *Board) checkAcross(stone Stone) bool {
	for r := range *b {
		if b.followAcross(0, r, BLACK) {
			return true
		}
	}

	return false
}

// checkDown checks for paths of a specific colour from the topmost hexes to a bottom hex
func (b *Board) checkDown(stone Stone) bool {
	for c := range (*b)[0] {
		if b.followDown(c, 0, WHITE) {
			return true
		}
	}

	return false
}

// directions defines the possible vectors between hexes
var directions = []struct {
	dc, dr int
}{
	{0, -1},
	{1, -1},
	{1, 0},
	{0, 1},
	{-1, 1},
	{-1, 0},
}

// followAcross checks a specific hex for a colour and recursively follows a path across the board
func (b *Board) followAcross(c, r int, s Stone) bool {
	if !b.visit(c, r, s) {
		return false
	}

	if c == len((*b)[r])-1 {
		return true
	}

	for _, d := range directions {
		if b.followAcross(c+d.dc, r+d.dr, s) {
			return true
		}
	}

	return false
}

// followDown checks a specific hex for a colour and recursively follows a path down the board
func (b *Board) followDown(c, r int, s Stone) bool {
	if !b.visit(c, r, s) {
		return false
	}

	if r == len(*b)-1 {
		return true
	}

	for _, d := range directions {
		if b.followDown(c+d.dc, r+d.dr, s) {
			return true
		}
	}

	return false
}

// visit a hex provided the coordinate is valid, the stone is the right colour, and it has not already been visited
func (b *Board) visit(c, r int, s Stone) bool {
	if r < 0 || r >= len(*b) {
		return false
	}

	if c < 0 || c >= len((*b)[r]) {
		return false
	}

	h := &(*b)[r][c]

	if h.stone != s {
		return false
	}

	if h.visited {
		return false
	}

	h.visited = true

	return true
}
