package registry

import (
	"fmt"

	"github.com/dwethmar/judoka/component"
	"github.com/dwethmar/judoka/entity"
)

// AddVelocity adds a velocity.
func (r *Registry) AddVelocity(c *component.Velocity) error {
	c.CID = r.VelocityIDGenerator.Next()
	if err := r.Velocity.AddComponent(c); err != nil {
		return fmt.Errorf("failed to add velocity: %w", err)
	}
	return nil
}

// RemoveVelocity removes a velocity.
func (r *Registry) RemoveVelocity(c *component.Velocity) { r.Velocity.RemoveComponent(c) }

// ListVelocities returns a list of all velocities.
func (r *Registry) GetVelocity(e entity.Entity) *component.Velocity {
	if t := r.Velocity.GetComponents(e); t != nil {
		return t[0]
	}
	return nil
}

// ListVelocities returns a list of all velocities.
func (r *Registry) ListVelocities() []*component.Velocity {
	velocities := []*component.Velocity{}
	for _, e := range r.Velocity.ListEntities() {
		if s := r.Velocity.GetComponents(e); s != nil {
			velocities = append(velocities, s...)
		}
	}

	return velocities
}
