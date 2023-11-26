package camera

import (
	"log/slog"

	"github.com/dwethmar/judoka/entity"
	"github.com/dwethmar/judoka/entity/registry"
	"github.com/dwethmar/judoka/system"
	ebiten "github.com/hajimehoshi/ebiten/v2"
)

var _ system.System = (*System)(nil)

type Options struct {
	Logger             *slog.Logger
	Register           *registry.Register
	PositionResolution int
	Viewport           entity.Entity
}

type System struct {
	logger             *slog.Logger
	register           *registry.Register
	positionResolution int
	viewport           entity.Entity
	Follow             entity.Entity
}

func New(opt Options) *System {
	return &System{
		logger:             opt.Logger,
		register:           opt.Register,
		positionResolution: opt.PositionResolution,
		viewport:           opt.Viewport,
	}
}

// Init implements system.System.
func (s *System) Init() error {
	s.Follow = s.register.Controller.Entities()[0]
	return nil
}

// Draw implements system.System.
func (s *System) Draw(screen *ebiten.Image) error {
	return nil
}

// Update implements system.System.
func (s *System) Update() error {
	// Get the transform of the entity we want to follow.
	transform, ok := s.register.Transform.First(s.Follow)
	if !ok {
		return nil
	}

	// Get the transform of the viewport.
	viewportTransform, ok := s.register.Transform.First(s.viewport)
	if !ok {
		return nil
	}

	// Get the size of the viewport.
	widht, height := ebiten.WindowSize()

	// Calculate the center of the viewport.
	centerX := (widht) / 2
	centerY := (height) / 2

	// Calculate the position of the entity we want to follow.
	x := (transform.X - centerX) / s.positionResolution
	y := (transform.Y - centerY) / s.positionResolution

	// Calculate the position of the viewport.
	viewportX := viewportTransform.X
	viewportY := viewportTransform.Y

	// Calculate the difference between the entity and the viewport.
	diffX := x - viewportX
	diffY := y - viewportY

	// Calculate the new position of the viewport.
	newViewportX := viewportX + diffX
	newViewportY := viewportY + diffY

	// Calculate the new position of the viewport.
	viewportTransform.X = newViewportX
	viewportTransform.Y = newViewportY

	return nil
}
