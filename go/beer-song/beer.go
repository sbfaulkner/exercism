package beer

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

const testVersion = 1

// ErrBeerBadVerse indicates an invalid verse number.
var ErrBeerBadVerse = errors.New("beer: invalid verse specified")

// Verse returns the lyrics for a single verse.
func Verse(n int) (verse string, err error) {
	if n < 0 || n > 99 {
		return "", ErrBeerBadVerse
	}

	verse = fmt.Sprintf(
		"%s, %s.\n%s, %s.\n",
		capitalize(bottlesOfBeerOnTheWall(n)),
		bottlesOfBeer(n),
		takeOneDownAndPassItAround(n),
		remainingBottlesOfBeerOnTheWall(n),
	)

	return
}

// Verses returns the lyrics for a series of verses.
func Verses(from int, to int) (string, error) {
	if to > from {
		return "", ErrBeerBadVerse
	}

	verses := []string{}

	for n := from; n >= to; n-- {
		v, err := Verse(n)

		if err != nil {
			return "", err
		}

		verses = append(verses, v)
	}

	return strings.Join(verses, "\n") + "\n", nil
}

// Song returns the lyrics for the entire song.
func Song() string {
	verses, _ := Verses(99, 0)
	return verses
}

// bottles returns the singular or plural of "bottle" based on the provided number.
func bottles(n int) string {
	if n == 1 {
		return "bottle"
	}

	return "bottles"
}

// bottles returns the string "N bottle(s) of beer" with the count and pluralization.
func bottlesOfBeer(n int) string {
	if n == 0 {
		return fmt.Sprintf("no more %s of beer", bottles(n))
	}

	return fmt.Sprintf("%d %s of beer", n, bottles(n))
}

// bottlesOfBeerOnTheWall returns the string "N bottle(s) of beer on the wall" with the count and pluralization.
func bottlesOfBeerOnTheWall(n int) string {
	return fmt.Sprintf("%s on the wall", bottlesOfBeer(n))
}

// one returns "one" or "it" to use as a the object of a sentence.
func one(n int) string {
	if n == 1 {
		return "it"
	}

	return "one"
}

// takeOneDownAndPassItAround uses "one" or "it" appropriately in the phrase.
func takeOneDownAndPassItAround(n int) string {
	if n == 0 {
		return "Go to the store and buy some more"
	}

	return fmt.Sprintf("Take %s down and pass it around", one(n))
}

// remainingBottlesOfBeerOnTheWall makes sure we wrap around to 99 if we have no beer.
func remainingBottlesOfBeerOnTheWall(n int) string {
	if n == 0 {
		return bottlesOfBeerOnTheWall(99)
	}

	return bottlesOfBeerOnTheWall(n - 1)
}

// capitalize is used to change the first character of a phrase to upper case.
func capitalize(s string) string {
	return string(unicode.ToUpper(rune(s[0]))) + s[1:]
}
