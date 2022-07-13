package shape

type compSquare struct {
	rect Rectangle
}

func NewCompSquare(side float64) compSquare {
	return compSquare{NewRectangle(side, side)}
}
