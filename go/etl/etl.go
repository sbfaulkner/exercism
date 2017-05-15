package etl

import (
	"strings"
)

const testVersion = 1

// Transform converts scrabble data from letters per score, to score per letter
func Transform(input map[int][]string) map[string]int {
	result := make(map[string]int)

	for points, letters := range input {
		for _, letter := range letters {
			result[strings.ToLower(letter)] = points
		}
	}

	return result
}
