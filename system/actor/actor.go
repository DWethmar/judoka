package actor

import (
	"fmt"
	"image/color"
	"log/slog"

	"github.com/dwethmar/judoka/assets"
	"github.com/dwethmar/judoka/component"
	"github.com/dwethmar/judoka/entity"
	"github.com/dwethmar/judoka/entity/registry"
	"github.com/dwethmar/judoka/system"
	"github.com/dwethmar/judoka/transform"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

var _ system.System = &System{}

// SubSystem is a system for a specific actor type.
type SubSystem interface {
	Init(root entity.Entity) error
	Update(actors []*component.Actor) error
	ActorType() component.ActorType
}

type System struct {
	logger             *slog.Logger
	register           *registry.Register
	positionResolution int // used to divide X and Y positions
	rootEntity         entity.Entity
	SubSystems         map[component.ActorType]SubSystem
}

// Options are used to configure a new actor system.
type Options struct {
	Logger             *slog.Logger
	Register           *registry.Register
	PositionResolution int
	ActorSubSystems    []SubSystem
}

func New(opt Options) *System {
	actorSubSystems := map[component.ActorType]SubSystem{}

	for _, subSystem := range opt.ActorSubSystems {
		actorSubSystems[subSystem.ActorType()] = subSystem
	}

	return &System{
		logger:             opt.Logger.WithGroup("actor"),
		register:           opt.Register,
		positionResolution: opt.PositionResolution,
		SubSystems:         actorSubSystems,
	}
}

// Init initializes the system and all sub systems.
func (s *System) Init() error {
	rootEntity, err := s.register.Create(s.register.Root())
	if err != nil {
		return fmt.Errorf("failed to create root entity: %w", err)
	}

	s.rootEntity = rootEntity

	for _, subSystem := range s.SubSystems {
		if err := subSystem.Init(s.rootEntity); err != nil {
			return fmt.Errorf("failed to init sub system: %w", err)
		}
	}

	return nil
}

// Draw implements system.System.
func (s *System) Draw(screen *ebiten.Image) error {
	for _, e := range s.register.Actor.Entities() {
		t, ok := s.register.Transform.First(e)
		if !ok {
			continue
		}

		var x, y = transform.Position(s.register, e)

		x /= s.positionResolution
		y /= s.positionResolution

		velocity, ok := s.register.Velocity.First(e)
		if !ok {
			continue
		}

		actor, ok := s.register.Actor.First(e)
		if !ok {
			continue
		}

		layer := 0
		if l, ok := s.register.Layer.First(e); ok {
			layer = l.Index
		}

		text.Draw(screen, fmt.Sprintf(`POS x: %d (%d), y: %d (%d)
VEL x: %d, y: %d
ACTOR:%d
Facing:%s
AnimationFrame:%d
Layer:%d
`,
			t.X,
			t.X/s.positionResolution,
			t.Y,
			t.Y/s.positionResolution,
			velocity.X, velocity.Y,
			actor.ActorType,
			actor.Facing,
			actor.AnimationFrame,
			layer,
		), assets.GetVGAFonts(2), x, y+30, color.White)
	}

	return nil
}

// Update implements system.System.
func (s *System) Update() error {
	actorsByType := map[component.ActorType][]*component.Actor{}

	for _, e := range s.register.Actor.Entities() {
		actor, ok := s.register.Actor.First(e)
		if !ok {
			continue
		}

		actorsByType[actor.ActorType] = append(actorsByType[actor.ActorType], actor)
	}

	for _, subSystem := range s.SubSystems {
		if err := subSystem.Update(actorsByType[subSystem.ActorType()]); err != nil {
			return fmt.Errorf("failed to update sub system: %w", err)
		}
	}

	return nil
}
