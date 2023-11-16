package assets

import (
	"image"

	_ "embed"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	skeletonKillHeight = 64
	skeletonKillWidth  = 64
)

var (
	SkeletonKillImg *ebiten.Image
)

var (
	SkeletonKill1 *ebiten.Image
	SkeletonKill2 *ebiten.Image
	SkeletonKill3 *ebiten.Image
	SkeletonKill4 *ebiten.Image
	SkeletonKill5 *ebiten.Image
	SkeletonKill6 *ebiten.Image
)

// animations
var (
	SkeletonKillFrames = []image.Image{}
)

func init() {
	img, err := loadPng(staticSpritesFS, "sprites/skeleton_kill.png")
	if err != nil {
		panic(err)
	}

	SkeletonKillImg = ebiten.NewImageFromImage(img)
	cells := CreateCells(6, 1, skeletonKillWidth, skeletonKillHeight)

	SkeletonKill1 = SkeletonKillImg.SubImage(cells[0][0]).(*ebiten.Image)
	SkeletonKill2 = SkeletonKillImg.SubImage(cells[1][0]).(*ebiten.Image)
	SkeletonKill3 = SkeletonKillImg.SubImage(cells[2][0]).(*ebiten.Image)
	SkeletonKill4 = SkeletonKillImg.SubImage(cells[3][0]).(*ebiten.Image)
	SkeletonKill5 = SkeletonKillImg.SubImage(cells[4][0]).(*ebiten.Image)
	SkeletonKill6 = SkeletonKillImg.SubImage(cells[5][0]).(*ebiten.Image)

	SkeletonKillFrames = []image.Image{
		SkeletonKill1,
		SkeletonKill2,
		SkeletonKill3,
		SkeletonKill4,
		SkeletonKill5,
		SkeletonKill6,
	}
}
