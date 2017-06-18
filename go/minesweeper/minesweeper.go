package minesweeper

import (
	"errors"
	"regexp"
	"strings"
)

const testVersion = 1

// Board errors
var (
	ErrInvalidWidth = errors.New("minesweeper: invalid board width")
	ErrInvalidEdge  = errors.New("minesweeper: invalid horizontal edge")
	ErrInvalidRow   = errors.New("minesweeper: invalid vertical edge")
)

var (
	validEdge = regexp.MustCompile(`^\+-*\+$`)
	validRow  = regexp.MustCompile(`^\|[ *]*\|$`)
)

// Count marks the mine count for all spaces on the board
func (b *Board) Count() error {
	l := len(*b)
	w := len((*b)[0])

	for r := 0; r < l; r++ {
		if err := b.validateRow(r, w, l); err != nil {
			return err
		}

		for c := 1; c < len((*b)[r])-1; c++ {
			if (*b)[r][c] != ' ' {
				continue
			}

			if m := b.count(r, c); m > 0 {
				(*b)[r][c] = '0' + byte(m)
			}
		}
	}

	return nil
}

func (b Board) count(r, c int) (n int) {
	n += strings.Count(string(b[r-1][c-1:c+2]), "*")
	n += strings.Count(string(b[r][c-1:c+2]), "*")
	n += strings.Count(string(b[r+1][c-1:c+2]), "*")

	return
}

func (b Board) validateRow(r, w, l int) error {
	row := b[r]

	if len(row) != w {
		return ErrInvalidWidth
	}

	l--

	if r%l == 0 {
		if !validEdge.Match(row) {
			return ErrInvalidEdge
		}
	} else {
		if !validRow.Match(row) {
			return ErrInvalidRow
		}
	}

	return nil
}
