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

// Camera is a shared struct that is used to follow an entity.
type Camera struct {
	Viewport           entity.Entity
	currentX, currentY int
	targetX, targetY   int
	Bounds             image.Rectangle
}

func (c *Camera) Target(x, y int) {
	c.targetX = x
	c.targetY = y
}

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
	camera             *Camera
}

func New(opt Options) *System {
	return &System{
		logger:             opt.Logger,
		register:           opt.Register,
		positionResolution: opt.PositionResolution,
		camera: &Camera{
			Viewport: opt.Viewport,
		},
		screenBounds: image.Rectangle{},
	}
}

// Init implements system.System.
func (s *System) Init() error {
	return nil
}

func (s *System) Camera() *Camera {
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

	// Get the transform of the viewport.
	viewportTransform, ok := s.register.Transform.First(s.camera.Viewport)
	if !ok {
		return nil
	}

	// Smooth transition: move current towards target with an offset
	offset := 0.1 // Adjust this value to change the smoothness
	s.camera.currentX += int(float64(s.camera.targetX-s.camera.currentX) * offset)
	s.camera.currentY += int(float64(s.camera.targetY-s.camera.currentY) * offset)

	// set viewport to center of the screen
	viewportTransform.X = (screenWidth / 2) * s.positionResolution
	viewportTransform.Y = (screenHeight / 2) * s.positionResolution

	// Set the viewport to the position of the entity we want to follow.
	viewportTransform.X -= s.camera.currentX
	viewportTransform.Y -= s.camera.currentY

	// set bounds on camera
	s.camera.Bounds = image.Rectangle{
		Min: image.Point{
			X: (s.camera.currentX / s.positionResolution) - (screenWidth / 2),
			Y: (s.camera.currentY / s.positionResolution) - (screenHeight / 2),
		},
		Max: image.Point{
			X: (s.camera.currentX / s.positionResolution) + (screenWidth / 2),
			Y: (s.camera.currentY / s.positionResolution) + (screenHeight / 2),
		},
	}

	return nil
}
