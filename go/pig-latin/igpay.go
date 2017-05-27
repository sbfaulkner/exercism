package igpay

import (
	"fmt"
	"strings"
	"unicode"
)

const testVersion = 1

const vowels = "aeiouy"

// PigLatin converts english text to pig latin.
func PigLatin(phrase string) string {
	in := strings.Fields(phrase)
	out := make([]string, len(in))

	for i, w := range in {
		out[i] = translate(w)
	}

	return strings.Join(out, " ")
}

func translate(w string) string {
	i := strings.IndexFunc(w, isVowel)

	if i > 0 && w[i-1:i+1] == "qu" {
		i++
	} else if i == 0 && w[0:1] == "y" && isVowel(rune(w[1])) {
		i += strings.IndexFunc(w[i+1:], isVowel) + 1
	}

	if w[i:len(w)] == "ay" {
		i = 0
	}

	return fmt.Sprintf("%s%say", w[i:len(w)], w[0:i])
}

func isVowel(r rune) bool {
	return strings.ContainsRune(vowels, unicode.ToLower(unicode.SimpleFold(r)))
}
