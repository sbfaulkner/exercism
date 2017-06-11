package binarysearch

import (
	"fmt"
)

const testVersion = 1

// SearchInts performs a binary search to determine the position to insert a new member in a slice
func SearchInts(slice []int, key int) int {
	if len(slice) == 0 {
		return 0
	}

	i := len(slice) / 2

	if key <= slice[i] {
		result := SearchInts(slice[0:i], key)

		if key == slice[i] && result == i {
			return i
		}

		return result
	}

	return SearchInts(slice[i+1:], key) + i + 1
}

// Message provides a textual representation of the result of a binary search
func Message(slice []int, key int) string {
	i := SearchInts(slice, key)

	if i == len(slice) {
		if i == 0 {
			return "slice has no values"
		}

		return fmt.Sprintf("%d > all %d values", key, len(slice))
	}

	if slice[i] == key {
		if i == 0 {
			return fmt.Sprintf("%d found at beginning of slice", key)
		}

		if i == len(slice)-1 {
			return fmt.Sprintf("%d found at end of slice", key)
		}

		return fmt.Sprintf("%d found at index %d", key, i)
	}

	if i == 0 {
		return fmt.Sprintf("%d < all values", key)
	}

	return fmt.Sprintf("%d > %d at index %d, < %d at index %d", key, slice[i-1], i-1, slice[i], i)
}
