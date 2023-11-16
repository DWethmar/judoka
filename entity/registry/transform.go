package registry

import (
	"fmt"

	"github.com/dwethmar/judoka/component"
	"github.com/dwethmar/judoka/entity"
)

// AddTransform adds a transform.
func (r *Registry) AddTransform(c *component.Transform) error {
	c.CID = r.TransformIDGenerator.Next()
	if err := r.Transform.AddComponent(c); err != nil {
		return fmt.Errorf("failed to add transform: %w", err)
	}
	return nil
}

// RemoveTransform removes a transform.
func (r *Registry) RemoveTransform(c *component.Transform) { r.Transform.RemoveComponent(c) }

// ListTransforms returns a list of all transforms.
func (r *Registry) GetTransform(e entity.Entity) *component.Transform {
	if t := r.Transform.GetComponents(e); t != nil {
		return t[0]
	}
	return nil
}
