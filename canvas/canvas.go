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
		// The points should be provided in grid space.
		// FIXME: This needs to handle float inputs for drawing sub grid shapes.
		Draw(a, b image.Point) error
	}
)
