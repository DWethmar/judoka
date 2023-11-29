package terrain

import (
	"github.com/dwethmar/judoka/assets"
	"github.com/dwethmar/judoka/level"
	"github.com/dwethmar/judoka/tilebitmasking"
	"github.com/hajimehoshi/ebiten/v2"
)

// https://code.tutsplus.com/how-to-use-tile-bitmasking-to-auto-tile-your-level-layouts--cms-25673t

var NotFound = assets.TerrainImg.SubImage(assets.TerrainCells[0][0]).(*ebiten.Image)

const (
	Water     int = 0
	GrassTile int = 1
	DirtTile  int = 2
)

func getNeighbors(x, y int, l *level.Level) *tilebitmasking.Neighborhood {
	return &tilebitmasking.Neighborhood{
		NorthWest: l.Tile(x-1, y-1, -1),
		North:     l.Tile(x, y-1, -1),
		NorthEast: l.Tile(x+1, y-1, -1),
		West:      l.Tile(x-1, y, -1),
		Center:    l.Tile(x, y, -1),
		East:      l.Tile(x+1, y, -1),
		SouthWest: l.Tile(x-1, y+1, -1),
		South:     l.Tile(x, y+1, -1),
		SouthEast: l.Tile(x+1, y+1, -1),
	}
}

func Neighbors(x, y int, l *level.Level) *tilebitmasking.Neighborhood {
	return getNeighbors(x, y, l)
}

func TileImage(currentTile int, bitmask int) *ebiten.Image {
	switch currentTile {
	case GrassTile:
		s, ok := GrassShapes[bitmask]
		if ok {
			return s
		}
	case DirtTile:
		return assets.OldTerrainDirtCC
	case Water:
		return assets.WaterImg.SubImage(assets.WaterCells[0][0]).(*ebiten.Image)
	default:
		// Default case
	}

	return NotFound
}
