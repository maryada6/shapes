package shape

type Rectangle struct {
	height float64
	width  float64
}

func NewRectangle(height, width float64) Rectangle {
	if height <= 0 || width <= 0 {
		panic("width and height should be positive")
	}
	return Rectangle{height, width}
}

func Perimeter(height float64, width float64) float64 {
	return width + width + height + height
}
