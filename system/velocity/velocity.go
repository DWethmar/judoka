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
	for _, e := range s.registry.Velocity.Entities() {
		transform, ok := s.registry.Transform.First(e)
		if !ok {
			continue
		}

		velocity, ok := s.registry.Velocity.First(e)
		if !ok {
			continue
		}

		if velocity.X == 0 && velocity.Y == 0 {
			continue
		}

		actor, ok := s.registry.Actor.First(e)
		if ok {
			actor.PreviousX = transform.X
			actor.PreviousY = transform.Y
		}

		transform.X += velocity.X
		transform.Y += velocity.Y

		// Apply drag
		if velocity.X > 0 {
			velocity.X -= drag
			if velocity.X < 0 {
				velocity.X = 0
			}
		}

		if velocity.X < 0 {
			velocity.X += drag
			if velocity.X > 0 {
				velocity.X = 0
			}
		}

		if velocity.Y > 0 {
			velocity.Y -= drag
			if velocity.Y < 0 {
				velocity.Y = 0
			}
		}

		if velocity.Y < 0 {
			velocity.Y += drag
			if velocity.Y > 0 {
				velocity.Y = 0
			}
		}
	}

	return nil
}

// Draw implements system.System.
func (*System) Draw(screen *ebiten.Image) error {
	return nil
}
