package bob

import (
	"strings"
	"unicode"
)

const testVersion = 3

func isQuestion(greeting string) bool {
	lastCharacter := greeting[len(greeting)-1]
	return lastCharacter == '?'
}

func isYell(greeting string) bool {
	for _, ch := range greeting {
		if unicode.IsLower(ch) {
			return false
		}
	}

	return isText(greeting)
}

func isText(greeting string) bool {
	for _, ch := range greeting {
		if unicode.IsLetter(ch) {
			return true
		}
	}

	return false
}

// Hey simulates teenager Bob's responses when you "talk" to him
func Hey(greeting string) (response string) {
	greeting = strings.TrimSpace(greeting)

	if greeting == "" {
		response = "Fine. Be that way!"
	} else if isYell(greeting) {
		response = "Whoa, chill out!"
	} else if isQuestion(greeting) {
		response = "Sure."
	} else {
		response = "Whatever."
	}

	return
}
