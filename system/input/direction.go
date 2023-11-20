package input

import "github.com/hajimehoshi/ebiten/v2"

// Direction returns the direction of the input.
func Direction() (X int, Y int) {
	// LEFT
	for _, k := range Left {
		if ebiten.IsKeyPressed(k) {
			X -= 1
			break
		}
	}

	// RIGHT
	for _, k := range Right {
		if ebiten.IsKeyPressed(k) {
			X += 1
			break
		}
	}

	// UP
	for _, k := range Up {
		if ebiten.IsKeyPressed(k) {
			Y -= 1
			break
		}
	}

	// DOWN
	for _, k := range Down {
		if ebiten.IsKeyPressed(k) {
			Y += 1
			break
		}
	}

	return
}
