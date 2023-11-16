package registry

import (
	"errors"
	"sync"

	"github.com/dwethmar/judoka/component"
	"github.com/dwethmar/judoka/entity"
)

var (
	ErrUnique = errors.New("component is unique")
)

type stores struct {
	TransformIDGenerator  *IDGenerator
	Transform             *Store[*component.Transform]
	SpriteIDGenerator     *IDGenerator
	Sprite                *Store[*component.Sprite]
	ControllerIDGenerator *IDGenerator
	Controller            *Store[*component.Controller]
	VelocityIDGenerator   *IDGenerator
	Velocity              *Store[*component.Velocity]
}

func (s *stores) DeleteEntity(e entity.Entity) {
	s.Transform.RemoveEntity(e)
	s.Sprite.RemoveEntity(e)
	s.Controller.RemoveEntity(e)
	s.Velocity.RemoveEntity(e)
}

// Registry keeps track of all entities and their components.
// It provides methods for entity creation and component management.
type Registry struct {
	stores
	idGenerator *IDGenerator
	mux         sync.RWMutex
	entities    map[entity.Entity]struct{}
}

// New creates and returns a new instance of Registry.
// It initializes the entity map and the component stores.
func New() *Registry {
	return &Registry{
		stores: stores{
			TransformIDGenerator: NewIDGenerator(0),
			Transform: NewStore[*component.Transform](
				func(s *Store[*component.Transform]) {
					s.BeforeAdd = func(c component.Component) error {
						if err := ValidateComponent(c); err != nil {
							return err
						}

						// unique component
						if len(s.store[c.Entity()]) > 0 {
							return ErrUnique
						}

						return nil
					}
				},
			),
			SpriteIDGenerator: NewIDGenerator(0),
			Sprite: NewStore[*component.Sprite](
				func(s *Store[*component.Sprite]) {
					s.BeforeAdd = func(c component.Component) error {
						if err := ValidateComponent(c); err != nil {
							return err
						}

						return nil
					}
				},
			),
			ControllerIDGenerator: NewIDGenerator(0),
			Controller: NewStore[*component.Controller](
				func(s *Store[*component.Controller]) {
					s.BeforeAdd = func(c component.Component) error {
						if err := ValidateComponent(c); err != nil {
							return err
						}

						// unique component
						if len(s.store[c.Entity()]) > 0 {
							return ErrUnique
						}

						return nil
					}
				},
			),
			VelocityIDGenerator: NewIDGenerator(0),
			Velocity: NewStore[*component.Velocity](
				func(s *Store[*component.Velocity]) {
					s.BeforeAdd = func(c component.Component) error {
						if err := ValidateComponent(c); err != nil {
							return err
						}

						// unique component
						if len(s.store[c.Entity()]) > 0 {
							return ErrUnique
						}

						return nil
					}
				},
			),
		},
		idGenerator: NewIDGenerator(0),
		mux:         sync.RWMutex{},
		entities:    make(map[entity.Entity]struct{}),
	}
}

// CreateEntity generates a new unique Entity ID, registers it as the child of the given parent, and returns it.
// If no parent is specified, and there is no root, the new entity becomes the root.
func (r *Registry) CreateEntity() (entity.Entity, error) {
	// Safely increment the ID counter to get the next entity ID.
	newID := r.idGenerator.Next()
	newEntity := entity.Entity(newID)

	// Add the new entity to the entity map.
	r.mux.Lock()
	defer r.mux.Unlock()
	r.entities[newEntity] = struct{}{}

	return newEntity, nil
}

// DeleteEntity removes an entity and all its components from the registry.
func (r *Registry) DeleteEntity(e entity.Entity) {
	r.mux.Lock()
	defer r.mux.Unlock()
	delete(r.entities, e)
	r.stores.DeleteEntity(e)
}
