package wordsearch

import (
	"fmt"
)

const testVersion = 3

// Puzzle holds a puzzle to be solved
type Puzzle []string

// Solve finds the specified words in the provided puzzle
func Solve(words []string, puzzle []string) (map[string][2][2]int, error) {
	solution := map[string][2][2]int{}

	p := Puzzle(puzzle)

	for _, w := range words {
		result, err := p.Search(w)
		if err != nil {
			return solution, err
		}
		solution[w] = result
	}

	return solution, nil
}

// Direction represents the directions to match words
type Direction int

// Direction constants for matching
const (
	E Direction = iota
	NE
	N
	NW
	W
	SW
	S
	SE
)

// Search finds a specific word in the Puzzle
func (puzzle Puzzle) Search(word string) ([2][2]int, error) {
	for r, row := range puzzle {
		for c, letter := range row {
			if letter == rune(word[0]) {
				for d := E; d < SE; d++ {
					if last, ok := puzzle.match(word, r, c, d); ok {
						return [2][2]int{[2]int{c, r}, last}, nil
					}
				}
			}
		}
	}

	return [2][2]int{}, fmt.Errorf("wordsearch: %q not found", word)
}

// match checks a Puzzle location for a word in a specific Direction
func (puzzle Puzzle) match(word string, r, c int, d Direction) (last [2]int, ok bool) {
	dx, dy := deltas(d)

	c2 := c + dx*(len(word)-1)
	if c2 < 0 || c2 >= len(puzzle[0]) {
		return
	}

	r2 := r + dy*(len(word)-1)
	if r2 < 0 || r2 >= len(puzzle) {
		return
	}

	for i := 1; i < len(word); i++ {
		if puzzle[r+i*dy][c+i*dx] != word[i] {
			return
		}
	}

	return [2]int{c2, r2}, true
}

// deltas is a helper to return the direction vector for matching in a specific Direction
func deltas(d Direction) (dx int, dy int) {
	switch d {
	case E:
		dx = 1
	case NE:
		dx = 1
		dy = -1
	case N:
		dy = -1
	case NW:
		dx = -1
		dy = -1
	case W:
		dx = -1
	case SW:
		dx = -1
		dy = 1
	case S:
		dy = 1
	case SE:
		dx = 1
		dy = 1
	}

	return
}
