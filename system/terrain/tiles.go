package terrain

import (
	"fmt"

	"github.com/dwethmar/judoka/assets"
	"github.com/dwethmar/judoka/level"
	"github.com/hajimehoshi/ebiten/v2"
)

type Tile int

const (
	GrassTile Tile = 0
	DirtTile  Tile = 1
)

var GrassShapes = map[int]*ebiten.Image{
	AllSame:      assets.TerrainGrassCC,
	AllDifferent: assets.TerrainGrassCC,
	// edges
	NorthDifferent: assets.TerrainGrassDirtTopEdge,
	EastDifferent:  assets.TerrainGrassDirtRightEdge,
	SouthDifferent: assets.TerrainGrassDirtBottomEdge,
	WestDifferent:  assets.TerrainGrassDirtLeftEdge,
	// corners
	NorthWestDifferent: assets.TerrainDirtGrassBottomRight,
	NorthEastDifferent: assets.TerrainDirtGrassBottomLeft,
	SouthWestDifferent: assets.TerrainDirtGrassTopRight,
	SouthEastDifferent: assets.TerrainDirtGrassTopLeft,
}

var defaultTiles = map[Tile]*ebiten.Image{
	GrassTile: assets.TerrainGrassCC,
	DirtTile:  assets.TerrainDirtCC,
}

const (
	North = 1 << 3 // 1000
	East  = 1 << 2 // 0100
	South = 1 << 1 // 0010
	West  = 1 << 0 // 0001
)

const (
	// Example constants for the new bitmask system
	AllSame        = 0b0000 // 0000
	AllDifferent   = 0b1111 // 1111
	NorthDifferent = 0b1000 // 1000
	EastDifferent  = 0b0100 // 0100
	SouthDifferent = 0b0010 // 0010
	WestDifferent  = 0b0001 // 0001
	// corners (two full edges)
	NorthWestDifferent = 0b1001
	NorthEastDifferent = 0b1100
	SouthWestDifferent = 0b0011
	SouthEastDifferent = 0b0110
)

type TileNeighbors struct {
	TopLeft     Tile
	Top         Tile
	TopRight    Tile
	Left        Tile
	Right       Tile
	BottomLeft  Tile
	Bottom      Tile
	BottomRight Tile
}

func getNeighbors(x, y int, l *level.Level) *TileNeighbors {
	return &TileNeighbors{
		TopLeft:     Tile(l.GetTile(x-1, y-1)),
		Top:         Tile(l.GetTile(x, y-1)),
		TopRight:    Tile(l.GetTile(x+1, y-1)),
		Left:        Tile(l.GetTile(x-1, y)),
		Right:       Tile(l.GetTile(x+1, y)),
		BottomLeft:  Tile(l.GetTile(x-1, y+1)),
		Bottom:      Tile(l.GetTile(x, y+1)),
		BottomRight: Tile(l.GetTile(x+1, y+1)),
	}
}

func calculateBitmask(neighbors *TileNeighbors, currentTile Tile) int {
	bitmask := 0

	if neighbors.Top != currentTile {
		bitmask |= North
	}
	if neighbors.Right != currentTile {
		bitmask |= East
	}
	if neighbors.Bottom != currentTile {
		bitmask |= South
	}
	if neighbors.Left != currentTile {
		bitmask |= West
	}

	return bitmask
}

func getTileImage(currentTile Tile, bitmask int) *ebiten.Image {
	switch currentTile {
	case GrassTile:
		s, ok := GrassShapes[bitmask]
		if ok {
			return s
		}
		return assets.TerrainGrassCC
	case DirtTile:
		return assets.TerrainDirtCC
	default:
		// Default case
	}

	return nil
}

func Shapes(x, y int, l *level.Level) *ebiten.Image {
	if x == 1 && y == 1 {
		fmt.Println("TopLeftCorner")
	}
	currentTile := Tile(l.GetTile(x, y))
	neighbors := getNeighbors(x, y, l)
	bitmask := calculateBitmask(neighbors, currentTile)
	img := getTileImage(currentTile, bitmask)

	fmt.Printf("bitmask X: %d Y: %d: %08b\n", x, y, bitmask)
	if img != nil {
		return img
	}

	return defaultTiles[currentTile]
}
