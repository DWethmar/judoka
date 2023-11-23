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
	OldTerrainGrassCC                  *ebiten.Image
	OldTerrainGrassDirtBottomRightEdge *ebiten.Image
	OldTerrainGrassDirtBottomEdge      *ebiten.Image
	OldTerrainGrassDirtBottomLeftEdge  *ebiten.Image
	OldTerrainGrassDirtRightEdge       *ebiten.Image
	OldTerrainGrassDirtLeftEdge        *ebiten.Image
	OldTerrainGrassDirtTopRightEdge    *ebiten.Image
	OldTerrainGrassDirtTopEdge         *ebiten.Image
	OldTerrainGrassDirtTopLeftEdge     *ebiten.Image

	OldTerrainDirtCC               *ebiten.Image
	OldTerrainDirtGrassBottomRight *ebiten.Image
	OldTerrainDirtGrassBottomLeft  *ebiten.Image
	OldTerrainDirtGrassTopRight    *ebiten.Image
	OldTerrainDirtGrassTopLeft     *ebiten.Image
)

func init() {
	img, err := loadPng(staticSpritesFS, "sprites/terrain.png")
	if err != nil {
		panic(err)
	}

	terrainImg = ebiten.NewImageFromImage(img)
	cells := CreateCells(10, 7, TerrainWidth, TerrainHeight)

	// grass
	OldTerrainGrassCC = terrainImg.SubImage(cells[0][0]).(*ebiten.Image)

	// Grass Dirt Corners 3 X 3
	OldTerrainGrassDirtBottomRightEdge = terrainImg.SubImage(cells[0][1]).(*ebiten.Image)
	OldTerrainGrassDirtBottomEdge = terrainImg.SubImage(cells[1][1]).(*ebiten.Image)
	OldTerrainGrassDirtBottomLeftEdge = terrainImg.SubImage(cells[2][1]).(*ebiten.Image)
	OldTerrainGrassDirtRightEdge = terrainImg.SubImage(cells[0][2]).(*ebiten.Image)
	OldTerrainGrassDirtLeftEdge = terrainImg.SubImage(cells[2][2]).(*ebiten.Image)
	OldTerrainGrassDirtTopRightEdge = terrainImg.SubImage(cells[0][3]).(*ebiten.Image)
	OldTerrainGrassDirtTopEdge = terrainImg.SubImage(cells[1][3]).(*ebiten.Image)
	OldTerrainGrassDirtTopLeftEdge = terrainImg.SubImage(cells[2][3]).(*ebiten.Image)

	// dirt
	OldTerrainDirtCC = terrainImg.SubImage(cells[1][2]).(*ebiten.Image)

	// Dirt Grass Corners 2 X 2
	OldTerrainDirtGrassBottomRight = terrainImg.SubImage(cells[3][1]).(*ebiten.Image)
	OldTerrainDirtGrassBottomLeft = terrainImg.SubImage(cells[4][1]).(*ebiten.Image)
	OldTerrainDirtGrassTopRight = terrainImg.SubImage(cells[3][2]).(*ebiten.Image)
	OldTerrainDirtGrassTopLeft = terrainImg.SubImage(cells[4][2]).(*ebiten.Image)
}
