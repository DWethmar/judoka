package level

import "github.com/dwethmar/judoka/component"

// Level is a collection of chunks. It is used to link chunks together.
type Level struct {
	Chunks    map[int]map[int]*component.Chunk
	ChunkSize int
}

func New(size int) *Level {
	return &Level{
		Chunks:    make(map[int]map[int]*component.Chunk),
		ChunkSize: size,
	}
}

func (l *Level) Chunk(chunkX, chunkY int) *component.Chunk {
	if l.Chunks[chunkX] == nil {
		return nil
	}

	return l.Chunks[chunkX][chunkY]
}

func (l *Level) SetChunk(c *component.Chunk) {
	if l.Chunks[c.X] == nil {
		l.Chunks[c.X] = make(map[int]*component.Chunk)
	}
	l.Chunks[c.X][c.Y] = c
}

func (l *Level) Tile(x, y int) int {
	chunkX, tileX := divMod(x, l.ChunkSize)
	chunkY, tileY := divMod(y, l.ChunkSize)

	c := l.Chunk(chunkX, chunkY)
	if c == nil || c.Tiles == nil {
		return 0
	}

	return c.Tiles.Get(tileX, tileY)
}

func (l *Level) SetTile(x, y, v int) {
	chunkX, tileX := divMod(x, l.ChunkSize)
	chunkY, tileY := divMod(y, l.ChunkSize)

	c := l.Chunk(chunkX, chunkY)
	if c == nil {
		return
	}

	c.Tiles.Set(tileX, tileY, v)
}

// divMod performs integer division and modulo with support for negative numbers.
func divMod(numerator, denominator int) (quotient, remainder int) {
	quotient = numerator / denominator
	remainder = numerator % denominator
	if remainder < 0 {
		remainder += denominator
		quotient--
	}
	return
}
