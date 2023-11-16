package component

import "github.com/dwethmar/judoka/entity"

const VelocityType = "velocity"

type Velocity struct {
	CID    uint32
	entity entity.Entity
	X, Y   int
}

func NewVelocity(id uint32, entity entity.Entity, x, y int) *Velocity {
	return &Velocity{
		CID:    id,
		entity: entity,
		X:      x,
		Y:      y,
	}
}

func (t *Velocity) ID() uint32            { return t.CID }
func (t *Velocity) Type() string          { return VelocityType }
func (t *Velocity) Entity() entity.Entity { return t.entity }
