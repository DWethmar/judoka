package hierarchy

import "github.com/dwethmar/judoka/entity"

// Walk collects all child entities starting from a given entity.
func Walk(h *Hierarchy, start entity.Entity) []entity.Entity {
	var result []entity.Entity
	walkHelper(h, start, &result)
	return result
}

// walkHelper is a recursive helper function for walking the tree.
func walkHelper(h *Hierarchy, current entity.Entity, result *[]entity.Entity) {
	node, ok := h.Get(current)
	if !ok {
		return // If the current node doesn't exist, return
	}

	for _, child := range node.Children {
		*result = append(*result, child.Entity)
		walkHelper(h, child.Entity, result) // Recursive call for each child
	}
}
