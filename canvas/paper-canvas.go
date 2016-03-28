package canvas

import (
	"image"
)

type (
	PaperCanvas struct {
		Canvas
	}
)

// TODO: Implement me!
func NewPaperCanvas(gridWidth, gridHeight int) (sc PaperCanvas) {
	return
}

// TODO: Implement me!
func (sc *PaperCanvas) Clear(rect image.Rectangle) error {
	return nil
}

// TODO: Implement me!
func (sc *PaperCanvas) Draw(a, b image.Point) error {
	return nil
}

// TODO: Implement me!
func (sc *PaperCanvas) WaitForQuit() {
	return
}
