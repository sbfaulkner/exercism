package acronym

import (
	"regexp"
	"strings"
)

const testVersion = 2

// Abbreviate converts an input string into the representative acronym
func Abbreviate(value string) (acronym string) {
	re := regexp.MustCompile("(\\b[A-Za-z])|([a-z][A-Z])")

	matches := re.FindAllString(value, -1)

	for _, match := range matches {
		acronym += strings.ToUpper(match[len(match)-1:])
	}

	return
}
