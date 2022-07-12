package shape

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewRectangle(t *testing.T) {
	t.Run("should be able to initialize with height and width", func(t *testing.T) {
		assert.NotPanics(t, func() {
			NewRectangle(4, 5)
		})
	})

	t.Run("new rectangle should return rectangle", func(t *testing.T) {
		assert.IsType(t, Rectangle{}, NewRectangle(4, 5))
	})
}

func TestPerimeter(t *testing.T) {
	t.Run("should return 18.0 for rectangle with height and width as 4 and 5", func(t *testing.T) {
		assert.Equal(t, Perimeter(4, 5), 18.0)
	})

	t.Run("should return 20.0 for rectangle with height and width as 5 and 5", func(t *testing.T) {
		assert.Equal(t, Perimeter(5, 5), 20.0)
	})
}
