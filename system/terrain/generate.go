package terrain

import (
	"image"

	"github.com/aquilax/go-perlin"
	"github.com/dwethmar/judoka/matrix"
)

const (
	alpha       = 2.
	beta        = 2.
	n           = 3
	seed  int64 = 100
)

func Generate(r image.Rectangle) matrix.Matrix {
	scale := 20.0 // Adjust this scale factor as needed
	p := perlin.NewPerlin(alpha, beta, n, seed)
	m := matrix.New(r.Dx(), r.Dy(), 0)

	for x := 0; x < r.Dx(); x++ {
		for y := 0; y < r.Dy(); y++ {
			noiseVal := p.Noise2D(float64(r.Min.X+x)/scale, float64(r.Min.Y+y)/scale)
			normalizedVal := (noiseVal + 1) / 2
			finalVal := normalizedVal * 5
			m.Set(x, y, int(finalVal))
		}
	}

	return m
}

func TestChunk2() matrix.Matrix {
	m := matrix.New(ChunkSize, ChunkSize, 0) // 0 = dirt

	m.Set(1, 0, 1)
	m.Set(2, 0, 1)
	m.Set(1, 1, 1)
	m.Set(2, 1, 1)

	m.Set(0, 2, 1)

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

	// 4x4
	var x, y int = 6, 8
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			m.Set(x+i, y+j, 1)
		}
	}

	return m
}

func TestChunk() matrix.Matrix {
	m := matrix.New(ChunkSize, ChunkSize, 0)
	for x := 0; x < ChunkSize; x++ {
		for y := 0; y < ChunkSize; y++ {
			m.Set(x, y, 0)
		}
	}

	// border dirt
	for x := 0; x < ChunkSize; x++ {
		m.Set(x, 0, 1)
		m.Set(x, ChunkSize-1, 1)
	}

	for y := 0; y < ChunkSize; y++ {
		m.Set(0, y, 1)
		m.Set(ChunkSize-1, y, 1)
	}

	// grass 2x2
	for x := 3; x < 5; x++ {
		for y := 3; y < 5; y++ {
			m.Set(x, y, 0)
		}
	}

	return m
}
