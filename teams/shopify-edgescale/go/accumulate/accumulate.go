package accumulate

// Accumulate returns the result of applying an operation to each element of
// the input collection
func Accumulate(in []string, op func(string) string) []string {
	out := make([]string, 0, len(in))

	for _, i := range in {
		out = append(out, op(i))
	}

	return out
}
