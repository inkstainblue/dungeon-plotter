package canvas

import (
	"image"
)

type (
	Canvas interface {
		// Clear ensures a given rectangle on the Canvas is empty,
		// ready for drawing to.
		Clear(rect image.Rectangle) error

		// Draw draws a line between two points on the canvas.
		Draw(a, b image.Point) error
	}
)
