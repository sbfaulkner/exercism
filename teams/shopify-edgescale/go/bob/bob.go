// Package bob simulates a lackadaisical teenager
package bob

import (
	"regexp"
)

var reQuestion = regexp.MustCompile(`\?\s*$`)
var reYelling = regexp.MustCompile(`^[^a-z]+[.!?]?\s*$`)
var reLetters = regexp.MustCompile(`(?i)[A-Z]`)
var reSilence = regexp.MustCompile(`^\s*$`)

func isQuestion(remark string) bool {
	return reQuestion.MatchString(remark)
}

func isYelling(remark string) bool {
	return reYelling.MatchString(remark)
}

func hasLetters(remark string) bool {
	return reLetters.MatchString(remark)
}

func isSilence(remark string) bool {
	return reSilence.MatchString(remark)
}

// Hey responds like a lackadaisical teenager
func Hey(remark string) string {
	if isSilence(remark) {
		return "Fine. Be that way!"
	}

	if hasLetters(remark) {
		if isYelling(remark) {
			if isQuestion(remark) {
				return "Calm down, I know what I'm doing!"
			}

			return "Whoa, chill out!"
		}
	}

	if isQuestion(remark) {
		return "Sure."
	}

	return "Whatever."
}
