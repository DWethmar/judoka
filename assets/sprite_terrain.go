package assets

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	TerrainHeight = 32
	TerrainWidth  = 32
)

var (
	terrainImg *ebiten.Image
)

var (
	TerrainGrassCC     *ebiten.Image
	TerrainGrassDirtTL *ebiten.Image
	TerrainGrassDirtTC *ebiten.Image
	TerrainGrassDirtTR *ebiten.Image
	TerrainGrassDirtLC *ebiten.Image
	TerrainGrassDirtRC *ebiten.Image
	TerrainGrassDirtBL *ebiten.Image
	TerrainGrassDirtBC *ebiten.Image
	TerrainGrassDirtBR *ebiten.Image
	TerrainDirtCC      *ebiten.Image
)

func init() {
	img, err := loadPng(staticSpritesFS, "sprites/terrain.png")
	if err != nil {
		panic(err)
	}

	terrainImg = ebiten.NewImageFromImage(img)
	cells := CreateCells(10, 7, TerrainWidth, TerrainHeight)

	// grass
	TerrainGrassCC = terrainImg.SubImage(cells[0][0]).(*ebiten.Image)

	// Grass Dirt
	TerrainGrassDirtTL = terrainImg.SubImage(cells[0][1]).(*ebiten.Image)
	TerrainGrassDirtTC = terrainImg.SubImage(cells[1][1]).(*ebiten.Image)
	TerrainGrassDirtTR = terrainImg.SubImage(cells[2][1]).(*ebiten.Image)
	TerrainGrassDirtLC = terrainImg.SubImage(cells[0][2]).(*ebiten.Image)
	TerrainGrassDirtRC = terrainImg.SubImage(cells[2][2]).(*ebiten.Image)
	TerrainGrassDirtBL = terrainImg.SubImage(cells[0][3]).(*ebiten.Image)
	TerrainGrassDirtBC = terrainImg.SubImage(cells[1][3]).(*ebiten.Image)
	TerrainGrassDirtBR = terrainImg.SubImage(cells[2][3]).(*ebiten.Image)

	// Dirt
	TerrainDirtCC = terrainImg.SubImage(cells[1][2]).(*ebiten.Image)
}
