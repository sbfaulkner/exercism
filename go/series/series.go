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
	return s[:n]
}

// First returns the first substring of s with length n and a boolean indicating whether it is successful.
func First(n int, s string) (first string, ok bool) {
	if n <= len(s) {
		first = UnsafeFirst(n, s)
		ok = true
	}

	return
}
