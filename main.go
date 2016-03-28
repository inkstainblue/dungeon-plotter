package main

import (
	"fmt"
	"image"

	"github.com/inkstainblue/dungeon-plotter/canvas"
	"github.com/inkstainblue/dungeon-plotter/controller"
	"github.com/inkstainblue/dungeon-plotter/input"
)

const (
	maxGridWidth  = 100
	maxGridHeight = 100
)

func main() {
	sc := canvas.NewScreenCanvas(maxGridWidth, maxGridHeight)

	fmt.Printf("%T is ready\n", sc)

	k := input.NewKeyboard()

	fmt.Printf("%T is ready\n", k)

	c := controller.New([]canvas.Canvas{&sc}, []input.InputHandler{&k})

	fmt.Printf("%T is ready\n", c)

	c.DrawWall(image.Pt(0, 0), image.Pt(10, 10))
	c.DrawWall(image.Pt(20, 10), image.Pt(70, 10))

	c.WaitForQuit()

	fmt.Printf("%T is done\n", c)
}
