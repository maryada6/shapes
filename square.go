package shape

type Square struct {
	rect Rectangle
}

func NewSquare(side float64) Square {
	rectangle := NewRectangle(side, side)
	if rectangle.height != rectangle.width {
		panic("square should have equal height and width")
	}
	return Square{rectangle}
}

func (square Square) Area() float64 {
	return square.rect.Area()
}

func (square Square) Perimeter() float64 {
	return square.rect.Perimeter()
}
