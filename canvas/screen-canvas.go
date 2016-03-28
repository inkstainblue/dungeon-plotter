package canvas

import (
	"errors"
	"fmt"
	"image"
	"sync"

	"github.com/veandco/go-sdl2/sdl"
)

type (
	ScreenCanvas struct {
		Canvas

		gridScale int

		eventLoop *sync.WaitGroup
		window    *sdl.Window
		renderer  *sdl.Renderer
	}
)

func NewScreenCanvas(gridWidth, gridHeight int) (sc ScreenCanvas) {
	sc.gridScale = 10

	var err error

	if err = sdl.Init(sdl.INIT_VIDEO); err != nil {
		// FIXME: Handle errors properly.
		panic(err)
	}

	title := "Dungeon Plotter"
	x, y, w, h := sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, gridWidth*sc.gridScale, gridHeight*sc.gridScale

	if sc.window, err = sdl.CreateWindow(title, x, y, w, h, sdl.WINDOW_SHOWN); err != nil {
		// FIXME: Handle errors properly.
		panic(err)
	}

	if sc.renderer, err = sdl.CreateRenderer(sc.window, -1, sdl.RENDERER_ACCELERATED); err != nil {
		// FIXME: Handle errors properly.
		panic(err)
	}

	if err = sc.Clear(image.Rect(0, 0, gridWidth, gridHeight)); err != nil {
		// FIXME: Handle errors properly.
		panic(err)
	}

	sc.eventLoop = new(sync.WaitGroup)
	sc.eventLoop.Add(1)

	go func() {
		defer func() {
			sc.renderer.Destroy()
			sc.window.Destroy()
			sc.eventLoop.Done()
		}()

		for {
			switch e := sdl.WaitEvent().(type) {
			case *sdl.QuitEvent:
				return
			default:
				fmt.Printf("%T: %+v\n", e, e)
			}
		}
	}()

	return
}

func (sc *ScreenCanvas) WaitForQuit() {
	sc.eventLoop.Wait()
}

// Clear overwrites the entire ScreenCanvas with a white background.
// The rectangle to clear is ignored, but an error is returned if the rectangle
// is larger than the drawable area.
func (sc *ScreenCanvas) Clear(rect image.Rectangle) error {
	if err := sc.renderer.SetDrawColor(255, 255, 255, 255); err != nil {
		return err
	}

	if err := sc.renderer.Clear(); err != nil {
		return err
	}

	sc.renderer.Present()

	r := sc.gridRectToScreen(rect)
	w, h, err := sc.renderer.GetRendererOutputSize()

	switch {
	case err != nil:
		return err
	case int(r.W) > w || int(r.H) > h:
		return errors.New("The rectangle to clear is larger than the drawable area")
	}

	return nil
}

// Draw draws a line between two points in grid space on the ScreenCanvas.
// The points are converted from grid space to screen space.
// FIXME: This needs to handle float inputs for drawing sub grid shapes.
func (sc *ScreenCanvas) Draw(a, b image.Point) error {
	points := []sdl.Point{
		sc.gridPointToScreen(a),
		sc.gridPointToScreen(b),
	}

	if err := sc.renderer.SetDrawColor(0, 0, 0, 255); err != nil {
		return err
	}

	if err := sc.renderer.DrawLines(points); err != nil {
		return err
	}

	sc.renderer.Present()

	return nil
}

func (sc *ScreenCanvas) gridPointToScreen(p image.Point) sdl.Point {
	return sdl.Point{
		X: int32(p.X * sc.gridScale),
		Y: int32(p.Y * sc.gridScale),
	}
}

func (sc *ScreenCanvas) gridRectToScreen(r image.Rectangle) sdl.Rect {
	return sdl.Rect{
		X: int32(r.Min.X * sc.gridScale),
		Y: int32(r.Min.Y * sc.gridScale),
		W: int32(r.Dx() * sc.gridScale),
		H: int32(r.Dy() * sc.gridScale),
	}
}
