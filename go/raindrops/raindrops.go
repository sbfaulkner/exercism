package raindrops

import (
	"fmt"
)

const testVersion = 3

var factors = []int{3, 5, 7}
var conversions = map[int]string{
	3: "Pling",
	5: "Plang",
	7: "Plong",
}

// Convert takes and input number, and returns a string of the number converted to raindrop-speak
func Convert(number int) (output string) {
	for _, factor := range factors {
		if number%factor == 0 {
			output += conversions[factor]
		}
	}

	if output == "" {
		output = fmt.Sprint(number)
	}
	return
}
