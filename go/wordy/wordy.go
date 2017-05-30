package wordy

import (
	"regexp"
	"strconv"
	"strings"
)

const testVersion = 1

// Answer parses and evaluates a simple math word problem.
func Answer(input string) (answer int, ok bool) {
	state := 0
	fields, parsed := problemFields(input)
	if !parsed {
		return 0, false
	}

	for _, token := range fields {
		switch state {
		case 0: // start state
			left, err := strconv.Atoi(token)
			if err != nil {
				break
			}
			answer = left
			state = 1
		case 1: // have left value
			switch token {
			case "plus":
				state = 2
			case "minus":
				state = 3
			case "multiplied":
				state = 4
			case "divided":
				state = 6
			default: // unknown operator
				return 0, false
			}
		case 2: // plus
			right, err := strconv.Atoi(token)
			if err != nil {
				break
			}
			answer += right
			state = 1
		case 3: // minus
			right, err := strconv.Atoi(token)
			if err != nil {
				break
			}
			answer -= right
			state = 1
		case 4: // multiplied
			if token != "by" {
				break
			}
			state = 5
		case 5: // multiplied by
			right, err := strconv.Atoi(token)
			if err != nil {
				break
			}
			answer *= right
			state = 1
		case 6: // divided
			if token != "by" {
				break
			}
			state = 7
		case 7: // divided by
			right, err := strconv.Atoi(token)
			if err != nil {
				break
			}
			answer /= right
			state = 1
		default:
			break
		}
	}

	if state != 1 {
		return
	}

	return answer, true
}

var problemRegexp = regexp.MustCompile(`\AWhat is (.*)\?`)

func problemFields(input string) ([]string, bool) {
	m := problemRegexp.FindStringSubmatch(input)
	if len(m) < 2 {
		return []string{}, false
	}
	return strings.Fields(m[1]), true
}
