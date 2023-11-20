package terrain

import (
	"github.com/dwethmar/judoka/assets"
	"github.com/dwethmar/judoka/level"
	"github.com/hajimehoshi/ebiten/v2"
)

type Tile int

const (
	GrassTile Tile = 0
	DirtTile  Tile = 1
)

var Tiles = map[Tile]*ebiten.Image{
	GrassTile: assets.TerrainGrassCC,
	DirtTile:  assets.TerrainDirtCC,
}

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

func getNeighbors(x, y int, l *level.Level) TileNeighbors {
	return TileNeighbors{
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

func Shapes(x, y int, l *level.Level) *ebiten.Image {
	currentTile := Tile(l.GetTile(x, y))
	neighbors := getNeighbors(x, y, l)

	switch currentTile {
	case GrassTile:
		// Corners
		if GrassToDirtTopLeftCorner(neighbors) {
			return assets.TerrainDirtGrassBottomLeft
		} else if GrassToDirtTopRightCorner(neighbors) {
			return assets.TerrainDirtGrassBottomRigtht
		} else if GrassToDirtBottomLeftCorner(neighbors) {
			return assets.TerrainDirtGrassTopRight
		} else if GrassToDirtBottomRightCorner(neighbors) {
			return assets.TerrainDirtGrassTopRight
		}

		// Transitions
		if GrassToDirtBottomRight(neighbors) {
			return assets.TerrainGrassDirtBottomRight
		} else if GrassToDirtTopLeft(neighbors) {
			return assets.TerrainGrassDirtTopLeft
		} else if GrassToDirtTopRight(neighbors) {
			return assets.TerrainGrassDirtTopRight
		} else if GrassToDirtBottomLeft(neighbors) {
			return assets.TerrainGrassDirtBottomLeft
		}

		// Sides
		if GrassToDirtLeftSide(neighbors) {
			return assets.TerrainGrassDirtLeft
		}

		if GrassToDirtRightSide(neighbors) {
			return assets.TerrainGrassDirtRight
		}

		if GrassToDirtTopSide(neighbors) {
			return assets.TerrainGrassDirtTop
		}

		if GrassToDirtBottomSide(neighbors) {
			return assets.TerrainGrassDirtBottom
		}

		return assets.TerrainGrassCC
	case DirtTile:
		return assets.TerrainDirtCC
	default:
		// Return a default image or log an error
	}

	return nil
}
