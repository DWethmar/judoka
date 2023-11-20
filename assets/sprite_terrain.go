package assets

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	TerrainHeight = 320 / 10 // 32
	TerrainWidth  = 224 / 7  // 32
)

var (
	terrainImg *ebiten.Image
)

var (
	TerrainGrassCC              *ebiten.Image
	TerrainGrassDirtBottomRight *ebiten.Image
	TerrainGrassDirtBottom      *ebiten.Image
	TerrainGrassDirtBottomLeft  *ebiten.Image
	TerrainGrassDirtLeft        *ebiten.Image
	TerrainGrassDirtRight       *ebiten.Image
	TerrainGrassDirtTopRight    *ebiten.Image
	TerrainGrassDirtTop         *ebiten.Image
	TerrainGrassDirtTopLeft     *ebiten.Image

	TerrainDirtCC                *ebiten.Image
	TerrainDirtGrassBottomRigtht *ebiten.Image
	TerrainDirtGrassBottomLeft   *ebiten.Image
	TerrainDirtGrassTopRight     *ebiten.Image
	TerrainDirtGrassTopLeft      *ebiten.Image
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

	// Grass Dirt Corners 3 X 3
	TerrainGrassDirtBottomRight = terrainImg.SubImage(cells[0][1]).(*ebiten.Image)
	TerrainGrassDirtBottom = terrainImg.SubImage(cells[1][1]).(*ebiten.Image)
	TerrainGrassDirtBottomLeft = terrainImg.SubImage(cells[2][1]).(*ebiten.Image)
	TerrainGrassDirtLeft = terrainImg.SubImage(cells[0][2]).(*ebiten.Image)
	TerrainGrassDirtRight = terrainImg.SubImage(cells[2][2]).(*ebiten.Image)
	TerrainGrassDirtTopRight = terrainImg.SubImage(cells[0][3]).(*ebiten.Image)
	TerrainGrassDirtTop = terrainImg.SubImage(cells[1][3]).(*ebiten.Image)
	TerrainGrassDirtTopLeft = terrainImg.SubImage(cells[2][3]).(*ebiten.Image)

	// dirt
	TerrainDirtCC = terrainImg.SubImage(cells[1][2]).(*ebiten.Image)

	// Dirt Grass Corners 2 X 2
	TerrainDirtGrassBottomRigtht = terrainImg.SubImage(cells[3][1]).(*ebiten.Image)
	TerrainDirtGrassBottomLeft = terrainImg.SubImage(cells[4][1]).(*ebiten.Image)
	TerrainDirtGrassTopRight = terrainImg.SubImage(cells[3][2]).(*ebiten.Image)
	TerrainDirtGrassTopLeft = terrainImg.SubImage(cells[4][2]).(*ebiten.Image)
}
