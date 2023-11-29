package terrain

import (
	"github.com/dwethmar/judoka/assets"
	"github.com/hajimehoshi/ebiten/v2"
)

var (
	WaterBackground1 *ebiten.Image
	WaterBackground2 *ebiten.Image
)

func init() {
	// draw water background
	WaterBackground1 = ebiten.NewImage(ChunkSize*TileSize, ChunkSize*TileSize)
	for i := 0; i < ChunkSize; i++ { // x
		x := i * TileSize
		for j := 0; j < ChunkSize; j++ { // y
			y := j * TileSize

			image := assets.WaterImg.SubImage(assets.WaterCells[0][0]).(*ebiten.Image)
			{
				w := image.Bounds().Dx()
				h := image.Bounds().Dy()

				op := &ebiten.DrawImageOptions{}
				op.GeoM.Scale(float64(TileSize)/float64(w), float64(TileSize)/float64(h))
				op.GeoM.Translate(float64(x), float64(y))
				WaterBackground1.DrawImage(image, op)
			}
		}
	}

	WaterBackground2 = ebiten.NewImage(ChunkSize*TileSize, ChunkSize*TileSize)
	for i := 0; i < ChunkSize; i++ { // x
		x := i * TileSize
		for j := 0; j < ChunkSize; j++ { // y
			y := j * TileSize

			image := assets.WaterImg.SubImage(assets.WaterCells[1][0]).(*ebiten.Image)
			{
				w := image.Bounds().Dx()
				h := image.Bounds().Dy()

				op := &ebiten.DrawImageOptions{}
				op.GeoM.Scale(float64(TileSize)/float64(w), float64(TileSize)/float64(h))
				op.GeoM.Translate(float64(x), float64(y))
				WaterBackground2.DrawImage(image, op)
			}
		}
	}

}
