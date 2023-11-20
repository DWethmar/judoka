package direction

type Direction string

const (
	None        Direction = "none"
	Top         Direction = "top"
	Bottom      Direction = "bottom"
	Left        Direction = "left"
	Right       Direction = "right"
	TopLeft     Direction = "top_left"
	TopRight    Direction = "top_right"
	BottomLeft  Direction = "bottom_left"
	BottomRight Direction = "bottom_right"
)

func Get(sX, sY, dX, dY int) Direction {
	deltaX := dX - sX
	deltaY := dY - sY

	// Diagonal checks
	if deltaX > 0 && deltaY > 0 {
		return BottomRight
	} else if deltaX < 0 && deltaY > 0 {
		return BottomLeft
	} else if deltaX > 0 && deltaY < 0 {
		return TopRight
	} else if deltaX < 0 && deltaY < 0 {
		return TopLeft
	}

	// Straight direction checks
	if sX == dX {
		if dY > sY {
			return Bottom
		}
		if dY < sY {
			return Top
		}
	} else if sY == dY {
		if dX > sX {
			return Right
		}
		if dX < sX {
			return Left
		}
	}

	// If we reach here, there's no clear direction or they're the same coordinates
	return None
}
