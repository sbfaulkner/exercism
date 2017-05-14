package secret

const testVersion = 1

var operations = []string{
	"wink",
	"double blink",
	"close your eyes",
	"jump",
}

func reverse(array []string) []string {
	rindex := len(array)
	reversed := make([]string, rindex)

	for _, value := range array {
		rindex--
		reversed[rindex] = value
	}

	return reversed
}

// Handshake returns the secret handshake for the provided input value
func Handshake(input uint) (handshake []string) {
	var i uint

	for i = 0; i < 4; i++ {
		if input&(1<<i) != 0 {
			handshake = append(handshake, operations[i])
		}
	}

	if input&(1<<4) != 0 {
		handshake = reverse(handshake)
	}

	return
}
