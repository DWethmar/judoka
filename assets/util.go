package assets

import "image"

// CreateCells creates a slice of image.Rectangle pointers
// that can be used to draw sprites from a sprite sheet.
//
// the rectangles are from left to right, top to bottom
// cells[row][column]
func CreateCells(columns, rows, width, height int) [][]image.Rectangle {
	cells := make([][]image.Rectangle, columns)

	for x := 0; x < columns; x++ {
		cells[x] = make([]image.Rectangle, rows)
		for y := 0; y < rows; y++ {
			cells[x][y] = image.Rect(
				x*width,
				y*height,
				(x*width)+width,
				(y*height)+height,
			)
		}
	}

	return cells
}
