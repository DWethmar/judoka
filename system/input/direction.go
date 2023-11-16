package input

import "github.com/hajimehoshi/ebiten/v2"

// Direction returns the direction of the input.
func Direction() (X int, Y int) {
	var left, right, up, down bool

	// LEFT
	for _, k := range Left {
		if ebiten.IsKeyPressed(k) {
			left = true
			break
		}
	}

	// RIGHT
	for _, k := range Right {
		if ebiten.IsKeyPressed(k) {
			right = true
			break
		}
	}

	// UP
	for _, k := range Up {
		if ebiten.IsKeyPressed(k) {
			up = true
			break
		}
	}

	// DOWN
	for _, k := range Down {
		if ebiten.IsKeyPressed(k) {
			down = true
			break
		}
	}

	if left {
		X = -1
	}

	if right {
		X = 1
	}

	if up {
		Y = -1
	}

	if down {
		Y = 1
	}

	return
}
