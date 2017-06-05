package matrix

import (
	"errors"
	"strconv"
	"strings"
)

const testVersion = 1

// Matrix of integers
type Matrix [][]int

// New creates a matrix from the newline delimited rows of space delimited integers
func New(input string) (m Matrix, err error) {
	rows := strings.Split(input, "\n")
	m = make([][]int, len(rows))

	for r, row := range rows {
		cols := strings.Fields(row)

		if r > 0 && len(cols) != len(m[r-1]) {
			err = errors.New("wrong number of columns")
			return
		}

		m[r] = make([]int, len(cols))

		for c, col := range cols {
			m[r][c], err = strconv.Atoi(col)
			if err != nil {
				return
			}
		}
	}

	return
}

// Rows return the rows of integers from the matrix
func (m Matrix) Rows() [][]int {
	output := make([][]int, len(m))

	for r, row := range m {
		output[r] = make([]int, len(row))
		copy(output[r], row)
	}

	return output
}

// Cols return the columns of integers from the matrix
func (m Matrix) Cols() [][]int {
	output := make([][]int, len(m[0]))

	for r := 0; r < len(m[0]); r++ {
		output[r] = make([]int, len(m))

		for c := 0; c < len(m); c++ {
			output[r][c] = m[c][r]
		}
	}

	return output
}

// Set updates an integer value within a matrix
func (m *Matrix) Set(row, col, val int) bool {
	if row < 0 || row >= len(*m) {
		return false
	}

	if col < 0 || col >= len((*m)[0]) {
		return false
	}

	(*m)[row][col] = val
	return true
}
