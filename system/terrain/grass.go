package terrain

import (
	"github.com/dwethmar/judoka/assets"
	"github.com/dwethmar/judoka/tilebitmasking"
	"github.com/hajimehoshi/ebiten/v2"
)

var GrassShapes = map[int]*ebiten.Image{
	tilebitmasking.AllEdges:                         assets.TerrainImg.SubImage(assets.TerrainCells[1][0]).(*ebiten.Image), //
	tilebitmasking.NoEdges:                          assets.TerrainImg.SubImage(assets.TerrainCells[5][0]).(*ebiten.Image),
	tilebitmasking.EastSoutheastSouthSouthwestWest:  assets.TerrainImg.SubImage(assets.TerrainCells[4][1]).(*ebiten.Image), // x
	tilebitmasking.NorthwestNorthSouthSouthwestWest: assets.TerrainImg.SubImage(assets.TerrainCells[5][1]).(*ebiten.Image),
	tilebitmasking.NorthwestNorthNortheastEastWest:  assets.TerrainImg.SubImage(assets.TerrainCells[7][1]).(*ebiten.Image),
	tilebitmasking.NorthNortheastEastSouthEastSouth: assets.TerrainImg.SubImage(assets.TerrainCells[6][1]).(*ebiten.Image),
	tilebitmasking.EastSoutheastSouthEdges:          assets.TerrainImg.SubImage(assets.TerrainCells[1][1]).(*ebiten.Image), // X
	tilebitmasking.SouthSouthwestWestEdges:          assets.TerrainImg.SubImage(assets.TerrainCells[0][1]).(*ebiten.Image),
	tilebitmasking.NorthwestNorthWestEdges:          assets.TerrainImg.SubImage(assets.TerrainCells[2][1]).(*ebiten.Image),
	tilebitmasking.NorthNortheastEastEdges:          assets.TerrainImg.SubImage(assets.TerrainCells[3][1]).(*ebiten.Image),
	tilebitmasking.HorizontalAndVerticalEdges:       assets.TerrainImg.SubImage(assets.TerrainCells[2][0]).(*ebiten.Image), //
	tilebitmasking.SouthEdge:                        assets.TerrainImg.SubImage(assets.TerrainCells[0][2]).(*ebiten.Image), //
	tilebitmasking.WestEdge:                         assets.TerrainImg.SubImage(assets.TerrainCells[1][2]).(*ebiten.Image),
	tilebitmasking.NorthEdge:                        assets.TerrainImg.SubImage(assets.TerrainCells[3][2]).(*ebiten.Image),
	tilebitmasking.EastEdge:                         assets.TerrainImg.SubImage(assets.TerrainCells[2][2]).(*ebiten.Image),
	tilebitmasking.HorizontalEdges:                  assets.TerrainImg.SubImage(assets.TerrainCells[3][0]).(*ebiten.Image),
	tilebitmasking.VerticalEdges:                    assets.TerrainImg.SubImage(assets.TerrainCells[4][0]).(*ebiten.Image),

	tilebitmasking.EastSouthWestEdges:  assets.TerrainImg.SubImage(assets.TerrainCells[7][2]).(*ebiten.Image),
	tilebitmasking.NorthEastWestEdges:  assets.TerrainImg.SubImage(assets.TerrainCells[4][2]).(*ebiten.Image),
	tilebitmasking.NorthSouthWestEdges: assets.TerrainImg.SubImage(assets.TerrainCells[6][2]).(*ebiten.Image),
	tilebitmasking.NorthEastSouthEdges: assets.TerrainImg.SubImage(assets.TerrainCells[5][2]).(*ebiten.Image),

	tilebitmasking.NorthwestNorthEastSoutheastSouthSouthWestWestEdges: assets.TerrainImg.SubImage(assets.TerrainCells[0][3]).(*ebiten.Image),
	tilebitmasking.NorthNortheastEastSoutheastSouthSouthWestWestEdges: assets.TerrainImg.SubImage(assets.TerrainCells[1][3]).(*ebiten.Image),
	tilebitmasking.NorthwestNorthNorthEastEastSouthSouthwestWestEdges: assets.TerrainImg.SubImage(assets.TerrainCells[2][3]).(*ebiten.Image),
	tilebitmasking.NorthwestNorthNortheastEastSoutheastSouthWestEdges: assets.TerrainImg.SubImage(assets.TerrainCells[3][3]).(*ebiten.Image),

	tilebitmasking.EastSoutheastSouthWestEdges:  assets.TerrainImg.SubImage(assets.TerrainCells[8][0]).(*ebiten.Image),
	tilebitmasking.NorthSouthSouthwestEastEdges: assets.TerrainImg.SubImage(assets.TerrainCells[9][0]).(*ebiten.Image),
	tilebitmasking.NorthEastSoutheastSouthEdges: assets.TerrainImg.SubImage(assets.TerrainCells[10][0]).(*ebiten.Image),
	tilebitmasking.NorthNortheastEastWestEdges:  assets.TerrainImg.SubImage(assets.TerrainCells[11][0]).(*ebiten.Image),
	tilebitmasking.EastSouthSouthwestWestEdges:  assets.TerrainImg.SubImage(assets.TerrainCells[8][1]).(*ebiten.Image),
	tilebitmasking.NorthwestNorthSouthWestEdges: assets.TerrainImg.SubImage(assets.TerrainCells[9][1]).(*ebiten.Image),
	tilebitmasking.NorthNorthEastEastSouthEdges: assets.TerrainImg.SubImage(assets.TerrainCells[10][1]).(*ebiten.Image),
	tilebitmasking.NorthwestNorthEastWestEdges:  assets.TerrainImg.SubImage(assets.TerrainCells[11][1]).(*ebiten.Image),

	tilebitmasking.NorthEastSouthEastSouthSouthwestWestEdges: assets.TerrainImg.SubImage(assets.TerrainCells[8][2]).(*ebiten.Image),
	tilebitmasking.NorthwestNorthEastSouthSouthwestWestEdges: assets.TerrainImg.SubImage(assets.TerrainCells[9][2]).(*ebiten.Image),
	tilebitmasking.NorthNortheastEastSoutheastSouthWestEdges: assets.TerrainImg.SubImage(assets.TerrainCells[10][2]).(*ebiten.Image),
	tilebitmasking.NorthwestNorthNortheastEastSouthWestEdges: assets.TerrainImg.SubImage(assets.TerrainCells[11][2]).(*ebiten.Image),
}
