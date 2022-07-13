package shape

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCompSquare(t *testing.T) {
	t.Run("should be able to initialize square with side", func(t *testing.T) {
		assert.NotPanics(t, func() {
			NewCompSquare(5)
		})
	})

	t.Run("new square should return square", func(t *testing.T) {
		assert.IsType(t, compSquare{}, NewCompSquare(5))
	})
}

func TestNewCompPerimeter(t *testing.T) {
	t.Run("should return 4.0 for square with side 1", func(t *testing.T) {
		assert.Equal(t, 4.0, NewCompSquare(1).Perimeter())
	})

	t.Run("should return 8.0 for square with side 2", func(t *testing.T) {
		assert.Equal(t, 8.0, NewCompSquare(2).Perimeter())
	})

	t.Run("should return 12.0 for square with side 3", func(t *testing.T) {
		assert.Equal(t, 12.0, NewCompSquare(3).Perimeter())
	})

	t.Run("should return 40.0 for square with side 10", func(t *testing.T) {
		assert.Equal(t, 40.0, NewCompSquare(10).Perimeter())
	})
}

func TestNewCompArea(t *testing.T) {
	t.Run("should return 1.0 for square with side 1", func(t *testing.T) {
		assert.Equal(t, 1.0, NewCompSquare(1).Area())
	})

	t.Run("should return 4.0 for square with side 2", func(t *testing.T) {
		assert.Equal(t, 4.0, NewCompSquare(2).Area())
	})

	t.Run("should return 9.0 for square with side 3", func(t *testing.T) {
		assert.Equal(t, 9.0, NewCompSquare(3).Area())
	})

	t.Run("should return 100.0 for square with side 10", func(t *testing.T) {
		assert.Equal(t, 100.0, NewCompSquare(10).Area())
	})
}
