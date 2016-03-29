package main

import (
	"fmt"

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

	c.DrawWall(canvas.Pt(0, 0), canvas.Pt(10, 10))

	c.DrawPath(canvas.Pt(20, 10), canvas.Pt(70, 10))

	c.DrawPath(canvas.Pt(20, 20), canvas.Pt(30, 20))
	c.DrawPath(canvas.Pt(30, 20), canvas.Pt(35, 20))
	c.DrawPath(canvas.Pt(35, 20), canvas.Pt(70, 20))

	c.DrawPath(canvas.Pt(40, 30), canvas.Pt(70, 50))

	c.WaitForQuit()

	fmt.Printf("%T is done\n", c)
}
