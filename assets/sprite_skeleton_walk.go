package assets

import (
	"image"

	_ "embed"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	skeletonWalkHeight = 64
	skeletonWalkWidth  = 64
)

var (
	SkeletonWalkImg *ebiten.Image
)

var (
	SkeletonUp1 *ebiten.Image // idle
	SkeletonUp2 *ebiten.Image
	SkeletonUp3 *ebiten.Image
	SkeletonUp4 *ebiten.Image
	SkeletonUp5 *ebiten.Image
	SkeletonUp6 *ebiten.Image
	SkeletonUp7 *ebiten.Image
	SkeletonUp8 *ebiten.Image
	SkeletonUp9 *ebiten.Image

	SkeletonLeft1 *ebiten.Image // idle
	SkeletonLeft2 *ebiten.Image
	SkeletonLeft3 *ebiten.Image
	SkeletonLeft4 *ebiten.Image
	SkeletonLeft5 *ebiten.Image
	SkeletonLeft6 *ebiten.Image
	SkeletonLeft7 *ebiten.Image
	SkeletonLeft8 *ebiten.Image
	SkeletonLeft9 *ebiten.Image

	SkeletonDown1 *ebiten.Image // idle
	SkeletonDown2 *ebiten.Image
	SkeletonDown3 *ebiten.Image
	SkeletonDown4 *ebiten.Image
	SkeletonDown5 *ebiten.Image
	SkeletonDown6 *ebiten.Image
	SkeletonDown7 *ebiten.Image
	SkeletonDown8 *ebiten.Image
	SkeletonDown9 *ebiten.Image

	SkeletonRight1 *ebiten.Image // idle
	SkeletonRight2 *ebiten.Image
	SkeletonRight3 *ebiten.Image
	SkeletonRight4 *ebiten.Image
	SkeletonRight5 *ebiten.Image
	SkeletonRight6 *ebiten.Image
	SkeletonRight7 *ebiten.Image
	SkeletonRight8 *ebiten.Image
	SkeletonRight9 *ebiten.Image
)

// animations
var (
	SkeletonMoveUpFrames    = []image.Image{}
	SkeletonMoveLeftFrames  = []image.Image{}
	SkeletonMoveDownFrames  = []image.Image{}
	SkeletonMoveRightFrames = []image.Image{}
)

