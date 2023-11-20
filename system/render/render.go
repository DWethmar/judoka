package render

import (
	"log/slog"

	"github.com/dwethmar/judoka/entity/registry"
	"github.com/dwethmar/judoka/system"
	"github.com/dwethmar/judoka/transform"
	"github.com/hajimehoshi/ebiten/v2"
)

var _ system.System = (*System)(nil)

type System struct {
	logger   *slog.Logger
	registry *registry.Registry
}

// New creates a new Render system.
func New(logger *slog.Logger, registry *registry.Registry) *System {
	return &System{
		logger:   logger,
		registry: registry,
	}
}

// Draw implements system.System.
func (r *System) Draw(screen *ebiten.Image) error {
	for _, e := range r.registry.Sprite.Entities() {
		x, y := transform.Position(r.registry, e)

		for _, sprite := range r.registry.Sprite.List(e) {
			nX := x / system.PositionResolution
			nY := y / system.PositionResolution

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
