package actor

import (
	"github.com/dwethmar/judoka/entity/registry"
	"github.com/dwethmar/judoka/level"
)

type Generator interface {
	Generate(level *level.Level, registry *registry.Register) error
}
