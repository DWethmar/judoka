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
	register *registry.Register
}

// Options are used to configure a new velocity system.
type Options struct {
	Logger   *slog.Logger
	Register *registry.Register
}

// New creates a new velocity system.
func New(opt Options) *System {
	return &System{
		logger:   opt.Logger,
		register: opt.Register,
	}
}

func (s *System) Init() error {
	return nil
}

// Update implements system.System.
func (s *System) Update() error {
	for _, e := range s.register.Velocity.Entities() {
		transform, ok := s.register.Transform.First(e)
		if !ok {
			continue
		}

		velocity, ok := s.register.Velocity.First(e)
		if !ok {
			continue
		}

		if velocity.X == 0 && velocity.Y == 0 {
			continue
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
