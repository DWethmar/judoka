package direction

//go:generate stringer -type=Direction
type Direction int

const (
	None Direction = iota
	Top
	Bottom
	Left
	Right
	TopLeft
	TopRight
	BottomLeft
	BottomRight
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
