package accumulate

const testVersion = 1

// Accumulate maps an array of values using the provided function
func Accumulate(values []string, convert func(string) string) []string {
	mappedValues := make([]string, len(values))

	for index, value := range values {
		mappedValues[index] = convert(value)
	}

	return mappedValues
}
