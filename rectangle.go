package shape

type Rectangle struct {
	height float64
	width  float64
}

func NewRectangle(height, width float64) Rectangle {
	return Rectangle{height, width}
}
