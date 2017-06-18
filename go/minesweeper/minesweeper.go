package minesweeper

import (
	"errors"
	"regexp"
)

const testVersion = 1

// Board errors
var (
	ErrInvalidWidth          = errors.New("minesweeper: invalid board width")
	ErrInvalidCorner         = errors.New("minesweeper: invalid corner")
	ErrInvalidHorizontalEdge = errors.New("minesweeper: invalid horizontal edge")
	ErrInvalidVerticalEdge   = errors.New("minesweeper: invalid vertical edge")
	ErrInvalidCell           = errors.New("minesweeper: invalid cell")
)

var (
	validEdge = regexp.MustCompile(`^\+-*\+$`)
	validRow  = regexp.MustCompile(`^\|[ *1-8]*\|$`)
)

// Count marks the mine count for all spaces on the board
func (b *Board) Count() error {
	l, w := len(*b), len((*b)[0])

	for r, row := range *b {
		if len(row) != w {
			return ErrInvalidWidth
		}

		for c, cell := range row {
			switch cell {
			case '+':
				if r != 0 && r != l-1 || c != 0 && c != w-1 {
					return ErrInvalidCell
				}
			case '-':
				if r != 0 && r != l-1 || c == 0 || c == w-1 {
					return ErrInvalidCell
				}
			case '|':
				if c != 0 && c != w-1 || r == 0 || r == l-1 {
					return ErrInvalidCell
				}
			case '*':
				if r == 0 || r == l-1 {
					if c == 0 || c == w-1 {
						return ErrInvalidCorner
					}
					return ErrInvalidHorizontalEdge
				}
				if c == 0 || c == w-1 {
					return ErrInvalidVerticalEdge
				}
				b.increment(r-1, c-1)
				b.increment(r-1, c)
				b.increment(r-1, c+1)
				b.increment(r, c-1)
				b.increment(r, c+1)
				b.increment(r+1, c-1)
				b.increment(r+1, c)
				b.increment(r+1, c+1)
			case ' ', '1', '2', '3', '4', '5', '6', '7', '8':
			default:
				return ErrInvalidCell
			}
		}
	}

	return nil
}

func (b Board) increment(r, c int) {
	switch b[r][c] {
	case ' ':
		b[r][c] = '1'
	case '1', '2', '3', '4', '5', '6', '7':
		b[r][c]++
	}

	return
}
