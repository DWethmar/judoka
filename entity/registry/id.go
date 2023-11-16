package registry

import (
	"sync/atomic"
)

type IDGenerator struct {
	nextID uint32
}

func NewIDGenerator(start uint32) *IDGenerator {
	return &IDGenerator{
		nextID: start,
	}
}

func (g *IDGenerator) Next() uint32 {
	atomic.AddUint32(&g.nextID, 1)
	return g.nextID
}
