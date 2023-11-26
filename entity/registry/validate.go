package registry

import (
	"errors"

	"github.com/dwethmar/judoka/component"
)

var (
	ErrComponentTypeRequied    = errors.New("component type is not set")
	ErrComponentEntityRequired = errors.New("component entity is not set")
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
		return ErrComponentTypeRequied
	}

	if c.Entity() <= 0 {
		return ErrComponentEntityRequired
	}

	return nil
}
