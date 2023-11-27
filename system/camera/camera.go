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

	screenWidth  int
	screenHeight int
}

func New(opt Options) *System {
	return &System{
		logger:             opt.Logger,
		register:           opt.Register,
		positionResolution: opt.PositionResolution,
		viewport:           opt.Viewport,

		screenWidth:  0,
		screenHeight: 0,
	}
}

// Init implements system.System.
func (s *System) Init() error {
	s.Follow = s.register.Controller.Entities()[0]
	return nil
}

// Draw implements system.System.
func (s *System) Draw(screen *ebiten.Image) error {
	s.screenWidth, s.screenHeight = screen.Bounds().Dx(), screen.Bounds().Dy()
	return nil
}

// Update implements system.System.
func (s *System) Update() error {
	if s.screenWidth == 0 || s.screenHeight == 0 {
		return nil
	}

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

	// set viewport to center of the screen
	viewportTransform.X = (s.screenWidth / 2) * s.positionResolution
	viewportTransform.Y = (s.screenHeight / 2) * s.positionResolution

	// Set the viewport to the position of the entity we want to follow.
	viewportTransform.X -= transform.X
	viewportTransform.Y -= transform.Y

	return nil
}
