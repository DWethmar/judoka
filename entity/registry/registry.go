package registry

import (
	"errors"
	"sync"

	"github.com/dwethmar/judoka/component"
	"github.com/dwethmar/judoka/entity"
	"github.com/dwethmar/judoka/entity/hierarchy"
	"github.com/dwethmar/judoka/ids"
)

var (
	ErrUniqueConstraintFailed = errors.New("component already exists")
)

type Stores struct {
	Transform  *Store[*component.Transform]  // 1
	Sprite     *Store[*component.Sprite]     // 2
	Controller *Store[*component.Controller] // 3
	Velocity   *Store[*component.Velocity]   // 4
	Actor      *Store[*component.Actor]      // 5
	Chunk      *Store[*component.Chunk]      // 6
}

func (s *Stores) RemoveFromStores(e entity.Entity) {
	s.Transform.RemoveAll(e)  // 1
	s.Sprite.RemoveAll(e)     // 2
	s.Controller.RemoveAll(e) // 3
	s.Velocity.RemoveAll(e)   // 4
	s.Actor.RemoveAll(e)      // 5
	s.Chunk.RemoveAll(e)      // 6
}

// Registry keeps track of all entities and their components.
// It provides methods for entity creation and component management.
type Registry struct {
	Stores
	idGenerator *ids.Generator
	mux         sync.RWMutex
	hierarchy   *hierarchy.Hierarchy
}

// New creates and returns a new instance of Registry.
// It initializes the entity map and the component stores.
func New() *Registry {
	idGen := ids.New(0)
	return &Registry{
		Stores:      stores,
		idGenerator: idGen,
		mux:         sync.RWMutex{},
		hierarchy:   hierarchy.New(entity.Entity(0)),
	}
}

// CreateEntity generates a new unique Entity ID, registers it as the child of the given parent, and returns it.
// If no parent is specified, and there is no root, the new entity becomes the root.
// A transform component is automatically created for the new entity.
func (r *Registry) Create(parent entity.Entity) (entity.Entity, error) {
	// check if parent exists
	if _, ok := r.hierarchy.Get(parent); !ok {
		return entity.Entity(0), errors.New("parent does not exist")
	}

	// Safely increment the ID counter to get the next entity ID.
	newID := r.idGenerator.Next()
	newEntity := entity.Entity(newID)

	// Add the new entity to the entity map.
	r.mux.Lock()
	defer r.mux.Unlock()
	r.hierarchy.AddChild(parent, newEntity)

	// create transform component
	if err := r.Transform.Add(component.NewTransform(newID, newEntity, 0, 0)); err != nil {
		return entity.Entity(0), err
	}

	return newEntity, nil
}

// DeleteEntity removes an entity and all its components from the registry.
func (r *Registry) Delete(e entity.Entity) {
	r.mux.Lock()
	defer r.mux.Unlock()

	// get all children and remove their components
	for _, child := range hierarchy.Walk(r.hierarchy, e) {
		r.Stores.RemoveFromStores(child)
	}

	r.hierarchy.Remove(e)
	r.Stores.RemoveFromStores(e)
}

// Root returns the root entity of the hierarchy.
func (r *Registry) Root() entity.Entity {
	return r.hierarchy.Root().Entity
}

// Parent returns the parent of the given entity.
func (r *Registry) Parent(e entity.Entity) (entity.Entity, bool) {
	if n, ok := r.hierarchy.Get(e); ok {
		return n.Parent.Entity, true
	}

	return entity.Entity(0), false
}

// Children returns the children of the given entity.
func (r *Registry) Children(e entity.Entity) []entity.Entity {
	var result []entity.Entity
	if n, ok := r.hierarchy.Get(e); ok {
		for _, child := range n.Children {
			result = append(result, child.Entity)
		}
	}

	return result
}
