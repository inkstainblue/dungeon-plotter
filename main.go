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
	k := input.NewKeyboard()

	fmt.Printf("%T is ready\n", k)

	sc := canvas.NewScreenCanvas(maxGridWidth, maxGridHeight)
	c := controller.New(&sc)

	fmt.Printf("%T is ready\n", c)

	c.DrawWall(image.Pt(0, 0), image.Pt(10, 10))
	c.DrawWall(image.Pt(20, 10), image.Pt(70, 10))

	c.WaitForQuit()

	fmt.Printf("%T is done\n", c)
}
