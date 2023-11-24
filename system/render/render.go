package render

import (
	"log/slog"

	"github.com/dwethmar/judoka/entity/registry"
	"github.com/dwethmar/judoka/system"
	"github.com/dwethmar/judoka/transform"
	"github.com/hajimehoshi/ebiten/v2"
)

var _ system.System = (*System)(nil)

// System is a render system.
type System struct {
	logger             *slog.Logger
	registry           *registry.Registry
	positionResolution int // used to divide X and Y positions
}

// Options are used to configure a new render system.
type Options struct {
	Logger             *slog.Logger
	Registry           *registry.Registry
	PositionResolution int
}

// New creates a new Render system.
func New(
	opt Options,
) *System {
	return &System{
		logger:             opt.Logger,
		registry:           opt.Registry,
		positionResolution: opt.PositionResolution,
	}
}

// Draw implements system.System.
func (r *System) Draw(screen *ebiten.Image) error {
	for _, e := range r.registry.Sprite.Entities() {
		x, y := transform.Position(r.registry, e)

		for _, sprite := range r.registry.Sprite.List(e) {
			nX := x / r.positionResolution
			nY := y / r.positionResolution

			x := float64(nX + sprite.OffsetX)
			y := float64(nY + sprite.OffsetY)

			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(x, y)
			screen.DrawImage(sprite.Image, op)
		}
	}

	return nil
}

// Update implements system.System.
func (*System) Update() error {
	return nil
}
