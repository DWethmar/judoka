package terrain

import (
	"github.com/dwethmar/judoka/matrix"
)

// Generator generates a matrix for use in terrain rendering.
type Generator interface {
	Generate(minX, maxX, minY, maxY int) matrix.Matrix
}
