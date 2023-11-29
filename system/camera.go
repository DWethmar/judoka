package system

import (
	"image"

	"github.com/dwethmar/judoka/entity"
)

// Camera is a shared struct that is used to follow an entity.
type Camera struct {
	Viewport  entity.Entity
	Following entity.Entity
	Bounds    image.Rectangle
}
