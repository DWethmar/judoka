package registry

import (
	"fmt"
	"sort"
	"sync"

	"github.com/dwethmar/judoka/component"
	"github.com/dwethmar/judoka/entity"
)

// Store is a generic type that associates entity.Entity identifiers with ComponentList.
// It manages the storage and retrieval of components for each entity.
type Store[T component.Component] struct {
	mu        sync.RWMutex
	store     map[entity.Entity][]T
	BeforeAdd func(c T) error
}

type StoreOption[T component.Component] func(*Store[T])

// New creates a new instance of a Store for the specified component type.
// It initializes the internal map that will hold the entity-component associations.
func NewStore[T component.Component](options ...StoreOption[T]) *Store[T] {
	return &Store[T]{
		store: make(map[entity.Entity][]T),
	}
}

// GetAllEntities returns a sorted slice of all entity identifiers present in the store.
func (ecs *Store[T]) Entities() []entity.Entity {
	ecs.mu.RLock()
	defer ecs.mu.RUnlock()

	var ids []entity.Entity
	for id := range ecs.store {
		ids = append(ids, id)
	}

	// Sort the slice of ids
	sort.Slice(ids, func(i, j int) bool { return ids[i] < ids[j] })

	return ids
}

// Add a new component to the ComponentList associated with an entity.Entity identifier.
// It initializes the component before adding it to the list. If the initialization fails,
// it returns an error.
func (ecs *Store[T]) Add(c T) error {
	ecs.mu.Lock()
	defer ecs.mu.Unlock()

	if ecs.BeforeAdd != nil {
		if err := ecs.BeforeAdd(c); err != nil {
			return fmt.Errorf("before add: %w", err)
		}
	}

	ecs.store[c.Entity()] = append(ecs.store[c.Entity()], c)
	return nil
}

// First returns the first component associated with an entity.Entity identifier,
// along with a boolean indicating whether a valid component was found.
func (ecs *Store[T]) First(e entity.Entity) (T, bool) {
	ecs.mu.RLock()
	defer ecs.mu.RUnlock()

	list, ok := ecs.store[e]
	if !ok || len(list) == 0 {
		var zero T         // Get the zero value of T
		return zero, false // Return zero value and false if no component is found
	}

	return list[0], true // Return the first component and true
}

// List the components associated with an entity.Entity identifier.
// It returns the list along with a boolean indicating whether the list was found.
// If the list is not found, it returns an empty list.
func (ecs *Store[T]) List(e entity.Entity) []T {
	ecs.mu.RLock()
	defer ecs.mu.RUnlock()

	list, ok := ecs.store[e]
	if !ok {
		return []T{}
	}

	return list
}

// Remove removes a single component associated with an entity.Entity identifier.
// If after removal, the ComponentList is empty, it deletes the entry for the entity from the internal map.
func (ecs *Store[T]) Remove(c T) {
	ecs.mu.Lock()
	defer ecs.mu.Unlock()

	list, ok := ecs.store[c.Entity()]
	if !ok {
		return // Entity does not exist, so there's nothing to remove
	}

	// Find the component to remove
	for i, component := range list {
		if component.ID() == c.ID() {
			// Remove the component from the list by filtering it out
			copy(list[i:], list[i+1:])
			var zero T                // Get the zero value of T
			list[len(list)-1] = zero  // Set the last element to the zero value
			list = list[:len(list)-1] // Truncate slice

			// If the list is empty, remove the entry from the map.
			if len(list) == 0 {
				delete(ecs.store, c.Entity())
			} else {
				ecs.store[c.Entity()] = list // Store the updated list
			}

			return // Component found and removed; exit the function
		}
	}
	// If we reach this point, the component was not found; nothing changes
}

// RemoveAll removes all components associated with an entity.Entity identifier.
func (ecs *Store[T]) RemoveAll(e entity.Entity) {
	ecs.mu.Lock()
	defer ecs.mu.Unlock()

	delete(ecs.store, e)
}
