package shape

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewRectangle(t *testing.T) {
	t.Run("should be able to initialize rectangle with height and width", func(t *testing.T) {
		assert.NotPanics(t, func() {
			NewRectangle(4, 5)
		})
	})

	t.Run("width and height should be positive", func(t *testing.T) {
		assert.Panics(t, func() {
			NewRectangle(-1, 0)
		})
	})

	t.Run("new rectangle should return rectangle", func(t *testing.T) {
		assert.IsType(t, Rectangle{}, NewRectangle(4, 5))
	})
}

func TestPerimeter(t *testing.T) {
	t.Run("should return 18.0 for rectangle with height and width as 4 and 5", func(t *testing.T) {
		r := Rectangle{4, 5}
		assert.Equal(t, 18.0, r.Perimeter())
	})
	t.Run("should return 18.0 for rectangle with height and width as 5 and 4", func(t *testing.T) {
		r := Rectangle{5, 4}
		assert.Equal(t, 18.0, r.Perimeter())
	})

	t.Run("should return 20.0 for rectangle with height and width as 5 and 5", func(t *testing.T) {
		r := Rectangle{5, 5}
		assert.Equal(t, 20.0, r.Perimeter())
	})
}
