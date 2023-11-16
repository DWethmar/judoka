package render

import (
	"log/slog"

	"github.com/dwethmar/judoka/entity/registry"
	"github.com/dwethmar/judoka/system"
	"github.com/hajimehoshi/ebiten/v2"
)

var _ system.System = (*Render)(nil)

type Render struct {
	logger   *slog.Logger
	registry *registry.Registry
}

// New creates a new Render system.
func New(logger *slog.Logger, registry *registry.Registry) *Render {
	return &Render{
		logger:   logger,
		registry: registry,
	}
}

// Debug implements system.System.
func (*Render) Debug(screen *ebiten.Image) error {
	return nil
}

// Draw implements system.System.
func (r *Render) Draw(screen *ebiten.Image) error {
	for _, sprite := range r.registry.ListSprites() {
		entity := sprite.Entity()
		transform := r.registry.GetTransform(entity)
		if transform == nil {
			continue
		}

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(transform.X+sprite.OffsetX, transform.Y+sprite.OffsetY)
		screen.DrawImage(sprite.Image, op)
	}

	return nil
}

// Update implements system.System.
func (*Render) Update() error {
	return nil
}
