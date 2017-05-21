package say

import (
	"fmt"
	"os/exec"
	"strings"
)

const testVersion = 1

var numbers = [][]string{
	{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"},
	{"ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"},
	{"twenty"},
	{"thirty"},
	{"forty"},
	{"fifty"},
	{"sixty"},
	{"seventy"},
	{"eighty"},
	{"ninety"},
}

// var powers = []string{"one", "ten", "hundred", "thousand", "million", "billion"}

// Say converts a number to text.
func Say(number uint64) string {
	text := spell(number)

	exec.Command("say", text).Run()

	return text
}

func spell(number uint64) string {
	var values []string

	hundreds := number / 100
	tens := number % 100 / 10
	ones := number % 10

	if hundreds > 0 {
		values = append(values, numbers[0][hundreds], "hundred")
	}

	if ones < uint64(len(numbers[tens])) {
		if ones > 0 || tens > 0 || number == 0 {
			values = append(values, numbers[tens][ones])
		}
	} else if tens > 0 {
		values = append(values, fmt.Sprintf("%s-%s", numbers[tens][0], numbers[0][ones]))
	} else {
		values = append(values, numbers[0][ones])
	}

	return strings.Join(values, " ")
}
