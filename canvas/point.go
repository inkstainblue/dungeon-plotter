package canvas

type (
	Point struct {
		X, Y float64
	}
)

// Pt is shorthand for Point{X, Y}.
func Pt(x, y float64) Point {
	return Point{
		X: x,
		Y: y,
	}
}

// Add returns the vector p+q.
func (p Point) Add(q Point) Point {
	return Pt(p.X+q.X, p.Y+q.Y)
}

// Sub returns the vector p-q.
func (p Point) Sub(q Point) Point {
	return Pt(p.X-q.X, p.Y-q.Y)
}

// Mul returns the vector p*k.
func (p Point) Mul(k float64) Point {
	return Pt(p.X*k, p.Y*k)
}

// Div returns the vector p/k.
func (p Point) Div(k float64) Point {
	return Pt(p.X/k, p.Y/k)
}
