package terrain

import (
	"fmt"

	"github.com/dwethmar/judoka/assets"
	"github.com/dwethmar/judoka/level"
	"github.com/dwethmar/judoka/tilebitmasking"
	"github.com/hajimehoshi/ebiten/v2"
)

// https://code.tutsplus.com/how-to-use-tile-bitmasking-to-auto-tile-your-level-layouts--cms-25673t

var NotFound = assets.TerrainImg.SubImage(assets.TerrainCells[0][0]).(*ebiten.Image)

const (
	DirtTile  int = 0
	GrassTile int = 1
)

func getNeighbors(x, y int, l *level.Level) *tilebitmasking.Neighborhood {
	return &tilebitmasking.Neighborhood{
		NorthWest: l.GetTile(x-1, y-1),
		North:     l.GetTile(x, y-1),
		NorthEast: l.GetTile(x+1, y-1),
		West:      l.GetTile(x-1, y),
		Center:    l.GetTile(x, y),
		East:      l.GetTile(x+1, y),
		SouthWest: l.GetTile(x-1, y+1),
		South:     l.GetTile(x, y+1),
		SouthEast: l.GetTile(x+1, y+1),
	}
}

func getTileImage(currentTile int, bitmask int) *ebiten.Image {
	switch currentTile {
	case GrassTile:
		s, ok := GrassShapes[bitmask]
		if ok {
			return s
		}
	case DirtTile:
		return assets.OldTerrainDirtCC
	default:
		// Default case
	}

	return NotFound
}

func Shapes(x, y int, l *level.Level) *ebiten.Image {
	if x == 1 && y == 8 {
		fmt.Println("debugger")
	}
	neighbors := getNeighbors(x, y, l)
	bitmask := tilebitmasking.Calculate(neighbors)

	// log bitmask
	fmt.Printf("x: %d, y: %d,  bitmask: %08b\n", x, y, bitmask)

	return getTileImage(neighbors.Center, bitmask)
}
