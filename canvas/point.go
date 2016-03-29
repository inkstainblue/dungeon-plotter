package canvas

type (
	Point struct {
		X, Y float64
	}
)

func Pt(x, y float64) Point {
	return Point{
		X: x,
		Y: y,
	}
}

// Add returns the vector p+q.
func (p Point) Add(q Point) Point {
	return Point{
		X: p.X + q.X,
		Y: p.Y + q.Y,
	}
}

// Sub returns the vector p-q.
func (p Point) Sub(q Point) Point {
	return Point{
		X: p.X - q.X,
		Y: p.Y - q.Y,
	}
}
