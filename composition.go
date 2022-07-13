package shape

type compSquare struct {
	rect Rectangle
}

func NewCompSquare(side float64) compSquare {
	return compSquare{NewRectangle(side, side)}
}

func (square compSquare) Area() float64 {
	return square.rect.Area()
}

func (square compSquare) Perimeter() float64 {
	return square.rect.Perimeter()
}
