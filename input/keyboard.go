package input

import (
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

	return
}

func (k *Keyboard) WaitForInput() (code int, label string) {
	code = <-k.keys
	label = xkg.KeyMap[code]

	return
}
