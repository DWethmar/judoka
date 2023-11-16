package registry

import (
	"sync"
	"sync/atomic"
)

type IDGenerator struct {
	nextID uint32
	mux    sync.RWMutex
	ids    map[uint32]struct{}
}

func NewIDGenerator() *IDGenerator {
	return &IDGenerator{
		ids:    make(map[uint32]struct{}),
		nextID: 0,
	}
}

func (g *IDGenerator) Next() uint32 {
	atomic.AddUint32(&g.nextID, 1)
	g.mux.Lock()
	defer g.mux.Unlock()
	g.ids[g.nextID] = struct{}{}
	return g.nextID
}
