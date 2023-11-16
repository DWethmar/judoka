package registry

import (
	"fmt"

	"github.com/dwethmar/judoka/component"
	"github.com/dwethmar/judoka/entity"
)

// AddController adds a transform.
func (r *Registry) AddController(c *component.Controller) error {
	c.CID = r.ControllerIDGenerator.Next()
	if err := r.Controller.AddComponent(c); err != nil {
		return fmt.Errorf("failed to add controller: %w", err)
	}
	return nil
}

// RemoveController removes a controller
func (r *Registry) RemoveController(c *component.Controller) { r.Controller.RemoveComponent(c) }

// GetController returns a controller
func (r *Registry) GetController(e entity.Entity) *component.Controller {
	if t := r.Controller.GetComponents(e); t != nil {
		return t[0]
	}
	return nil
}

// ListControllers returns a list of all controllers.
func (r *Registry) ListControllers() []*component.Controller {
	controllers := []*component.Controller{}
	for _, e := range r.Controller.ListEntities() {
		if t := r.Controller.GetComponents(e); t != nil {
			controllers = append(controllers, t...)
		}
	}

	return controllers
}
