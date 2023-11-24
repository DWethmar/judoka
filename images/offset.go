package images

import "github.com/hajimehoshi/ebiten/v2"

// ActorOffsets returns the offsets for an actor image
// The offsets are calculated by:
//
//	move sprite image left by half width.
//	move sprite image up by height.
func ActorOffsets(image *ebiten.Image) (int, int) {
	w := image.Bounds().Max.X - image.Bounds().Min.X
	h := image.Bounds().Max.Y - image.Bounds().Min.Y
	return -(w / 2), -h
}
