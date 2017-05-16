package foodchain

import (
	"fmt"
	"strings"
)

const testVersion = 3

var verseData = []struct {
	animal      string
	description string
}{
	{"fly", "I don't know why she swallowed the fly. Perhaps she'll die."},
	{"spider", "It wriggled and jiggled and tickled inside her."},
	{"bird", "How absurd to swallow a bird!"},
	{"cat", "Imagine that, to swallow a cat!"},
	{"dog", "What a hog, to swallow a dog!"},
	{"goat", "Just opened her throat and swallowed a goat!"},
	{"cow", "I don't know how she swallowed a cow!"},
	{"horse", "She's dead, of course!"},
}

// Song returns the entire song.
func Song() string {
	return Verses(1, len(verseData))
}

// Verses returns `count` verses beginning with the verse specified by `number`.
func Verses(number int, count int) string {
	verses := make([]string, count)

	for i := range verses {
		verses[i] = Verse(number + i)
	}

	return strings.Join(verses, "\n\n")
}

// Verse returns the verse specified by `number`.
func Verse(number int) string {
	stanzas := []string{}

	for i, n := 0, number-1; i < number; i, n = i+1, n-1 {
		if i == 0 {
			stanzas = append(
				stanzas,
				fmt.Sprintf("I know an old lady who swallowed a %s.", verseData[n].animal),
				verseData[n].description,
			)

			if number == len(verseData) {
				break
			}
		} else if n == 1 {
			stanza := fmt.Sprintf(
				"She swallowed the %s to catch the %s that %s",
				verseData[n+1].animal,
				verseData[n].animal,
				verseData[n].description[3:],
			)
			stanzas = append(stanzas, stanza)
		} else {
			stanzas = append(
				stanzas,
				fmt.Sprintf("She swallowed the %s to catch the %s.", verseData[n+1].animal, verseData[n].animal),
			)

			if n == 0 {
				stanzas = append(stanzas, verseData[n].description)
			}
		}
	}

	return strings.Join(stanzas, "\n")
}
