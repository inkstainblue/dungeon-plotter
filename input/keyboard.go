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

func (k *Keyboard) WaitForInput() (code int) {
	return <-k.keys
}
