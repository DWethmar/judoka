package input

import (
	"fmt"
	"log/slog"
	"math"

	"github.com/dwethmar/judoka/entity/registry"
	"github.com/dwethmar/judoka/system"
	ebiten "github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const defaultSpeed = 2

var _ system.System = (*System)(nil)

type System struct {
	logger             *slog.Logger
	register           *registry.Register
	PositionResolution int // used to divide X and Y positions
}

// Options are used to configure a new input system.
type Options struct {
	Logger             *slog.Logger
	Register           *registry.Register
	PositionResolution int
}

// New creates a new input system.
func New(opt Options) *System {
	return &System{
		logger:             opt.Logger,
		register:           opt.Register,
		PositionResolution: opt.PositionResolution,
	}
}

func (s *System) Init() error {
	return nil
}

// Draw implements system.System.
func (*System) Draw(screen *ebiten.Image) error {
	dx, dy := Direction()
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("INPUT X: %d, Y: %d", dx, dy), 500, 50)
	return nil
}

// Update implements system.System.
func (s *System) Update() error {
	dx, dy := Direction()
	if dx == 0 && dy == 0 {
		// Reset controller
		for _, e := range s.register.Controller.Entities() {
			controller, ok := s.register.Controller.First(e)
			if !ok {
				continue
			}

			controller.X = 0
			controller.Y = 0
		}

		return nil
	}

	// Calculate the length of the vector (dx, dy)
	length := math.Sqrt(float64(dx*dx + dy*dy))

	// Normalize dx and dy
	normalizedDx := float64(dx) / length
	normalizedDy := float64(dy) / length

	// Use the normalized values for transformation
	for _, e := range s.register.Controller.Entities() {
		vel := s.register.Velocity.List(e)[0]
		if vel == nil {
			continue
		}

		// Update position based on normalized direction
		vel.X = int(math.Round(normalizedDx)) * defaultSpeed * s.PositionResolution
		vel.Y = int(math.Round(normalizedDy)) * defaultSpeed * s.PositionResolution

		// update controller
		controller, ok := s.register.Controller.First(e)
		if !ok {
			continue
		}

		controller.X = dx
		controller.Y = dy
	}

	return nil
}
