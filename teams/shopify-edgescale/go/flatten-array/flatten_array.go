package flatten

// Flatten returns a single flattened list with all values except nil/null
func Flatten(input interface{}) []interface{} {
	ai, ok := input.([]interface{})

	if !ok {
		return []interface{}{input}
	}

	out := make([]interface{}, 0, len(ai))

	for _, i := range ai {
		if i == nil {
			continue
		}

		out = append(out, Flatten(i)...)
	}

	return out
}
