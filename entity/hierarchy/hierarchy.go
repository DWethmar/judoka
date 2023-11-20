package hierarchy

import (
	"fmt"
	"sync"

	"github.com/dwethmar/judoka/entity"
)

// Node represents a single entity in the hierarchy.
type Node struct {
	Entity   entity.Entity
	Children []*Node
	Parent   *Node
}

// Hierarchy represents the entire hierarchy of entities.
type Hierarchy struct {
	mux   sync.RWMutex
	nodes map[entity.Entity]*Node
	root  *Node // The root node of the hierarchy.
}

// New creates a new, empty Hierarchy.
func New(root entity.Entity) *Hierarchy {
	rootNode := &Node{
		Entity: root,
	}
	return &Hierarchy{
		root: rootNode,
		nodes: map[entity.Entity]*Node{
			root: rootNode,
		},
	}
}

// Get retrieves the Node associated with the given entity.
// It returns the found Node and a boolean indicating whether the Node was found.
func (h *Hierarchy) Get(e entity.Entity) (*Node, bool) {
	h.mux.RLock()
	defer h.mux.RUnlock()
	node, ok := h.nodes[e]
	return node, ok
}

// Root returns the root entity of the hierarchy.
func (h *Hierarchy) Root() *Node {
	h.mux.RLock()
	defer h.mux.RUnlock()
	return h.root
}

// AddChild adds a child entity to a parent entity in the hierarchy.
func (h *Hierarchy) AddChild(parentEntity, childEntity entity.Entity) error {
	// Check if the child entity already exists in the hierarchy.
	if _, exists := h.Get(childEntity); exists {
		return fmt.Errorf("child entity %d already exists in the hierarchy", childEntity)
	}

	// Retrieve the parent node.
	parentNode, exists := h.Get(parentEntity)
	if !exists {
		return fmt.Errorf("parent entity %d does not exist in the hierarchy", parentEntity)
	}

	h.mux.Lock()
	defer h.mux.Unlock()

	// Create a new node for the child entity.
	childNode := &Node{
		Entity: childEntity,
		Parent: parentNode,
	}

	// Add the child node to the parent's children and to the hierarchy's nodes.
	parentNode.Children = append(parentNode.Children, childNode)
	h.nodes[childEntity] = childNode
	return nil
}

// Move changes the parent of a node to a new parent within the hierarchy.
func (h *Hierarchy) Move(entity entity.Entity, newParent entity.Entity) error {
	h.mux.Lock()
	defer h.mux.Unlock()
	// Cannot move the root node
	if h.root != nil && h.root.Entity == entity {
		return fmt.Errorf("cannot move the root node")
	}

	node, ok := h.nodes[entity]
	if !ok {
		return fmt.Errorf("entity to move does not exist")
	}

	// If the new parent is the same as the current parent, nothing to do
	if node.Parent != nil && node.Parent.Entity == newParent {
		return nil
	}

	// Remove node from current parent's children list, if it has one
	if node.Parent != nil {
		h.removeChildFromParent(node)
	}

	// Set the new parent for the node
	newParentNode, ok := h.nodes[newParent]
	if !ok {
		return fmt.Errorf("new parent does not exist")
	}

	// Update node's parent
	node.Parent = newParentNode
	// Add node to new parent's children list
	newParentNode.Children = append(newParentNode.Children, node)

	return nil
}

// Remove deletes an entity from the hierarchy.
// If the entity has children, it recursively removes them as well.
func (h *Hierarchy) Remove(e entity.Entity) error {
	h.mux.Lock()
	defer h.mux.Unlock()
	node, ok := h.nodes[e]
	if !ok {
		// Entity not found in the hierarchy
		return fmt.Errorf("entity %d not found in the hierarchy", e)
	}

	// If the node is the root, we remove the entire hierarchy.
	if h.root == node {
		h.root = nil
		h.nodes = make(map[entity.Entity]*Node)
		return nil
	}

	// Recursively remove all children
	h.removeChildren(node)

	// Remove the node from its parent's children slice
	h.removeChildFromParent(node)

	// Finally, remove the node from the nodes map
	delete(h.nodes, e)
	return nil
}

// removeChildren recursively removes all children of a given node.
func (h *Hierarchy) removeChildren(node *Node) {
	for _, childNode := range node.Children {
		// Recursively remove all descendants of the child
		h.removeChildren(childNode)
		// Remove the child from the nodes map
		delete(h.nodes, childNode.Entity)
	}
	// Clear the children slice to help garbage collection
	node.Children = nil
}

// removeChildFromParent removes a node from its parent's children slice.
func (h *Hierarchy) removeChildFromParent(node *Node) {
	parentNode := node.Parent
	if parentNode == nil {
		return // If the node has no parent, nothing to do
	}
	for i, child := range parentNode.Children {
		if child == node {
			// Remove the node from the slice
			parentNode.Children = append(parentNode.Children[:i], parentNode.Children[i+1:]...)
			break
		}
	}
}
