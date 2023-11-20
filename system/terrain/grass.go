package terrain

func GrassToDirtTransition(neighbors TileNeighbors) bool {
	return neighbors.TopLeft == GrassTile &&
		neighbors.Top == GrassTile &&
		neighbors.TopRight == GrassTile &&
		neighbors.Left == GrassTile &&
		neighbors.Right == GrassTile &&
		neighbors.BottomLeft == GrassTile &&
		neighbors.Bottom == GrassTile
}

// DIAGANALS

// GrassToDirtBottomRight
//
// G G G
// G G G
// G G D
func GrassToDirtBottomRight(neighbors TileNeighbors) bool {
	return neighbors.TopLeft == GrassTile &&
		neighbors.Top == GrassTile &&
		neighbors.TopRight == GrassTile &&
		neighbors.Left == GrassTile &&
		neighbors.Right == GrassTile &&
		neighbors.BottomLeft == GrassTile &&
		neighbors.Bottom == GrassTile &&
		neighbors.BottomRight == DirtTile
}

// GrassToDirtTopLeft
//
// D G G
// G G G
// G G G
func GrassToDirtTopLeft(neighbors TileNeighbors) bool {
	return neighbors.TopLeft == DirtTile &&
		neighbors.Top == GrassTile &&
		neighbors.TopRight == GrassTile &&
		neighbors.Left == GrassTile &&
		neighbors.Right == GrassTile &&
		neighbors.BottomLeft == GrassTile &&
		neighbors.Bottom == GrassTile &&
		neighbors.BottomRight == GrassTile
}

// GrassToDirtTopRightTransition
//
// G G D
// G G G
// G G G
func GrassToDirtTopRight(neighbors TileNeighbors) bool {
	return neighbors.TopLeft == GrassTile &&
		neighbors.Top == GrassTile &&
		neighbors.TopRight == DirtTile &&
		neighbors.Left == GrassTile &&
		neighbors.Right == GrassTile &&
		neighbors.BottomLeft == GrassTile &&
		neighbors.Bottom == GrassTile &&
		neighbors.BottomRight == GrassTile
}

// GrassToDirtBottomLeftTransition
//
// G G G
// G G G
// D G G
func GrassToDirtBottomLeft(neighbors TileNeighbors) bool {
	return neighbors.TopLeft == GrassTile &&
		neighbors.Top == GrassTile &&
		neighbors.TopRight == GrassTile &&
		neighbors.Left == GrassTile &&
		neighbors.Right == GrassTile &&
		neighbors.BottomLeft == DirtTile &&
		neighbors.Bottom == GrassTile &&
		neighbors.BottomRight == GrassTile
}

// CORNERS

// GrassToDirtBottomLeftCorner
// ? G G
// D X G
// D D ?
func GrassToDirtBottomLeftCorner(neighbors TileNeighbors) bool {
	return neighbors.Top == GrassTile &&
		neighbors.TopRight == GrassTile &&
		neighbors.Left == DirtTile &&
		neighbors.Right == GrassTile &&
		neighbors.BottomLeft == DirtTile &&
		neighbors.Bottom == DirtTile
}

// GrassToDirtBottomRightCorner
// G G D
// G X D
// D D D
func GrassToDirtBottomRightCorner(neighbors TileNeighbors) bool {
	return neighbors.TopLeft == GrassTile &&
		neighbors.Top == GrassTile &&
		neighbors.TopRight == DirtTile &&
		neighbors.Left == GrassTile &&
		neighbors.Right == DirtTile &&
		neighbors.BottomLeft == DirtTile &&
		neighbors.Bottom == DirtTile &&
		neighbors.BottomRight == DirtTile
}

// GrassToDirtTopLeftCorner
// D D D
// D X G
// D G G
func GrassToDirtTopLeftCorner(neighbors TileNeighbors) bool {
	return neighbors.TopLeft == DirtTile &&
		neighbors.Top == DirtTile &&
		neighbors.TopRight == DirtTile &&
		neighbors.Left == DirtTile &&
		neighbors.Right == GrassTile &&
		neighbors.BottomLeft == GrassTile &&
		neighbors.Bottom == GrassTile &&
		neighbors.BottomRight == DirtTile
}

// GrassToDirtTopRightCorner
// D D D
// G X D
// G G D
func GrassToDirtTopRightCorner(neighbors TileNeighbors) bool {
	return neighbors.TopLeft == DirtTile &&
		neighbors.Top == DirtTile &&
		neighbors.TopRight == DirtTile &&
		neighbors.Left == GrassTile &&
		neighbors.Right == DirtTile &&
		neighbors.BottomLeft == DirtTile &&
		neighbors.Bottom == GrassTile &&
		neighbors.BottomRight == GrassTile
}

// SIDES

// GrassToDirtTopSide
// D D D
// G X G
// G G G
func GrassToDirtTopSide(neighbors TileNeighbors) bool {
	return neighbors.TopLeft == DirtTile &&
		neighbors.Top == DirtTile &&
		neighbors.TopRight == DirtTile &&
		neighbors.Left == GrassTile &&
		neighbors.Right == GrassTile &&
		neighbors.BottomLeft == GrassTile &&
		neighbors.Bottom == GrassTile &&
		neighbors.BottomRight == GrassTile
}

// GrassToDirtBottomSide
// G G G
// G X G
// D D D
func GrassToDirtBottomSide(neighbors TileNeighbors) bool {
	return neighbors.TopLeft == GrassTile &&
		neighbors.Top == GrassTile &&
		neighbors.TopRight == GrassTile &&
		neighbors.Left == GrassTile &&
		neighbors.Right == GrassTile &&
		neighbors.BottomLeft == DirtTile &&
		neighbors.Bottom == DirtTile &&
		neighbors.BottomRight == DirtTile
}

// GrassToDirtLeftSide
// D G G
// D X G
// D G G
func GrassToDirtLeftSide(neighbors TileNeighbors) bool {
	return neighbors.TopLeft == DirtTile &&
		neighbors.Top == GrassTile &&
		neighbors.TopRight == GrassTile &&
		neighbors.Left == DirtTile &&
		neighbors.Right == GrassTile &&
		neighbors.BottomLeft == DirtTile &&
		neighbors.Bottom == GrassTile &&
		neighbors.BottomRight == DirtTile
}

// GrassToDirtRightSide
// G G D
// G X D
// G G D
func GrassToDirtRightSide(neighbors TileNeighbors) bool {
	return neighbors.TopLeft == GrassTile &&
		neighbors.Top == GrassTile &&
		neighbors.TopRight == DirtTile &&
		neighbors.Left == GrassTile &&
		neighbors.Right == DirtTile &&
		neighbors.BottomLeft == GrassTile &&
		neighbors.Bottom == GrassTile &&
		neighbors.BottomRight == DirtTile
}
