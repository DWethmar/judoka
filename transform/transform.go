package transform

import (
	"github.com/dwethmar/judoka/entity"
	"github.com/dwethmar/judoka/entity/registry"
)

// Position returns the position of an entity.
// I takes in consideration the positions of parent entities.
func Position(r *registry.Register, e entity.Entity) (x, y int) {
	transform, ok := r.Transform.First(e)
	if !ok {
		return 0, 0
	}

	x += transform.X
	y += transform.Y

	if p, ok := r.Parent(e); ok {
		x2, y2 := Position(r, p)
		x += x2
		y += y2
	}

	return x, y
}
