package canvas

import (
	"image"
)

type (
	Canvas interface {
		// Clear ensures a given rectangle on the canvas is empty,
		// ready for drawing to.
		Clear(rect image.Rectangle) error

		// Draw draws a line between two points on the canvas.
		// The points should be provided in grid space.
		Draw(a, b Point) error

		// WaitForQuit blocks until the canvas has exited.
		WaitForQuit()
	}
)
