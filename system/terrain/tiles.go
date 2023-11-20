package terrain

import (
	"github.com/dwethmar/judoka/assets"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	GrassTile = 0
	DirtTile  = 1
)

var Tiles = map[int32]*ebiten.Image{
	GrassTile: assets.TerrainGrassCC,
	DirtTile:  assets.TerrainDirtCC,
}
