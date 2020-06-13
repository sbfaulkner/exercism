package flatten

func flatEach(input interface{}, f func(interface{})) {
	if input == nil {
		return
	}

	ai, ok := input.([]interface{})
	if ok {
		for _, i := range ai {
			flatEach(i, f)
		}
	} else {
		f(input)
	}
}

// Flatten returns a single flattened list with all values except nil/null
func Flatten(input interface{}) []interface{} {
	count := 0
	flatEach(input, func(interface{}) { count++ })

	out := make([]interface{}, 0, count)
	flatEach(input, func(i interface{}) { out = append(out, i) })

	return out
}
