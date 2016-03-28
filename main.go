package main

import (
	"fmt"
	"image"

	"github.com/inkstainblue/dungeon-plotter/canvas"
)

const (
	maxGridWidth  = 100
	maxGridHeight = 100
)

func main() {
	sc := canvas.NewScreenCanvas(maxGridWidth, maxGridHeight)

	fmt.Printf("%T is ready\n", sc)

	sc.Draw(image.Pt(0, 0), image.Pt(10, 10))
	sc.Draw(image.Pt(20, 10), image.Pt(70, 10))

	sc.WaitForQuit()

	fmt.Printf("%T is done\n", sc)
}
