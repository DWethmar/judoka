package debug

import (
	"fmt"
	"image/color"
	"log/slog"

	"github.com/dwethmar/judoka/assets"
	"github.com/dwethmar/judoka/entity/registry"
	"github.com/dwethmar/judoka/system"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

type System struct {
	logger   *slog.Logger
	registry *registry.Registry
}

// New creates a new debug system.
func New(logger *slog.Logger, registry *registry.Registry) *System {
	return &System{
		logger:   logger,
		registry: registry,
	}
}

// Draw implements system.System.
func (s *System) Draw(screen *ebiten.Image) error {
	for _, c := range s.registry.ListControllers() {
		entity := c.Entity()
		transform := s.registry.GetTransform(entity)
		if transform == nil {
			continue
		}

		x := transform.X / system.PositionResolution
		y := transform.Y / system.PositionResolution

		text.Draw(screen, fmt.Sprintf("TRANS x: %d (%d), y: %d (%d)", transform.X, x, transform.Y, y), assets.GetVGAFonts(2), x, y, color.White)

		velocity := s.registry.GetVelocity(entity)
		if velocity == nil {
			continue
		}

		text.Draw(screen, fmt.Sprintf("VELOC x: %d, y: %d", velocity.X, velocity.Y), assets.GetVGAFonts(2), x, y+15, color.White)
	}

	return nil
}

// Update implements system.System.
func (*System) Update() error {
	return nil
}
