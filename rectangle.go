package shape

type Shape interface {
	Perimeter() float64
	Area() float64
}

type Rectangle struct {
	height float64
	width  float64
}

type Square struct {
	side float64
}

func NewRectangle(height, width float64) Rectangle {
	if height <= 0 || width <= 0 {
		panic("width and height should be positive")
	}
	return Rectangle{height, width}
}

func NewSquare(side float64) Square {
	return Square{side}
}

func (rect Rectangle) Perimeter() float64 {
	return 2 * (rect.width + rect.height)
}

func (rect Rectangle) Area() float64 {
	return rect.width * rect.height
}

func (square Square) Perimeter() float64 {
	return square.side * 4
}
