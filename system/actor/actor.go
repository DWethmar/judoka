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
	logger             *slog.Logger
	registry           *registry.Registry
	PositionResolution int // used to divide X and Y positions
	managers           map[component.ActorType]Manager
}

// Options are used to configure a new actor system.
type Options struct {
	Logger             *slog.Logger
	Registry           *registry.Registry
	PositionResolution int
	Managers           map[component.ActorType]Manager
}

func New(opt Options) *System {
	return &System{
		logger:             opt.Logger.WithGroup("actor"),
		registry:           opt.Registry,
		PositionResolution: opt.PositionResolution,
		managers:           opt.Managers,
	}
}

// Draw implements system.System.
func (s *System) Draw(screen *ebiten.Image) error {
	for _, e := range s.registry.Actor.Entities() {
		transform, ok := s.registry.Transform.First(e)
		if !ok {
			continue
		}

		x := transform.X / s.PositionResolution
		y := transform.Y / s.PositionResolution

		text.Draw(screen, fmt.Sprintf("POS x: %d (%d), y: %d (%d)", transform.X, x, transform.Y, y), assets.GetVGAFonts(2), x, y, color.White)

		velocity, ok := s.registry.Velocity.First(e)
		if !ok {
			continue
		}

		text.Draw(screen, fmt.Sprintf("VEL x: %d, y: %d", velocity.X, velocity.Y), assets.GetVGAFonts(2), x, y+15, color.White)

		actor, ok := s.registry.Actor.First(e)
		if !ok {
			continue
		}

		text.Draw(screen, fmt.Sprintf(`ACTOR %d
Facing %s
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
			s.logger.Error("no actor found for entity (which is weird)", slog.Int("entity", int(e)))
			continue
		}

		manager, ok := s.managers[actor.ActorType]
		if !ok {
			s.logger.Info("no manager found for actor type", slog.Int("actor_type", int(actor.ActorType)))
			continue
		}

		if err := manager.Update(actor); err != nil {
			return fmt.Errorf("failed to update actor: %w", err)
		}
	}

	return nil
}
