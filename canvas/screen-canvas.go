package canvas

import (
	"image"
)

type (
	ScreenCanvas struct {
		Canvas
	}
)

func (sc *ScreenCanvas) Clear(rect image.Rectangle) error {
	return nil
}

func (sc *ScreenCanvas) Draw(a, b image.Point) error {
	return nil
}
