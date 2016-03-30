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

// TODO: Implement me! Block until complete.
func (sc *PaperCanvas) Clear(rect image.Rectangle) error {
	return nil
}

// TODO: Implement me! Block until complete.
func (sc *PaperCanvas) Draw(a, b Point) error {
	return nil
}

// TODO: Implement me!
func (sc *PaperCanvas) WaitForQuit() {
	return
}
