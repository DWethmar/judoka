package component

import (
	"github.com/dwethmar/judoka/entity"
	"github.com/hajimehoshi/ebiten/v2"
)

const SpriteType = "sprite"

var _ Component = (*Sprite)(nil)

type Sprite struct {
	CID     uint32
	entity  entity.Entity
	OffsetX float64
	OffsetY float64
	Image   *ebiten.Image
}

func NewSprite(id uint32, entity entity.Entity, offsetX, offsetY float64, image *ebiten.Image) *Sprite {
	return &Sprite{
		CID:     id,
		entity:  entity,
		OffsetX: offsetX,
		OffsetY: offsetY,
		Image:   image,
	}
}

func (s *Sprite) ID() uint32            { return s.CID }
func (s *Sprite) Type() string          { return SpriteType }
func (s *Sprite) Entity() entity.Entity { return s.entity }
