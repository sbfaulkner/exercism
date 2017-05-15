package series

const testVersion = 2

// All returns a list of all substrings of s with length n.
func All(n int, s string) (all []string) {
	all = make([]string, len(s)-n+1)

	for i := 0; i < len(all); i++ {
		all[i] = s[i : i+n]
	}

	return all
}

// UnsafeFirst returns the first substring of s with length n.
func UnsafeFirst(n int, s string) string {
	if n > len(s) {
		return ""
	}
	return s[:n]
}
