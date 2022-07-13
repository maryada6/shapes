package shape

type compSquare struct {
	rect Rectangle
}

func NewCompSquare(side float64) compSquare {
	rectangle := NewRectangle(side, side)
	if rectangle.height != rectangle.width {
		panic("square should have equal height and width")
	}
	return compSquare{rectangle}
}

func (square compSquare) Area() float64 {
	return square.rect.Area()
}

func (square compSquare) Perimeter() float64 {
	return square.rect.Perimeter()
}
