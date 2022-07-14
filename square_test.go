package shape

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSquare(t *testing.T) {
	t.Run("should be able to initialize square with side", func(t *testing.T) {
		assert.NotPanics(t, func() {
			NewSquare(5)
		})
	})

	t.Run("new square should return square", func(t *testing.T) {
		assert.IsType(t, Square{}, NewSquare(4))
	})
}

func TestSquarePerimeter(t *testing.T) {
	t.Run("should return 4.0 for square with side 1", func(t *testing.T) {
		assert.Equal(t, 4.0, NewSquare(1).Perimeter())
	})

	t.Run("should return 8.0 for square with side 2", func(t *testing.T) {
		assert.Equal(t, 8.0, NewSquare(2).Perimeter())
	})

	t.Run("should return 12.0 for square with side 3", func(t *testing.T) {
		assert.Equal(t, 12.0, NewSquare(3).Perimeter())
	})

	t.Run("should return 40.0 for square with side 10", func(t *testing.T) {
		assert.Equal(t, 40.0, NewSquare(10).Perimeter())
	})
}

func TestSquareArea(t *testing.T) {
	t.Run("should return 1.0 for square with side 1", func(t *testing.T) {
		assert.Equal(t, 1.0, NewSquare(1).Area())
	})

	t.Run("should return 4.0 for square with side 2", func(t *testing.T) {
		assert.Equal(t, 4.0, NewSquare(2).Area())
	})

	t.Run("should return 9.0 for square with side 3", func(t *testing.T) {
		assert.Equal(t, 9.0, NewSquare(3).Area())
	})

	t.Run("should return 100.0 for square with side 10", func(t *testing.T) {
		assert.Equal(t, 100.0, NewSquare(10).Area())
	})
}
