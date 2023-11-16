package component

import "github.com/dwethmar/judoka/entity"

type Component interface {
	ID() uint32
	Type() string
	Entity() entity.Entity
}
