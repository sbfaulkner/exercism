package diamond

import (
	"errors"
	"fmt"
	"strings"
)

const testVersion = 1

// Errors defined for diamond.
var (
	ErrNotALetter = errors.New("diamond: not a letter")
)

// Gen will generate a diamond starting with 'A' with the supplied letter at the widest point.
func Gen(char byte) (string, error) {
	if char < 'A' || char > 'Z' {
		return "", ErrNotALetter
	}

	var rows = []string{row('A', char)}

	for c := byte('B'); c <= char; c++ {
		rows = append(rows, row(c, char))
	}

	for i := len(rows) - 2; i >= 0; i-- {
		rows = append(rows, rows[i])
	}

	return strings.Join(rows, ""), nil
}

func pad(count byte) string {
	return strings.Repeat(" ", int(count))
}

func row(c byte, char byte) string {
	if c == 'A' {
		return fmt.Sprintf("%sA%s\n", pad(char-c), pad(char-c))
	}

	return fmt.Sprintf("%s%c%s%c%s\n", pad(char-c), c, pad(2*(c-'A')-1), c, pad(char-c))
}
