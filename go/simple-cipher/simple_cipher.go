package cipher

import (
	"regexp"
	"strings"
	"unicode"
)

const testVersion = 1

type vigenereCipher struct {
	key []rune
}

type shiftCipher struct {
	vigenereCipher
}

type caesarCipher struct {
	shiftCipher
}

// NewCaesar returns a cipher that implements the Caesar cipher
func NewCaesar() Cipher {
	return caesarCipher{shiftCipher{vigenereCipher{[]rune{3}}}}
}

// NewShift returns a cipher thtn implements a shift cipher
func NewShift(s int) Cipher {
	if s == 0 || s < -25 || s > 25 {
		return nil
	}

	return shiftCipher{vigenereCipher{[]rune{rune(s)}}}
}

// NewVigenere returns a cipher that implements the Vigenère cipher
func NewVigenere(k string) Cipher {
	if len(k) == 0 {
		return nil
	}

	if k == strings.Repeat("a", len(k)) {
		return nil
	}

	for _, r := range k {
		if !unicode.IsLower(r) {
			return nil
		}
	}

	key := make([]rune, len(k))

	for i, b := range k {
		key[i] = b - 'a'
	}

	return vigenereCipher{key}
}

func (c vigenereCipher) Encode(input string) string {
	return c.shift(input, 1)
}

func (c vigenereCipher) Decode(input string) string {
	return c.shift(input, -1)
}

func (c vigenereCipher) shift(input string, d int) string {
	input = normalize(input)

	output := make([]rune, len(input))

	for i, r := range input {
		s := c.key[i%len(c.key)] * rune(d)
		output[i] = (r-'a'+s+26)%26 + 'a'
	}

	return string(output)
}

var regexpAlpha = regexp.MustCompile(`[^a-z]`)

func normalize(input string) (output string) {
	for _, i := range input {
		if unicode.IsLetter(i) {
			output += string(unicode.ToLower(i))
		}
	}
	return
}
