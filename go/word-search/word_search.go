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
type Direction struct {
	dx, dy int
}

var directions = []Direction{
	{1, 0},
	{1, -1},
	{0, -1},
	{-1, -1},
	{-1, 0},
	{-1, 1},
	{0, 1},
	{1, 1},
}

// Search finds a specific word in the Puzzle
func (puzzle Puzzle) Search(word string) ([2][2]int, error) {
	for r, row := range puzzle {
		for c, letter := range row {
			if letter == rune(word[0]) {
				for _, d := range directions {
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
	c2 := c + d.dx*(len(word)-1)
	if c2 < 0 || c2 >= len(puzzle[0]) {
		return
	}

	r2 := r + d.dy*(len(word)-1)
	if r2 < 0 || r2 >= len(puzzle) {
		return
	}

	for i := 1; i < len(word); i++ {
		if puzzle[r+i*d.dy][c+i*d.dx] != word[i] {
			return
		}
	}

	return [2]int{c2, r2}, true
}
