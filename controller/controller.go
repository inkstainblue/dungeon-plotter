package controller

import (
	"fmt"
	"math"
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
func (c *Controller) DrawPath(a, b canvas.Point) error {
	// TODO: Stop repeated paths from overlapping with each other?
	//		 Maybe use a different offset based on the draw direction.
	return c.drawDashed(centerPoint(a), centerPoint(b), 0.5, 0.5)
}

// DrawWall draws a wall between two points in grid space.
func (c *Controller) DrawWall(a, b canvas.Point) error {
	// TODO: Draw more interesting lines.
	return c.draw(centerPoint(a), centerPoint(b))
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

func (c *Controller) drawDashed(a, b canvas.Point, dashLen, gapLen float64) error {
	ab := b.Sub(a)
	abLen := math.Sqrt(ab.X*ab.X + ab.Y*ab.Y)

	if abLen < 0.0001 {
		return c.draw(a, b)
	}

	ab = ab.Div(abLen)

	dash, gap := ab.Mul(dashLen), ab.Mul(gapLen)

	p0, p1 := a, a.Add(dash.Mul(0.5))

	min, max := a, b

	if a.X > b.X {
		max.X = a.X
		min.X = b.X
	}

	if a.Y > b.Y {
		max.Y = a.Y
		min.Y = b.Y
	}

	for {
		if err := c.draw(p0, p1); err != nil {
			return err
		}

		p0 = p1.Add(gap)
		p1 = p0.Add(dash)

		// TODO: Only break if x and y are both out of range.
		//	     Otherwise clamp it.
		if p0.X < min.X || p0.X > max.X || p0.Y < min.Y || p0.Y > max.Y {
			break
		}

		switch {
		case p1.X < min.X:
			p1.X = min.X
		case p1.X > max.X:
			p1.X = max.X
		}

		switch {
		case p1.Y < min.Y:
			p1.Y = min.Y
		case p1.Y > max.Y:
			p1.Y = max.Y
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

func centerPoint(p canvas.Point) canvas.Point {
	return p.Add(canvas.Pt(0.5, 0.5))
}
