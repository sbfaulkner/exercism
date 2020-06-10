package armstrong

func digits(num int) []int {
	d := []int{}

	for num > 0 {
		d = append([]int{num % 10}, d...)
		num /= 10
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
	d := digits(num)
	l := len(d)

	s := 0

	for _, dd := range d {
		s += pow(dd, l)
	}

	return s == num
}
