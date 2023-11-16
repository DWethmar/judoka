package velocity

import (
	"log/slog"

	"github.com/dwethmar/judoka/entity/registry"
	"github.com/dwethmar/judoka/system"
	"github.com/hajimehoshi/ebiten/v2"
)

const drag = 1

var _ system.System = (*System)(nil)

type System struct {
	logger   *slog.Logger
	registry *registry.Registry
}

// New creates a new velocity system.
func New(logger *slog.Logger, registry *registry.Registry) *System {
	return &System{
		logger:   logger,
		registry: registry,
	}
}

// Update implements system.System.
func (s *System) Update() error {
	for _, c := range s.registry.ListVelocities() {
		transform := s.registry.GetTransform(c.Entity())
		if transform == nil {
			continue
		}

		transform.X += c.X
		transform.Y += c.Y

		// Apply drag
		if c.X > 0 {
			c.X -= drag
			if c.X < 0 {
				c.X = 0
			}
		}

		if c.X < 0 {
			c.X += drag
			if c.X > 0 {
				c.X = 0
			}
		}

		if c.Y > 0 {
			c.Y -= drag
			if c.Y < 0 {
				c.Y = 0
			}
		}

		if c.Y < 0 {
			c.Y += drag
			if c.Y > 0 {
				c.Y = 0
			}
		}
	}

	return nil
}

// Draw implements system.System.
func (*System) Draw(screen *ebiten.Image) error {
	return nil
}
