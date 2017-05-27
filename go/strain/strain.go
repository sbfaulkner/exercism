package strain

const testVersion = 1

// Ints defines a collection of int values
type Ints []int

// Lists defines a collection of arrays of ints
type Lists [][]int

// Strings defines a collection of strings
type Strings []string

// Keep filters a collection of ints to only contain the members where the provided function returns true.
func (i Ints) Keep(strainer func(int) bool) Ints {
	o := Ints{}

	for _, v := range i {
		if strainer(v) {
			o = append(o, v)
		}
	}

	if len(o) == 0 {
		return nil
	}

	return o
}

// Discard filters a collection to only contain the members where the provided function returns false.
func (i Ints) Discard(strainer func(int) bool) Ints {
	return i.Keep(func(n int) bool { return !strainer(n) })
}

// Keep filters a collection of lists to only contain the members where the provided function returns true.
func (l Lists) Keep(strainer func([]int) bool) Lists {
	o := Lists{}

	for _, v := range l {
		if strainer(v) {
			o = append(o, v)
		}
	}

	if len(o) == 0 {
		return nil
	}

	return o
}

// Keep filters a collection of strings to only contain the members where the provided function returns true.
func (s Strings) Keep(strainer func(string) bool) Strings {
	o := Strings{}

	for _, v := range s {
		if strainer(v) {
			o = append(o, v)
		}
	}

	if len(o) == 0 {
		return nil
	}

	return o
}
