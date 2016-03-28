package main

import (
	"fmt"
	"image"

	"github.com/inkstainblue/dungeon-plotter/canvas"
)

func main() {
	sc := canvas.NewScreenCanvas()

	fmt.Printf("%T is ready\n", sc)

	sc.Draw(image.Pt(0, 0), image.Pt(80, 100))
	sc.Draw(image.Pt(200, 100), image.Pt(700, 100))

	sc.WaitForQuit()

	fmt.Printf("%T is done\n", sc)
}
