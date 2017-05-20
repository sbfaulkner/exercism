package brackets

import (
	"strings"
)

const testVersion = 5

var matchingBrackets = map[rune]rune{
	'[': ']',
	'{': '}',
	'(': ')',
}

type bracketStack []rune

func (stack *bracketStack) push(b rune) *bracketStack {
	*stack = append(*stack, b)
	return stack
}

func (stack *bracketStack) pop() rune {
	b := (*stack)[len(*stack)-1]
	*stack = (*stack)[:len(*stack)-1]
	return b
}

// Bracket makes sure the brackets and braces all match.
func Bracket(input string) (bool, error) {
	var brackets bracketStack

	for _, r := range input {
		if strings.ContainsRune("[{(", r) {
			brackets.push(r)
		} else if strings.ContainsRune("]})", r) {
			if len(brackets) == 0 {
				return false, nil
			}
			if b := brackets.pop(); matchingBrackets[b] != r {
				return false, nil
			}
		}
	}

	return len(brackets) == 0, nil
}
