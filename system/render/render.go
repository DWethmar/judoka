package render

import (
	"log/slog"

	"github.com/dwethmar/judoka/component"
	"github.com/dwethmar/judoka/entity"
	"github.com/dwethmar/judoka/entity/registry"
	"github.com/dwethmar/judoka/system"
	"github.com/dwethmar/judoka/transform"
	"github.com/hajimehoshi/ebiten/v2"
)

var _ system.System = (*System)(nil)

// System is a render system.
type System struct {
	logger             *slog.Logger
	register           *registry.Register
	positionResolution int // used to divide X and Y positions
}

// Options are used to configure a new render system.
type Options struct {
	Logger             *slog.Logger
	Register           *registry.Register
	PositionResolution int
}

// New creates a new Render system.
func New(
	opt Options,
) *System {
	return &System{
		logger:             opt.Logger,
		register:           opt.Register,
		positionResolution: opt.PositionResolution,
	}
}

func (s *System) Init() error {
	return nil
}

// Draw implements system.System.
func (r *System) Draw(screen *ebiten.Image) error {
	spriteLookup := map[entity.Entity][]*component.Sprite{}

	for _, e := range r.register.Sprite.Entities() {
		for _, sprite := range r.register.Sprite.List(e) {
			spriteLookup[e] = append(spriteLookup[e], sprite)
		}
	}

	for _, e := range r.register.List() {
		sprites, ok := spriteLookup[e]
		if !ok {
			continue
		}
		x, y := transform.Position(r.register, e)
		for _, sprite := range sprites {
			r.drawSprite(screen, x, y, sprite)
		}
	}

	return nil
}

func (r *System) drawSprite(screen *ebiten.Image, x, y int, sprite *component.Sprite) {
	nX := x / r.positionResolution
	nY := y / r.positionResolution

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(
		float64(nX+sprite.OffsetX),
		float64(nY+sprite.OffsetY),
	)
	screen.DrawImage(sprite.Image, op)
}

// Update implements system.System.
func (*System) Update() error {
	return nil
}
