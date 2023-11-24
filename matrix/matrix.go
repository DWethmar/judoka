package matrix

type Matrix [][]int

// Get returns the value at the given position.
// If the position is out of bounds, the default value is returned.
func (m Matrix) Get(x, y, d int) int {
	if y < 0 || y >= len(m) {
		return d
	}

	return m[y][x]
}

func (m Matrix) Set(x, y int, v int) {
	m[y][x] = v
}

func New(width, height, v int) Matrix {
	m := make([][]int, height)
	for i := range m {
		m[i] = make([]int, width)
		for j := range m[i] {
			m[i][j] = v
		}
	}
	return m
}

func Iterate(m Matrix, fn func(x, y, v int)) {
	for y := range m {
		for x, v := range m[y] {
			fn(x, x, v)
		}
	}
}
