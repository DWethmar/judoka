package component

import (
	"github.com/dwethmar/judoka/entity"
)

const LayerType = "layer"

var _ Component = (*Layer)(nil)

// Layer is a component that stores the rendering layer that an entity belongs to.
type Layer struct {
	CID    uint32
	entity entity.Entity
	Index  int
}

func NewLayer(id uint32, entity entity.Entity) *Layer {
	return &Layer{
		CID:    id,
		entity: entity,
		Index:  0,
	}
}

func (s *Layer) ID() uint32            { return s.CID }
func (s *Layer) Type() string          { return LayerType }
func (s *Layer) Entity() entity.Entity { return s.entity }
