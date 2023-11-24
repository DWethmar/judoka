package player

import (
	"fmt"
	"log/slog"

	"github.com/dwethmar/judoka/component"
	"github.com/dwethmar/judoka/entity"
	"github.com/dwethmar/judoka/entity/registry"
	"github.com/dwethmar/judoka/system/actor"
)

var _ actor.SubSystem = &System{}

// System is a system for a specific actor type.
type System struct {
	logger   *slog.Logger
	register *registry.Register
	root     entity.Entity // root entity for actors
}

// Options are used to configure a new player manager.
type Options struct {
	Logger   *slog.Logger
	Register *registry.Register
}

func New(opt Options) *System {
	return &System{
		logger:   opt.Logger,
		register: opt.Register,
	}
}

func (m *System) Init(root entity.Entity) error {
	m.root = root
	return nil
}

func (m *System) ActorType() component.ActorType { return component.ActorTypePlayer }

func (m *System) Update(actors []*component.Actor) error {
	for _, actor := range actors {
		if err := m.animate(actor); err != nil {
			return fmt.Errorf("failed to animate actor: %w", err)
		}
	}

	return nil
}
