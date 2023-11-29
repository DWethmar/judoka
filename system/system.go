package system

import "github.com/hajimehoshi/ebiten/v2"

// System is a system that can be updated.
type System interface {
	Update() error
	Draw(screen *ebiten.Image) error
}
