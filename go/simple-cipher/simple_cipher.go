// For Step 1, implement the Caesar cipher, which seems clear enough except
// maybe for figuring out whether to add or subtract.
//
// For Step 2, implement a shift cipher, like Caesar except the shift amount
// is specified as an int.  (Each letter of "ddddddddddddddddd", is 'd';
// 'd'-'a' == 3, so the corresponding shift amount is 3.)
//
// Steps 2 and 3 seem to be describing the Vigenère cipher (Google it)
// so let's do that too.  The random thing, don't worry about.  There is
// no test for that.
//
// API:
//
// NewCaesar() Cipher
// NewShift(int) Cipher
// NewVigenere(string) Cipher
//
// Interface Cipher is in file cipher.go.
//
// Argument for NewShift must be in the range 1 to 25 or -1 to -25.
// Zero is disallowed.  For invalid arguments NewShift returns nil.
//
// Argument for NewVigenere must consist of lower case letters a-z only.
// Values consisting entirely of the letter 'a' are disallowed.
// For invalid arguments NewVigenere returns nil.

package cipher

import (
	"regexp"
	"unicode"
)

const testVersion = 1

type vigenereCipher []rune

type shiftCipher struct {
	shift int
}

type caesarCipher struct {
	shiftCipher
}

// NewCaesar returns a cipher that implements the Caesar cipher
func NewCaesar() Cipher {
	return caesarCipher{shiftCipher{3}}
}

// NewShift returns a cipher thtn implements a shift cipher
func NewShift(s int) Cipher {
	if s == 0 || s < -25 || s > 25 {
		return nil
	}

	return shiftCipher{s}
}

// NewVigenere returns a cipher that implements the Vigenère cipher
func NewVigenere(k string) Cipher {
	key := make([]rune, len(k))

	for i, b := range k {
		key[i] = b - 'a'
	}

	return vigenereCipher(key)
}

func (c shiftCipher) Encode(input string) string {
	return shift(input, c.shift)
}

func (c shiftCipher) Decode(input string) string {
	return shift(input, -c.shift)
}

func (c vigenereCipher) Encode(input string) string {
	return input
}

func (c vigenereCipher) Decode(input string) string {
	return input
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

func shift(input string, s int) string {
	input = normalize(input)

	output := make([]rune, len(input))

	for i, r := range input {
		output[i] = (r-'a'+rune(s)+26)%26 + 'a'
	}

	return string(output)
}
