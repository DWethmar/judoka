package assets

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	WaterTileRows    = 1
	WaterTileColumns = 2
	WaterTileWidth   = 32
	WaterTileHeight  = 32
)

var (
	WaterImg   *ebiten.Image
	WaterCells = CreateCells(WaterTileColumns, WaterTileRows, TerrainTileWidth, TerrainTileHeight)
)

func init() {
	img, err := loadPng(staticSpritesFS, "sprites/water_tiles.png")
	if err != nil {
		panic(err)
	}

	WaterImg = ebiten.NewImageFromImage(img)
}
