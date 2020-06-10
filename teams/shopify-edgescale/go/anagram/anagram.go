package anagram

import (
	"sort"
	"strings"
)

func sorted(s string) string {
	ss := []rune(strings.ToLower(s))

	less := func(i, j int) bool {
		if ss[i] < ss[j] {
			return true
		}

		return false
	}

	sort.Slice(ss, less)

	return string(ss)
}

// Detect select the sublist of anagrams of the given word
func Detect(subject string, candidates []string) []string {
	anagrams := make([]string, 0, len(candidates))

	s := sorted(subject)

	for _, c := range candidates {
		if strings.ToLower(subject) == strings.ToLower(c) {
			continue
		}

		if s == sorted(c) {
			anagrams = append(anagrams, c)
		}
	}

	return anagrams
}
