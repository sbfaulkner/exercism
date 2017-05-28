package transpose

import (
	"strings"
)

const testVersion = 1

// Transpose transposes a matrix of input text.
func Transpose(input []string) []string {
	if len(input) == 0 {
		return []string{}
	}

	// determine the length of the longest line to allocate the rows required.
	width := 0

	for _, row := range input {
		length := len(row)
		if length > width {
			width = length
		}
	}

	output := make([]string, width)

	// transpose the character matrix, padding short columns with nul characters
	for _, row := range input {
		for c := 0; c < width; c++ {
			if c < len(row) {
				output[c] += row[c : c+1]
			} else {
				output[c] += "\x00"
			}
		}
	}

	// remove any right padding, and change any remaining padding to spaces
	for r, row := range output {
		output[r] = strings.Replace(strings.TrimRight(row, "\x00"), "\x00", " ", -1)
	}

	return output
}
