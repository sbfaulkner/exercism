package ocr

import (
	"fmt"
	"strings"
)

const testVersion = 1

// Recognize returns multiple sets of digits within the provided input lines.
func Recognize(input string) (output []string) {
	lines := strings.Split(input, "\n")[1:]

	for l := 0; l < len(lines); l += 4 {
		output = append(output, recognizeDigits(lines[l:l+3]))
	}

	return
}

// recognizeDigits returns the set of digits within the provided input lines.
func recognizeDigits(lines []string) (digits string) {
	length := maxLength(lines)

	for i := 0; i < length; i += 3 {
		digits += recognizeDigit(lines, i)
	}

	return
}

// maxLength returns the length of the longest line within the provided lines.
func maxLength(lines []string) (length int) {
	for _, line := range lines {
		if len(line) > length {
			length = len(line)
		}
	}

	return
}

// recognitionState is used to maintain internal state when recognizing digits.
type recognitionState int

// states for recognizing digits
const (
	digitUnknown recognitionState = 0x00
	digitStart   recognitionState = 0x01

	digit02356789 recognitionState = 0x10
	digit14       recognitionState = 0x11

	digit0  recognitionState = 0x20
	digit1  recognitionState = 0x21
	digit23 recognitionState = 0x22
	digit4  recognitionState = 0x24
	digit56 recognitionState = 0x25
	digit7  recognitionState = 0x27
	digit89 recognitionState = 0x29

	digitFinal0 recognitionState = 0x30
	digitFinal1 recognitionState = 0x31
	digitFinal2 recognitionState = 0x32
	digitFinal3 recognitionState = 0x33
	digitFinal4 recognitionState = 0x34
	digitFinal5 recognitionState = 0x35
	digitFinal6 recognitionState = 0x36
	digitFinal7 recognitionState = 0x37
	digitFinal8 recognitionState = 0x38
	digitFinal9 recognitionState = 0x39
)

// state transition table for digit recognition
var digitPatterns = map[recognitionState]map[string]recognitionState{
	digitStart: {
		" _ ": digit02356789,
		"   ": digit14,
	},
	digit02356789: {
		"| |": digit0,
		" _|": digit23,
		"|_ ": digit56,
		"  |": digit7,
		"|_|": digit89,
	},
	digit14: {
		"  |": digit1,
		"|_|": digit4,
	},
	digit0: {
		"|_|": digitFinal0,
	},
	digit23: {
		"|_ ": digitFinal2,
		" _|": digitFinal3,
	},
	digit56: {
		" _|": digitFinal5,
		"|_|": digitFinal6,
	},
	digit7: {
		"  |": digitFinal7,
	},
	digit89: {
		"|_|": digitFinal8,
		" _|": digitFinal9,
	},
	digit1: {
		"  |": digitFinal1,
	},
	digit4: {
		"  |": digitFinal4,
	},
}

// recognizeDigit returns the digit at the specified offset within the provided input lines.
func recognizeDigit(input []string, offset int) string {
	state := digitStart

	for _, line := range input {
		l := fmt.Sprintf("%-3s", line[offset:])[0:3]
		state = digitPatterns[state][l]
		if state == digitUnknown {
			return "?"
		}
	}

	return string(state)
}
