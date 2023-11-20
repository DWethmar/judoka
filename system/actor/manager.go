package actor

import "github.com/dwethmar/judoka/component"

type Manager interface {
	Update(actor *component.Actor) error
}
