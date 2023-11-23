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
}
