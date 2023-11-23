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

	// Edge Tiles, These tiles have one edge that differs from the central tile.

	// // NorthEdge is a tile with a different north edge
	// //	101
	// //	1X1
	// //	111
	// NorthEdge = 0b11111101
	// // EdgeRight is a tile with a different east edge
	// //	111
	// //	1X0
	// //	111
	// EastEdge = 0b11101111
	// // SouthEdge is a tile with a different south edge
	// //	111
	// //	1X1
	// //	101
	// SouthEdge = 0b10111111
	// // WestEdge is a tile with a different west edge
	// //	111
	// //	0X1
	// //	111
	// WestEdge = 0b11110111

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

	// EastSouthWestEdges is a tile that has edges on the east, south and west side.
	//	000
	//	1X1
	//	010
	EastSouthWestEdges = 0b01011000
	//
)
