package registry

import (
	"errors"
	"fmt"
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
	Layer      *Store[*component.Layer]      // 7
}

func (s *Stores) RemoveFromStores(e entity.Entity) {
	s.Transform.RemoveAll(e)  // 1
	s.Sprite.RemoveAll(e)     // 2
	s.Controller.RemoveAll(e) // 3
	s.Velocity.RemoveAll(e)   // 4
	s.Actor.RemoveAll(e)      // 5
	s.Chunk.RemoveAll(e)      // 6
	s.Layer.RemoveAll(e)      // 7
}

// Register keeps track of all entities and their components.
// It provides methods for entity creation and component management.
type Register struct {
	Stores
	idGenerator *ids.Generator
	mux         sync.RWMutex
	hierarchy   *hierarchy.Hierarchy
}

// New creates and returns a new instance of Registry.
// It initializes the entity map and the component stores.
func New() (*Register, error) {
	idGen := ids.New(0)
	root := entity.Entity(idGen.Next())
	r := &Register{
		Stores:      stores,
		idGenerator: idGen,
		mux:         sync.RWMutex{},
		hierarchy:   hierarchy.New(root),
	}

	// create transform component
	if err := r.Transform.Add(component.NewTransform(0, root, 0, 0)); err != nil {
		return nil, fmt.Errorf("could not create transform component for root: %w", err)
	}

	return r, nil
}

// List returns a list of all entities in the registry. Sorted by its hierarchy.
func (r *Register) List() []entity.Entity {
	return hierarchy.Walk(r.hierarchy, r.Root())
}

// CreateEntity generates a new unique Entity ID, registers it as the child of the given parent, and returns it.
// If no parent is specified, and there is no root, the new entity becomes the root.
// A transform component is automatically created for the new entity.
func (r *Register) Create(parent entity.Entity) (entity.Entity, error) {
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
func (r *Register) Delete(e entity.Entity) {
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
func (r *Register) Root() entity.Entity {
	return r.hierarchy.Root().Entity
}

// Parent returns the parent of the given entity.
func (r *Register) Parent(e entity.Entity) (entity.Entity, bool) {
	if n, ok := r.hierarchy.Get(e); ok {
		if n.Parent == nil {
			return entity.Entity(0), false
		}
		return n.Parent.Entity, true
	}

	return entity.Entity(0), false
}

// Children returns the children of the given entity.
func (r *Register) Children(e entity.Entity) []entity.Entity {
	var result []entity.Entity
	if n, ok := r.hierarchy.Get(e); ok {
		for _, child := range n.Children {
			result = append(result, child.Entity)
		}
	}

	return result
}
