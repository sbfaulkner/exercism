package etl

import (
	"unicode"
)

// Transform legacy scrabble scores to the new system
func Transform(input map[int][]string) map[string]int {
	output := make(map[string]int, 26)

	for score, tiles := range input {
		for _, tile := range tiles {
			output[string(unicode.ToLower([]rune(tile)[0]))] = score
		}
	}

	return output
}
