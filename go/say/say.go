package say

import (
	"fmt"
	"os/exec"
	"strings"
)

const testVersion = 1

const zero = "zero"

var ones = []string{
	"", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
}
var teens = []string{
	"", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen",
}
var tens = []string{
	"", "ten", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety",
}
var scales = []string{
	"", "thousand", "million", "billion", "trillion", "quadrillion", "quintillion",
}

// Say converts a number to text.
func Say(number uint64) string {
	var text string
	var chunks []string

	if number == 0 {
		text = zero
	} else {
		for i := 0; number > 0 || i == 0; i++ {
			count := number % 1000
			if count > 0 {
				chunks = append([]string{scales[i]}, chunks...)
				chunks = append([]string{spell(count)}, chunks...)
			}
			number /= 1000
		}

		text = strings.TrimSpace(strings.Join(chunks, " "))
	}

	exec.Command("say", text).Run()

	return text
}

func spell(number uint64) string {
	var values []string

	h := number / 100
	t := number % 100 / 10
	o := number % 10

	if h > 0 {
		values = append(values, ones[h], "hundred")
	}

	if t > 1 {
		if o > 0 {
			values = append(values, fmt.Sprintf("%s-%s", tens[t], ones[o]))
		} else if t > 0 {
			values = append(values, tens[t])
		}
	} else if o > 0 {
		if t == 1 {
			values = append(values, teens[o])
		} else {
			values = append(values, ones[o])
		}
	}

	return strings.Join(values, " ")
}
