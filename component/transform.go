package component

import "github.com/dwethmar/judoka/entity"

const TransformType = "transform"

var _ Component = (*Transform)(nil)

type Transform struct {
	CID    uint32
	entity entity.Entity
	X, Y   int
}

func NewTransform(id uint32, entity entity.Entity, x, y int) *Transform {
	return &Transform{
		CID:    id,
		entity: entity,
		X:      x,
		Y:      y,
	}
}

func (t *Transform) ID() uint32            { return t.CID }
func (t *Transform) Type() string          { return TransformType }
func (t *Transform) Entity() entity.Entity { return t.entity }
