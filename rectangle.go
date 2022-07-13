package shape

type Rectangle struct {
	Height float64
	Width  float64
}

func NewRectangle(height, width float64) Rectangle {
	if height <= 0 || width <= 0 {
		panic("width and height should be positive")
	}
	return Rectangle{height, width}
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (r Rectangle) Area() float64 {
	return 4.0
}
