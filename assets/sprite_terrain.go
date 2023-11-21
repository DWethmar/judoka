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
	TerrainGrassCC                  *ebiten.Image
	TerrainGrassDirtBottomRightEdge *ebiten.Image
	TerrainGrassDirtBottomEdge      *ebiten.Image
	TerrainGrassDirtBottomLeftEdge  *ebiten.Image
	TerrainGrassDirtRightEdge       *ebiten.Image
	TerrainGrassDirtLeftEdge        *ebiten.Image
	TerrainGrassDirtTopRightEdge    *ebiten.Image
	TerrainGrassDirtTopEdge         *ebiten.Image
	TerrainGrassDirtTopLeftEdge     *ebiten.Image

	TerrainDirtCC               *ebiten.Image
	TerrainDirtGrassBottomRight *ebiten.Image
	TerrainDirtGrassBottomLeft  *ebiten.Image
	TerrainDirtGrassTopRight    *ebiten.Image
	TerrainDirtGrassTopLeft     *ebiten.Image
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
	TerrainGrassDirtBottomRightEdge = terrainImg.SubImage(cells[0][1]).(*ebiten.Image)
	TerrainGrassDirtBottomEdge = terrainImg.SubImage(cells[1][1]).(*ebiten.Image)
	TerrainGrassDirtBottomLeftEdge = terrainImg.SubImage(cells[2][1]).(*ebiten.Image)
	TerrainGrassDirtRightEdge = terrainImg.SubImage(cells[0][2]).(*ebiten.Image)
	TerrainGrassDirtLeftEdge = terrainImg.SubImage(cells[2][2]).(*ebiten.Image)
	TerrainGrassDirtTopRightEdge = terrainImg.SubImage(cells[0][3]).(*ebiten.Image)
	TerrainGrassDirtTopEdge = terrainImg.SubImage(cells[1][3]).(*ebiten.Image)
	TerrainGrassDirtTopLeftEdge = terrainImg.SubImage(cells[2][3]).(*ebiten.Image)

	// dirt
	TerrainDirtCC = terrainImg.SubImage(cells[1][2]).(*ebiten.Image)

	// Dirt Grass Corners 2 X 2
	TerrainDirtGrassBottomRight = terrainImg.SubImage(cells[3][1]).(*ebiten.Image)
	TerrainDirtGrassBottomLeft = terrainImg.SubImage(cells[4][1]).(*ebiten.Image)
	TerrainDirtGrassTopRight = terrainImg.SubImage(cells[3][2]).(*ebiten.Image)
	TerrainDirtGrassTopLeft = terrainImg.SubImage(cells[4][2]).(*ebiten.Image)
}
