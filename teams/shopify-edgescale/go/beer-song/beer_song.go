package beer

import (
	"errors"
	"fmt"
	"strings"
)

func bottles(n int) string {
	switch n {
	case 0:
		return "No more bottles of beer"

	case 1:
		return "1 bottle of beer"
	}

	return fmt.Sprintf("%d bottles of beer", n)
}

func it(n int) string {
	if n == 1 {
		return "it"
	}

	return "one"
}

func action(n int) string {
	if n == 0 {
		return "Go to the store and buy some more"
	}

	return fmt.Sprintf("Take %s down and pass it around", it(n))
}

// Verse returns the lyrics for a single verse of 99 bottles of beer
func Verse(n int) (string, error) {
	if n < 0 || n > 99 {
		return "", errors.New("number of bottles must be in range 0 <= n <= 99")
	}

	b1 := bottles(n)
	b2 := bottles((n + 99) % 100)

	v := fmt.Sprintf(
		"%s on the wall, %s.\n%s, %s on the wall.\n",
		b1,
		strings.ToLower(b1),
		action(n),
		strings.ToLower(b2),
	)

	return v, nil
}

// Verses returns the lyrics for a range of verses of 99 bottles of beer
func Verses(start, end int) (string, error) {
	if start < end {
		return "", errors.New("end must be less than start")
	}

	vs := make([]string, start-end+2)

	for i := start; i >= end; i-- {
		v, err := Verse(i)
		if err != nil {
			return "", err
		}

		vs[start-i] = v
	}

	return strings.Join(vs, "\n"), nil
}

// Song returns the lyrics for all verses of 99 bottles of beer
func Song() string {
	vs, _ := Verses(99, 0)

	return vs
}
