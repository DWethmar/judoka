package assets

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	TerrainTileRows    = 5
	TerrainTileColumns = 12
	TerrainTileHeight  = 32
	TerrainTileWidth   = 32
)

var (
	TerrainImg   *ebiten.Image
	TerrainCells = CreateCells(12, 5, TerrainTileWidth, TerrainTileHeight)
)

func init() {
	img, err := loadPng(staticSpritesFS, "sprites/terrain_tiles2.png")
	if err != nil {
		panic(err)
	}

	TerrainImg = ebiten.NewImageFromImage(img)
}
