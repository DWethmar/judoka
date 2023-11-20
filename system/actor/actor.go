package actor

import (
	"fmt"
	"image/color"
	"log/slog"

	"github.com/dwethmar/judoka/assets"
	"github.com/dwethmar/judoka/component"
	"github.com/dwethmar/judoka/entity/registry"
	"github.com/dwethmar/judoka/system"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

var _ system.System = &System{}

type System struct {
	logger   *slog.Logger
	registry *registry.Registry
	managers map[component.ActorType]Manager
}

func New(logger *slog.Logger, registry *registry.Registry) *System {
	return &System{
		logger:   logger.WithGroup("player"),
		registry: registry,
		managers: map[component.ActorType]Manager{
			component.ActorTypePlayer: NewPlayerManager(logger, registry),
		},
	}
}

// Draw implements system.System.
func (s *System) Draw(screen *ebiten.Image) error {
	for _, e := range s.registry.Actor.Entities() {
		transform, ok := s.registry.Transform.First(e)
		if !ok {
			continue
		}

		x := transform.X / system.PositionResolution
		y := transform.Y / system.PositionResolution

		text.Draw(screen, fmt.Sprintf("TRANS x: %d (%d), y: %d (%d)", transform.X, x, transform.Y, y), assets.GetVGAFonts(2), x, y, color.White)

		velocity, ok := s.registry.Velocity.First(e)
		if !ok {
			continue
		}

		text.Draw(screen, fmt.Sprintf("VELOC x: %d, y: %d", velocity.X, velocity.Y), assets.GetVGAFonts(2), x, y+15, color.White)

		actor, ok := s.registry.Actor.First(e)
		if !ok {
			continue
		}

		text.Draw(screen, fmt.Sprintf(`ACTOR %d
Facing %d
AnimationFrame: %d
`,
			actor.ActorType,
			actor.Facing,
			actor.AnimationFrame,
		), assets.GetVGAFonts(2), x, y+30, color.White)
	}

	return nil
}

// Update implements system.System.
func (s *System) Update() error {
	for _, e := range s.registry.Actor.Entities() {
		actor, ok := s.registry.Actor.First(e)
		if !ok {
			continue
		}

		manager, ok := s.managers[actor.ActorType]
		if !ok {
			continue
		}

		if err := manager.Update(actor); err != nil {
			return fmt.Errorf("failed to update actor: %w", err)
		}
	}

	return nil
}

func GetOffsets(image *ebiten.Image) (int, int) {
	w := image.Bounds().Max.X - image.Bounds().Min.X
	h := image.Bounds().Max.Y - image.Bounds().Min.Y
	return -(w / 2), -h
}
