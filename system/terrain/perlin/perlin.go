package perlin

import (
	"github.com/aquilax/go-perlin"
	"github.com/dwethmar/judoka/matrix"
	"github.com/dwethmar/judoka/system/terrain"
)

var _ terrain.Generator = &Generator{}

type Generator struct {
	ChunkSize int

	// Perlin noise settings
	alpha float64
	beta  float64
	n     int32
	seed  int64
}

func New() *Generator {
	return &Generator{
		ChunkSize: terrain.ChunkSize,
		alpha:     2.,
		beta:      2.,
		n:         3,
		seed:      100,
	}
}

// Generate implements terrain.Generator.
func (g *Generator) Generate(minX int, maxX int, minY int, maxY int) matrix.Matrix {
	w := maxX - minX
	h := maxY - minY

	scale := 20.0 // Adjust this scale factor as needed
	p := perlin.NewPerlin(g.alpha, g.beta, g.n, g.seed)
	m := matrix.New(w, h, 0)

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			noiseVal := p.Noise2D(float64(minX+x)/scale, float64(minY+y)/scale)
			normalizedVal := (noiseVal + 1) / 2
			finalVal := normalizedVal * 5
			m.Set(x, y, int(finalVal))
		}
	}

	return m
}
