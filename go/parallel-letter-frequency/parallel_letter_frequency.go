package letter

const testVersion = 1

// ConcurrentFrequency executes the Frequency function in concurrent threads and merges the results
func ConcurrentFrequency(inputs []string) FreqMap {
	c := make(chan FreqMap, len(inputs))

	for _, input := range inputs {
		go func(in string) { c <- Frequency(in) }(input)
	}

	result := make(FreqMap)

	for _ = range inputs {
		for r, count := range <-c {
			result[r] += count
		}
	}

	return result
}
