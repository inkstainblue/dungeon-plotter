package main

import (
	"fmt"

	"github.com/inkstainblue/dungeon-plotter/dungeoneer"
)

const (
	maxGridWidth  = 100
	maxGridHeight = 100
)

func main() {
	d := dungeoneer.New(dungeoneer.DungeoneerOptions{})

	fmt.Println("Dungeoneer is ready")
	fmt.Println("Starting...")

	d.Start()
	defer d.WaitForQuit()
}
