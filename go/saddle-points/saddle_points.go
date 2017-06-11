package matrix

const testVersion = 1

// Pair represents a position in a matrix
type Pair struct{ row, column int }

// Saddle locates any saddle points in a matrix
func (m Matrix) Saddle() []Pair {
	pairs := []Pair{}

	max := []int{}

	for _, row := range m {
		max = append(max, maximum(row))
	}

	min := []int{}

	for _, col := range m.Cols() {
		min = append(min, minimum(col))
	}

	for r, row := range m {
		for c, value := range row {
			if value >= max[r] && value <= min[c] {
				pairs = append(pairs, Pair{r, c})
			}
		}
	}

	return pairs
}

func maximum(input []int) int {
	max := input[0]

	for _, v := range input[1:] {
		if v > max {
			max = v
		}
	}

	return max
}

func minimum(input []int) int {
	min := input[0]

	for _, v := range input[1:] {
		if v < min {
			min = v
		}
	}

	return min
}
