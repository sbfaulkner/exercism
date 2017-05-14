package pascal

const testVersion = 1

// Triangle calculates Pascal's triangle given a size
func Triangle(size int) [][]int {
	if size == 0 {
		return [][]int{}
	}

	if size == 1 {
		return [][]int{{1}}
	}

	triangle := append(Triangle(size-1), make([]int, size, size))

	for i := 0; i < size; i++ {
		if i > 0 {
			triangle[size-1][i] += triangle[size-2][i-1]
		}
		if i < size-1 {
			triangle[size-1][i] += triangle[size-2][i]
		}
	}

	return triangle
}
