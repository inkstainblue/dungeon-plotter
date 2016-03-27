package canvas

import (
	"fmt"
	"image"
	"sync"

	"github.com/veandco/go-sdl2/sdl"
)

type (
	ScreenCanvas struct {
		Canvas

		window    *sdl.Window
		renderer  *sdl.Renderer
		eventLoop *sync.WaitGroup
	}
)

func NewScreenCanvas() (sc ScreenCanvas) {
	var err error

	if err = sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		// FIXME: Handle errors properly.
		panic(err)
	}

	title := "Dungeon Plotter"
	x, y, w, h := sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 800, 600

	if sc.window, err = sdl.CreateWindow(title, x, y, w, h, sdl.WINDOW_SHOWN); err != nil {
		// FIXME: Handle errors properly.
		panic(err)
	}

	if sc.renderer, err = sdl.CreateRenderer(sc.window, -1, sdl.RENDERER_ACCELERATED); err != nil {
		// FIXME: Handle errors properly.
		panic(err)
	}

	sc.Clear(image.Rect(0, 0, w, h))

	sc.eventLoop = &sync.WaitGroup{}
	sc.eventLoop.Add(1)

	go func() {
		defer sc.eventLoop.Done()
		defer sc.window.Destroy()
		defer sc.renderer.Destroy()

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

func (sc *ScreenCanvas) Clear(rect image.Rectangle) error {
	sc.renderer.Clear()

	return nil
}

func (sc *ScreenCanvas) Draw(a, b image.Point) error {
	return nil
}
