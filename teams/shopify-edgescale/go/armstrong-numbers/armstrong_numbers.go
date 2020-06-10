package armstrong

func digits(num int) int {
	d := 0

	for n := num; n > 0; n /= 10 {
		d++
	}

	return d
}

func pow(num int, power int) int {
	if power == 0 {
		return 1
	}
	return num * pow(num, power-1)
}

// IsNumber determines whether a number is an Armstrong number
func IsNumber(num int) bool {
	l := digits(num)

	s := 0

	for n := num; n > 0; n /= 10 {
		s += pow(n%10, l)
	}

	return s == num
}
