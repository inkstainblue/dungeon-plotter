package controller

import (
	"fmt"
	"sync"

	"github.com/inkstainblue/dungeon-plotter/canvas"
	"github.com/inkstainblue/dungeon-plotter/input"
)

type (
	Controller struct {
		canvases []canvas.Canvas
		inputs   []input.InputHandler
	}
)

// New creates a new controller using the given canvases for output.
func New(canvases []canvas.Canvas, inputs []input.InputHandler) (c Controller) {
	c.canvases = canvases
	c.inputs = inputs

	c.handleInput()

	return
}

// DrawPath draws a path between two points in grid space.
// FIXME: Draw in the correct direction.
// FIXME: Draw along the vector joining the two points.
func (c *Controller) DrawPath(a, b canvas.Point) error {
	half := canvas.Pt(0.5, 0.5)

	a, b = a.Add(half), b.Add(half)

	p0 := a
	p1 := a.Add(canvas.Pt(0.25, 0))

	for {
		if err := c.draw(p0, p1); err != nil {
			return err
		}

		p0 = p1.Add(canvas.Pt(0.5, 0))
		p1 = p0.Add(canvas.Pt(0.5, 0))

		if p0.X >= b.X {
			break
		}

		if p1.X > b.X {
			p1 = b
		}
	}

	return nil
}

// DrawWall draws a wall between two points in grid space.
func (c *Controller) DrawWall(a, b canvas.Point) error {
	half := canvas.Pt(0.5, 0.5)

	// TODO: Draw more interesting lines.
	return c.draw(a.Add(half), b.Add(half))
}

// WaitForQuit blocks until the controller has exited.
func (c *Controller) WaitForQuit() {
	wg := new(sync.WaitGroup)
	wg.Add(len(c.canvases))

	for _, cv := range c.canvases {
		go func() {
			cv.WaitForQuit()
			wg.Done()
		}()
	}

	wg.Wait()
}

func (c *Controller) draw(a, b canvas.Point) error {
	for _, cv := range c.canvases {
		if err := cv.Draw(a, b); err != nil {
			return err
		}
	}

	return nil
}

func (c *Controller) handleInput() {
	for _, in := range c.inputs {
		go func() {
			for {
				code, label := in.WaitForInput()

				fmt.Println(label, code)
			}
		}()
	}
}
