package anagram

import (
	"sort"
	"strings"
)

const testVersion = 1

// Detect selects the correct sublist, given a word and a list of possible anagrams.
func Detect(word string, list []string) []string {
	anagrams := []string{}

	word = strings.ToLower(word)

	for _, w := range list {
		w = strings.ToLower(w)
		if word != w && sortRunes(word) == sortRunes(w) {
			anagrams = append(anagrams, w)
		}
	}

	return anagrams
}

type byRune []rune

func (r byRune) Len() int           { return len(r) }
func (r byRune) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r byRune) Less(i, j int) bool { return r[i] < r[j] }

func sortRunes(s string) string {
	r := []rune(s)
	sort.Sort(byRune(r))
	return string(r)
}
