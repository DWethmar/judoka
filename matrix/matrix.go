package matrix

type Matrix [][]int32

func (m Matrix) Get(x, y int32) int32 {
	return m[y][x]
}

func (m Matrix) Set(x, y int32, v int32) {
	m[y][x] = v
}

func New(width, height int, v int32) Matrix {
	m := make([][]int32, height)
	for i := range m {
		m[i] = make([]int32, width)
		for j := range m[i] {
			m[i][j] = v
		}
	}
	return m
}

func Iterate(m Matrix, fn func(x, y, v int32)) {
	for y := range m {
		for x, v := range m[y] {
			fn(int32(x), int32(y), v)
		}
	}
}