func init() {
	img, err := loadPng(staticSpritesFS, "sprites/skeleton_move.png")
	if err != nil {
		panic(err)
	}

	SkeletonWalkImg = ebiten.NewImageFromImage(img)
	cells := CreateCells(9, 4, skeletonWalkWidth, skeletonWalkHeight)

	SkeletonUp1 = SkeletonWalkImg.SubImage(cells[0][0]).(*ebiten.Image)
	SkeletonUp2 = SkeletonWalkImg.SubImage(cells[1][0]).(*ebiten.Image)
	SkeletonUp3 = SkeletonWalkImg.SubImage(cells[2][0]).(*ebiten.Image)
	SkeletonUp4 = SkeletonWalkImg.SubImage(cells[3][0]).(*ebiten.Image)
	SkeletonUp5 = SkeletonWalkImg.SubImage(cells[4][0]).(*ebiten.Image)
	SkeletonUp6 = SkeletonWalkImg.SubImage(cells[5][0]).(*ebiten.Image)
	SkeletonUp7 = SkeletonWalkImg.SubImage(cells[6][0]).(*ebiten.Image)
	SkeletonUp8 = SkeletonWalkImg.SubImage(cells[7][0]).(*ebiten.Image)
	SkeletonUp9 = SkeletonWalkImg.SubImage(cells[8][0]).(*ebiten.Image)

	SkeletonLeft1 = SkeletonWalkImg.SubImage(cells[0][1]).(*ebiten.Image)
	SkeletonLeft2 = SkeletonWalkImg.SubImage(cells[1][1]).(*ebiten.Image)
	SkeletonLeft3 = SkeletonWalkImg.SubImage(cells[2][1]).(*ebiten.Image)
	SkeletonLeft4 = SkeletonWalkImg.SubImage(cells[3][1]).(*ebiten.Image)
	SkeletonLeft5 = SkeletonWalkImg.SubImage(cells[4][1]).(*ebiten.Image)
	SkeletonLeft6 = SkeletonWalkImg.SubImage(cells[5][1]).(*ebiten.Image)
	SkeletonLeft7 = SkeletonWalkImg.SubImage(cells[6][1]).(*ebiten.Image)
	SkeletonLeft8 = SkeletonWalkImg.SubImage(cells[7][1]).(*ebiten.Image)
	SkeletonLeft9 = SkeletonWalkImg.SubImage(cells[8][1]).(*ebiten.Image)

	SkeletonDown1 = SkeletonWalkImg.SubImage(cells[0][2]).(*ebiten.Image)
	SkeletonDown2 = SkeletonWalkImg.SubImage(cells[1][2]).(*ebiten.Image)
	SkeletonDown3 = SkeletonWalkImg.SubImage(cells[2][2]).(*ebiten.Image)
	SkeletonDown4 = SkeletonWalkImg.SubImage(cells[3][2]).(*ebiten.Image)
	SkeletonDown5 = SkeletonWalkImg.SubImage(cells[4][2]).(*ebiten.Image)
	SkeletonDown6 = SkeletonWalkImg.SubImage(cells[5][2]).(*ebiten.Image)
	SkeletonDown7 = SkeletonWalkImg.SubImage(cells[6][2]).(*ebiten.Image)
	SkeletonDown8 = SkeletonWalkImg.SubImage(cells[7][2]).(*ebiten.Image)
	SkeletonDown9 = SkeletonWalkImg.SubImage(cells[8][2]).(*ebiten.Image)

	SkeletonRight1 = SkeletonWalkImg.SubImage(cells[0][3]).(*ebiten.Image)
	SkeletonRight2 = SkeletonWalkImg.SubImage(cells[1][3]).(*ebiten.Image)
	SkeletonRight3 = SkeletonWalkImg.SubImage(cells[2][3]).(*ebiten.Image)
	SkeletonRight4 = SkeletonWalkImg.SubImage(cells[3][3]).(*ebiten.Image)
	SkeletonRight5 = SkeletonWalkImg.SubImage(cells[4][3]).(*ebiten.Image)
	SkeletonRight6 = SkeletonWalkImg.SubImage(cells[5][3]).(*ebiten.Image)
	SkeletonRight7 = SkeletonWalkImg.SubImage(cells[6][3]).(*ebiten.Image)
	SkeletonRight8 = SkeletonWalkImg.SubImage(cells[7][3]).(*ebiten.Image)
	SkeletonRight9 = SkeletonWalkImg.SubImage(cells[8][3]).(*ebiten.Image)

	SkeletonMoveUpFrames = []image.Image{
		// SkeletonUp1Sprite, // idle
		SkeletonUp2,
		SkeletonUp3,
		SkeletonUp4,
		SkeletonUp5,
		SkeletonUp6,
		SkeletonUp7,
		SkeletonUp8,
		SkeletonUp9,
	}

	SkeletonMoveLeftFrames = []image.Image{
		// SkeletonLeft1Sprite, // idle
		SkeletonLeft2,
		SkeletonLeft3,
		SkeletonLeft4,
		SkeletonLeft5,
		SkeletonLeft6,
		SkeletonLeft7,
		SkeletonLeft8,
		SkeletonLeft9,
	}

	SkeletonMoveDownFrames = []image.Image{
		// SkeletonDown1Sprite, // idle
		SkeletonDown2,
		SkeletonDown3,
		SkeletonDown4,
		SkeletonDown5,
		SkeletonDown6,
		SkeletonDown7,
		SkeletonDown8,
		SkeletonDown9,
	}

	SkeletonMoveRightFrames = []image.Image{
		// SkeletonRight1Sprite, // idle
		SkeletonRight2,
		SkeletonRight3,
		SkeletonRight4,
		SkeletonRight5,
		SkeletonRight6,
		SkeletonRight7,
		SkeletonRight8,
		SkeletonRight9,
	}
}
