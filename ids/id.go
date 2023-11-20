package ids

import (
	"sync/atomic"
)

type Generator struct {
	nextID uint32
}

func New(start uint32) *Generator {
	return &Generator{
		nextID: start,
	}
}

func (g *Generator) Next() uint32 {
	atomic.AddUint32(&g.nextID, 1)
	return g.nextID
}
