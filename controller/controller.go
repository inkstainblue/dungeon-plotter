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

// DrawWall draws a wall between two points in grid space.
func (c *Controller) DrawWall(a, b canvas.Point) error {
	half := canvas.Pt(0.5, 0.5)

	for _, cv := range c.canvases {
		// TODO: Draw more interesting lines.
		if err := cv.Draw(a.Add(half), b.Add(half)); err != nil {
			return err
		}
	}

	return nil
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
