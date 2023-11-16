package registry

import (
	"errors"

	"github.com/dwethmar/judoka/component"
)

var (
	ErrComponentTypeEmpty   = errors.New("component type is empty")
	ErrComponentEntityEmpty = errors.New("component entity is empty")
)

// ValidateComponent is a business rule that validates a component.
func ValidateComponent(c component.Component) error {
	if c == nil {
		return errors.New("component is nil")
	}

	if c.ID() <= 0 {
		return errors.New("component id is empty")
	}

	if c.Type() == "" {
		return ErrComponentTypeEmpty
	}

	if c.Entity() <= 0 {
		return ErrComponentEntityEmpty
	}

	return nil
}
