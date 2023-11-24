package debug

import (
	"github.com/dwethmar/judoka/matrix"
	"github.com/dwethmar/judoka/system/terrain"
)

var _ terrain.Generator = &Generator{}

type Generator struct{}

func New() *Generator {
	return &Generator{}
}

// Generate implements terrain.Generator.
func (*Generator) Generate(minX int, maxX int, minY int, maxY int) matrix.Matrix {
	chunkX := minX / terrain.ChunkSize
	chunkY := minY / terrain.ChunkSize
	ChunkSize := terrain.ChunkSize

	if chunkX == 0 && chunkY == 0 {
		m := matrix.New(ChunkSize, ChunkSize, 0) // 0 = dirt

		m.Set(1, 0, 1)
		m.Set(2, 0, 1)
		m.Set(1, 1, 1)
		m.Set(2, 1, 1)

		m.Set(14, 2, 1)

		// line verical
		for y := 0; y < ChunkSize; y++ {
			m.Set(4, y, 1)
		}

		m.Set(2, 5, 1)

		// line horizontal
		for x := 0; x < ChunkSize; x++ {
			m.Set(x, 4, 1)
		}

		m.Set(5, 1, 1)
		m.Set(6, 3, 1)
		m.Set(3, 14, 1)

		// 5x5
		var x, y int = 7, 7
		for i := 0; i < 5; i++ {
			for j := 0; j < 5; j++ {
				m.Set(x+i, y+j, 1)
			}
		}

		// punch a hole
		m.Set(9, 9, 0)

		// 5x5 extrusions
		m.Set(9, 6, 1)  // north
		m.Set(12, 9, 1) // east
		m.Set(9, 12, 1) // south
		m.Set(6, 9, 1)  // west

		// 5x5 corner extrusions
		m.Set(6, 7, 1)
		m.Set(12, 7, 1)
		m.Set(6, 11, 1)
		m.Set(12, 11, 1)

		// 3x3
		x, y = 8, 0
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				m.Set(x+i, y+j, 1)
			}
		}

		return m
	}

	if chunkX == 1 && chunkY == 0 {
		m := matrix.New(ChunkSize, ChunkSize, 0) // 0 = dirt

		// 5x5
		var x, y int = 3, 3
		for i := 0; i < 5; i++ {
			for j := 0; j < 5; j++ {
				m.Set(x+i, y+j, 1)
			}
		}

		// 5x5 extrusions
		m.Set(3, 2, 1)
		m.Set(7, 2, 1)
		m.Set(3, 8, 1)
		m.Set(7, 8, 1)

		// 5x2
		x, y = 0, 12
		for i := 0; i < 10; i++ {
			for j := 0; j < 2; j++ {
				m.Set(x+i, y+j, 1)
			}
		}

		m.Set(2, 12, 0)
		m.Set(5, 13, 0)

		return m
	}

	return matrix.New(ChunkSize, ChunkSize, 0)
}
