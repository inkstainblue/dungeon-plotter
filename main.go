package main

import (
	"fmt"

	"github.com/inkstainblue/dungeon-plotter/canvas"
)

func main() {
	sc := canvas.NewScreenCanvas()

	fmt.Printf("%T is ready\n", sc)

	sc.WaitForQuit()

	fmt.Printf("%T is done\n", sc)
}
