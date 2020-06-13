package etl

import (
	"strings"
)

// Transform legacy scrabble scores to the new system
func Transform(input map[int][]string) map[string]int {
	output := map[string]int{}

	for score, tiles := range input {
		for _, tile := range tiles {
			output[strings.ToLower(tile)] = score
		}
	}

	return output
}
