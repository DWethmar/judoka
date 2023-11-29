package camera

import (
	"image"
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
	screenBounds       image.Rectangle
	camera             *system.Camera
}

func New(opt Options) *System {
	return &System{
		logger:             opt.Logger,
		register:           opt.Register,
		positionResolution: opt.PositionResolution,
		camera: &system.Camera{
			Viewport: opt.Viewport,
		},
		screenBounds: image.Rectangle{},
	}
}

// Init implements system.System.
func (s *System) Init() error {
	s.camera.Following = s.register.Controller.Entities()[0]
	return nil
}

func (s *System) Camera() *system.Camera {
	return s.camera
}

// Draw implements system.System.
func (s *System) Draw(screen *ebiten.Image) error {
	s.screenBounds = screen.Bounds()
	return nil
}

// Update implements system.System.
func (s *System) Update() error {
	if s.screenBounds.Empty() {
		return nil
	}

	screenWidth := s.screenBounds.Dx()
	screenHeight := s.screenBounds.Max.Y

	// Get the transform of the entity we want to follow.
	transform, ok := s.register.Transform.First(s.camera.Following)
	if !ok {
		return nil
	}

	// Get the transform of the viewport.
	viewportTransform, ok := s.register.Transform.First(s.camera.Viewport)
	if !ok {
		return nil
	}

	// set viewport to center of the screen
	viewportTransform.X = (screenWidth / 2) * s.positionResolution
	viewportTransform.Y = (screenHeight / 2) * s.positionResolution

	// Set the viewport to the position of the entity we want to follow.
	viewportTransform.X -= transform.X
	viewportTransform.Y -= transform.Y

	// set bounds on camera
	s.camera.Bounds = image.Rectangle{
		Min: image.Point{
			X: (transform.X / s.positionResolution) - (screenWidth / 2),
			Y: (transform.Y / s.positionResolution) - (screenHeight / 2),
		},
		Max: image.Point{
			X: (transform.X / s.positionResolution) + (screenWidth / 2),
			Y: (transform.Y / s.positionResolution) + (screenHeight / 2),
		},
	}

	return nil
}
