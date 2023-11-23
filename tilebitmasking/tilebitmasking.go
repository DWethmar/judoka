// tilebitmasking is a method for automatically selecting the appropriate sprite from a defined
// tileset. This allows you to place a generic placeholder tile everywhere you want a particular
// type of terrain to appear instead of hand placing a potentially enormous selection of various
// tiles.
package tilebitmasking

// Bitmasking
const (
	// neighbors
	// 0 represents an different tile and 1 represents a tile of the same type.
	//
	//	1	2	4
	//	8	X	16
	//	32	64	128
	NorthWest = 0b00000001 // 1
	North     = 0b00000010 // 2
	NorthEast = 0b00000100 // 4
	West      = 0b00001000 // 8
	East      = 0b00010000 // 16
	SouthWest = 0b00100000 // 32
	South     = 0b01000000 // 64
	SouthEast = 0b10000000 // 128
)

// Neighborhood represents the neighbors of a tile in the cardinal directions.
// The center tile is the tile we are currently checking.
type Neighborhood struct {
	NorthWest int
	North     int
	NorthEast int
	West      int
	Center    int
	East      int
	SouthWest int
	South     int
	SouthEast int
}

// To eliminate the redundancies, we add an extra condition to our Boolean directional
// check: when checking for the presence of bordering corner tiles, we also have to
// check for neighboring tiles in the four cardinal directions (directly North, East, South, or West).
// For example, the tile to the North-East is neighbored by existing tiles,
// whereas the tiles to the South-West and South-East are not.
// This means that the South-West and South-East tiles are not included in the bitmasking calculation.
func Calculate(neighbors *Neighborhood) int {
	currentTile := neighbors.Center
	bitmask := 0

	// cardinal directions
	northSame := neighbors.North == currentTile
	eastSame := neighbors.East == currentTile
	southSame := neighbors.South == currentTile
	westSame := neighbors.West == currentTile

	if northSame {
		bitmask |= North
	}

	if eastSame {
		bitmask |= East
	}

	if southSame {
		bitmask |= South
	}

	if westSame {
		bitmask |= West
	}

	// When considering corner tiles (Northeast, Northwest, Southeast, Southwest),
	// also check if the neighboring tiles in the
	// cardinal directions (North, East, South, West) are different.
	// If a corner tile is different, but the adjacent cardinal tiles are not,
	// ignore this corner tile in the bitmask calculation.
	if neighbors.NorthWest == currentTile && northSame && westSame {
		bitmask |= NorthWest
	}

	if neighbors.NorthEast == currentTile && northSame && eastSame {
		bitmask |= NorthEast
	}

	if neighbors.SouthEast == currentTile && southSame && eastSame {
		bitmask |= SouthEast
	}

	if neighbors.SouthWest == currentTile && southSame && westSame {
		bitmask |= SouthWest
	}

	return bitmask
}
