package controller

import (
	"image"

	"github.com/inkstainblue/dungeon-plotter/canvas"
)

type (
	Controller struct {
		canvas canvas.Canvas
	}
)

// New creates a new Controller using the given canvas.Canvas.
func New(canvas canvas.Canvas) (c Controller) {
	c.canvas = canvas

	return
}

// DrawWall draws a wall between two points in grid space.
func (c *Controller) DrawWall(a, b image.Point) error {
	// TODO: Draw more interesting lines.
	return c.canvas.Draw(a, b)
}

// WaitForQuit blocks until the Controller has exited.
func (c *Controller) WaitForQuit() {
	c.canvas.WaitForQuit()
}