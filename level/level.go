package level

import "github.com/dwethmar/judoka/matrix"

type Level struct {
	Chunks    map[int]map[int]matrix.Matrix
	ChunkSize int
}

func New(size int) *Level {
	return &Level{
		Chunks:    make(map[int]map[int]matrix.Matrix),
		ChunkSize: size,
	}
}

func (l *Level) GetChunk(chunkX, chunkY int) matrix.Matrix {
	if l.Chunks[chunkX] == nil {
		return nil
	}

	return l.Chunks[chunkX][chunkY]
}

func (l *Level) SetChunk(chunkX, chunkY int, m matrix.Matrix) {
	if l.Chunks[chunkX] == nil {
		l.Chunks[chunkX] = make(map[int]matrix.Matrix)
	}
	l.Chunks[chunkX][chunkY] = m
}

func (l *Level) GetTile(x, y int) int {
	chunkX, tileX := divMod(x, l.ChunkSize)
	chunkY, tileY := divMod(y, l.ChunkSize)

	c := l.GetChunk(chunkX, chunkY)
	if c == nil {
		return 0
	}

	return c.Get(tileX, tileY)
}

func (l *Level) SetTile(x, y, v int) {
	chunkX, tileX := divMod(x, l.ChunkSize)
	chunkY, tileY := divMod(y, l.ChunkSize)

	c := l.GetChunk(chunkX, chunkY)
	if c == nil {
		return
	}

	c.Set(tileX, tileY, v)
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
