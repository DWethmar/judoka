package tilebitmasking

// shapes
// 1 represents a tile of the same type and 0 represents an different tile.
// Naming convention is all edges that have the same type: <NorthWest><North><NorthEast><East><SouthEast><South><SouthWest><West>Edges
const (
	// AllEdges is a fully enclosed tile of the same type
	//	111
	//	1X1
	//	111
	AllEdges = 0b11111111
	// NoEdges is a fully enclosed tile of different types
	//	000
	//	0X0
	//	000
	NoEdges = 0b00000000

	// Multiple edge Tiles, These tiles have three edges on the same side that differs from the central tile.

	// EastSoutheastSouthSouthwestWest is a tile that has edges on the east, south and west side that differs from the central tile.
	//	000
	//	1X1
	//	111
	EastSoutheastSouthSouthwestWest = 0b11111000
	// NorthwestNorthSouthSouthwestWest is a tile that has edges on the north, south and west side that differs from the central tile.
	//	110
	//	1X0
	//	110
	NorthwestNorthSouthSouthwestWest = 0b01101011
	// NorthwestNorthNortheastEastWest is a tile that has edges on the north, east and west side that differs from the central tile.
	//	111
	//	1X1
	//	000
	NorthwestNorthNortheastEastWest = 0b00011111
	// NorthNortheastEastSouthEastSouth is a tile that has edges on the north, east and south side that differs from the central tile.
	//	0X1
	//	011
	NorthNortheastEastSouthEastSouth = 0b11010110

	// corner tiles

	// EastSoutheastSouthEdges is a tile with different north-west, north, north-east, south-west and west edges
	//	000
	//	0X1
	//	011
	EastSoutheastSouthEdges = 0b11010000
	// SouthSouthwestWestEdges is a tile with different north-west, north, north-east, east and south-east edges
	//	000
	//	1X0
	//	110
	SouthSouthwestWestEdges = 0b01101000
	// NorthwestNorthWestEdges is a tile with different north-east, east, south-east, south, south-west edges
	//	110
	//	1X0
	//	000
	NorthwestNorthWestEdges = 0b00001011
	// NorthNortheastEastEdges is a tile with different north, south-east, south, south-west and west edges
	//	011
	//	0X1
	//	000
	NorthNortheastEastEdges = 0b00010110

	// Two diagonal edges, These tiles have two diagonal edges on the same axis that differs from the central tile.

	// NorthNortheastEastSouthSouthwestWest is a tile that has edges on all sides except the north-west and south-east.
	//	011
	//	1X1
	//	110
	NorthNortheastEastSouthSouthwestWest = 0b01111110
	// NorthwestNorthEastSoutheastSouthWest is a tile that has edges on all sides except the north-west and south-east.
	// 110
	// 1X1
	// 011
	NorthwestNorthEastSoutheastSouthWest = 0b11011011

	// Four diagonal edges, These tiles have four diagonal edges that differs from the central tile.

	// HorizontalAndVerticalEdges is a tile with different north west, north east, south east and south west edges
	//	010
	//	1X1
	//	010
	HorizontalAndVerticalEdges = 0b01011010

	// line ends, These tiles can be used to create lines.

	// SouthEdge is a tile with different north west, north, north east, east, south east, south west and west edges
	//	000
	//	0X0
	//	010
	SouthEdge = 0b01000000
	// WestEdge is a tile with different north west, north, north east, east, south east, south and south west edges
	//	000
	//	1X0
	//	000
	WestEdge = 0b00001000
	// NorthEdge is a tile with different north west, north east, east, south east, south, south west and west edges
	//	010
	//	0X0
	//	000
	NorthEdge = 0b00000010
	// EastEdge is a tile with different north west, north, north east, south east, south, south west and west edges
	//	000
	//	0X1
	//	000
	EastEdge = 0b00010000

	// lines, These tiles can be used to create lines.

	// HorizontalEdges is a tile with west and east edges
	//	000
	//	1X1
	//	000
	HorizontalEdges = 0b00011000
	// VerticalEdges is a tile with north and south edges
	//	010
	//	0X0
	//	010
	VerticalEdges = 0b01000010

	// T shaped tiles, These tiles can be used to create T shaped lines.

	// EastSouthWestEdges is a T shaped tile that has edges on the east, south and west side.
	//	000
	//	1X1
	//	010
	EastSouthWestEdges = 0b01011000
	// NorthEastWest is a T shaped tile that has edges on the north, east and west side.
	//	010
	//	1X1
	//	000
	NorthEastWestEdges = 0b00011010
	// NorthSouthWestEdges is a T shaped tile that has edges on the north, south and west side.
	//	010
	//	1X0
	//	010
	NorthSouthWestEdges = 0b01001010
	// NorthEastSouthEdges is a T shaped tile that has edges on the north, east and south side.
	//	010
	//	0X1
	//	010
	NorthEastSouthEdges = 0b01010010

	// inward corners, These tiles can be used to create inward corners.

	// NorthwestNorthEastSoutheastSouthSouthWestWestEdges is a tile that has edges on the north, east, south and west side, but not on the north-east side.
	//	110
	//	1X1
	//	111
	NorthwestNorthEastSoutheastSouthSouthWestWestEdges = 0b11111011
	// NorthNortheastEastSoutheastSouthSouthWestWestEdges is a tile that has edges on the north, east, south and west side, but not on the north-west side.
	//	011
	//	1X1
	//	111
	NorthNortheastEastSoutheastSouthSouthWestWestEdges = 0b11111110
	// NorthwestNorthNorthEastEastSouthSouthwestWestEdges is a tile that has edges on the north, east, south and west side, but not on the south-east side.
	//	111
	//	1X1
	//	110
	NorthwestNorthNorthEastEastSouthSouthwestWestEdges = 0b01111111
	// NorthwestNorthNortheastEastSoutheastSouthWestEdges is a tile that has edges on the north, east, south and west side, but not on the south-west side.
	//	111
	//	1X1
	//	011
	NorthwestNorthNortheastEastSoutheastSouthWestEdges = 0b11011111

	// inward corner on edges. This combines the inward corner with an edge.

	// EastSoutheastSouthWestEdges is a tile that has edges on the north side and south west size.
	//	000
	//	1X1
	//	011
	EastSoutheastSouthWestEdges = 0b11011000
	// NorthSouthSouthwestEastEdges is a tile that has edges on the north west, and east side
	//	010
	//	1X0
	//	110
	NorthSouthSouthwestEastEdges = 0b01101010
	// NorthEastSoutheastSouthEdges is a tile that has edges on the west side and north east.
	//	010
	//	0X1
	//	011
	NorthEastSoutheastSouthEdges = 0b11010010
	// NorthNortheastEastWestEdges is a tile that has edges on the north west and south side.
	//	011
	//	1x1
	//	000
	NorthNortheastEastWestEdges = 0b00011110
	// EastSouthSouthwestWestEdges is a tile that has edges on the north side and south east.
	//	000
	//	1X1
	//	110
	EastSouthSouthwestWestEdges = 0b01111000
	// NorthwestNorthSouthWestEdges is a tile that has edges on east side and south west.
	//	110
	//	1X0
	//	010
	NorthwestNorthSouthWestEdges = 0b01001011
	// NorthNorthEastEastSouthEdges is a tile that has edges on the north west and south east.
	//	011
	//	0X1
	//	010
	NorthNorthEastEastSouthEdges = 0b01010110
	// NorthwestNorthEastWestEdges is a tile that has edges on the south side and north east.
	//	110
	//	1X1
	//	000
	NorthwestNorthEastWestEdges = 0b00011011

	// double inward corners. This combines the two inward corners on the same side.

	// NorthEastSouthEastSouthSouthwestWestEdges is a tile that has alls edges except the north west and north east.
	//	010
	//	1X1
	//	111
	NorthEastSouthEastSouthSouthwestWestEdges = 0b11111010
	// NorthwestNorthEastSouthSouthwestWestEdges is a tile that has alls edges except the north east and south east.
	//	110
	//	1X1
	//	110
	NorthwestNorthEastSouthSouthwestWestEdges = 0b01111011
	// NorthNortheastEastSoutheastSouthWestEdges is a tile that has alls edges except the north west and south west.
	//	011
	//	1X1
	//	011
	NorthNortheastEastSoutheastSouthWestEdges = 0b11011110
	// NorthwestNorthNortheastEastSouthWestEdges is a tile that has alls edges except the north west and south east.
	//	111
	//	1X1
	//	010
	NorthwestNorthNortheastEastSouthWestEdges = 0b01011111

	// turn corners. These tiles can be used to create turn corners.

	// EastSouthEdges is a tile that has edges on the north side and south east.
	//	000
	//	0x1
	//	010
	EastSouthEdges = 0b01010000
	// SouthWestEdges is a tile that has edges on the south and west side.
	//	000
	//	1X0
	//	010
	SouthWestEdges = 0b01001000
	// NorthEastEdges is a tile that has edges on the north and east side.
	//	010
	//	0X1
	//	000
	NorthEastEdges = 0b00010010
	// NorthWestEdges is a tile that has edges on the north and west side.
	//	010
	//	1X0
	//	000
	NorthWestEdges = 0b00001010

	// fish shaped corners, all diagonal edges are the different except one.

	// NorthEastSouthSouthwestWestEdges
	//	010
	//	1X1
	//	110
	NorthEastSouthSouthwestWestEdges = 0b01111010
	// NorthwestNorthEastSouthWestEdges
	//	110
	//	1X1
	//	010
	NorthwestNorthEastSouthWestEdges = 0b01011011
	// NorthNorthEastEastSouthEdges
	//	011
	//	1X1
	//	010
	NorthNortheastEastSouthWestEdges = 0b01011110
	// NorthEastSoutheastSouthWestEdges
	//	010
	//	1X1
	//	011
	NorthEastSoutheastSouthWestEdges = 0b11011010
)
