package say

import "fmt"

import "os/exec"

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
	var text string

	tens := number / 10
	ones := number % 10

	if ones < uint64(len(numbers[tens])) {
		text = numbers[tens][ones]
	} else if tens > 0 {
		text = fmt.Sprintf("%s-%s", numbers[tens][0], numbers[0][ones])
	} else {
		text = numbers[0][ones]
	}

	exec.Command("say", text).Run()

	return text
}
