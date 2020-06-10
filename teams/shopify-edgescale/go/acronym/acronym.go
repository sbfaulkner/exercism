// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"regexp"
	"unicode"
	"unicode/utf8"
)

var delim = regexp.MustCompile("[ _-]+")

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	a := ""

	for _, w := range delim.Split(s, -1) {
		r, _ := utf8.DecodeRuneInString(w)
		a += string(unicode.ToUpper(r))
	}

	return a
}
