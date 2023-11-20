package component

import (
	"github.com/dwethmar/judoka/entity"
	"github.com/dwethmar/judoka/matrix"
)

const ChunkType = "Chunk"

type Chunk struct {
	CID     uint32
	CEntity entity.Entity
	X, Y    int
	Tiles   matrix.Matrix
}

func NewChunk(id uint32, entity entity.Entity) *Chunk {
	return &Chunk{
		CID:     id,
		CEntity: entity,
	}
}

func (c *Chunk) ID() uint32            { return c.CID }
func (c *Chunk) Type() string          { return ChunkType }
func (c *Chunk) Entity() entity.Entity { return c.CEntity }
