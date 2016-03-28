package input

import (
	"fmt"

	"gopkg.in/xkg.v0"
)

type (
	Keyboard struct {
		keys chan int
	}
)

func NewKeyboard() (k Keyboard) {
	k.keys = make(chan int, 100)

	go xkg.StartXGrabber(k.keys)

	go func() {
		for {
			keycode := <-k.keys

			if key, ok := xkg.KeyMap[keycode]; ok {
				fmt.Printf("[%s]\n", key)
			}
		}
	}()

	return
}
