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

	c.DrawWall(canvas.Pt(0, 0), canvas.Pt(0, 99))
	c.DrawWall(canvas.Pt(0, 0), canvas.Pt(99, 0))
	c.DrawWall(canvas.Pt(99, 0), canvas.Pt(99, 99))
	c.DrawWall(canvas.Pt(0, 99), canvas.Pt(99, 99))

	c.DrawPath(canvas.Pt(50, 50), canvas.Pt(50, 50))

	c.DrawPath(canvas.Pt(25, 25), canvas.Pt(75, 25))
	c.DrawPath(canvas.Pt(75, 25), canvas.Pt(75, 75))
	c.DrawPath(canvas.Pt(75, 75), canvas.Pt(25, 75))
	c.DrawPath(canvas.Pt(25, 75), canvas.Pt(25, 25))

	c.DrawPath(canvas.Pt(50, 25), canvas.Pt(75, 50))
	c.DrawPath(canvas.Pt(75, 50), canvas.Pt(50, 75))
	c.DrawPath(canvas.Pt(50, 75), canvas.Pt(25, 50))
	c.DrawPath(canvas.Pt(25, 50), canvas.Pt(50, 25))

	fmt.Println("Waiting for quit")

	c.WaitForQuit()

	fmt.Printf("%T is done\n", c)
}
