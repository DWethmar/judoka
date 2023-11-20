package component

import (
	"github.com/dwethmar/judoka/direction"
	"github.com/dwethmar/judoka/entity"
)

const ActorComponentType = "Actor"

type ActorType int

const (
	ActorTypeUnknown ActorType = iota
	ActorTypePlayer
	ActorTypeEnemy
)

type Actor struct {
	CID            uint32
	entity         entity.Entity
	ActorType      ActorType
	Facing         direction.Direction
	AnimationFrame int
}

func NewActor(id uint32, entity entity.Entity) *Actor {
	return &Actor{
		CID:    id,
		entity: entity,
	}
}

func (c *Actor) ID() uint32            { return c.CID }
func (c *Actor) Type() string          { return ActorComponentType }
func (c *Actor) Entity() entity.Entity { return c.entity }
