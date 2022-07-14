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

func TestArea(t *testing.T) {
	t.Run("should return 4.0 for rectangle with height and width as 4 and 1", func(t *testing.T) {
		r := Rectangle{4, 1}
		assert.Equal(t, 4.0, r.Area())
	})

	t.Run("should return 8.0 for rectangle with height and width as 4 and 2", func(t *testing.T) {
		r := Rectangle{4, 2}
		assert.Equal(t, 8.0, r.Area())
	})

	t.Run("should return 12.0 for rectangle with height and width as 4 and 3", func(t *testing.T) {
		r := Rectangle{4, 3}
		assert.Equal(t, 12.0, r.Area())
	})

	t.Run("should return 4.0 for rectangle with height and width as 1 and 4", func(t *testing.T) {
		r := Rectangle{1, 4}
		assert.Equal(t, 4.0, r.Area())
	})

	t.Run("should return 8.0 for rectangle with height and width as 2 and 4", func(t *testing.T) {
		r := Rectangle{2, 4}
		assert.Equal(t, 8.0, r.Area())
	})

	t.Run("should return 12.0 for rectangle with height and width as 3 and 4", func(t *testing.T) {
		r := Rectangle{3, 4}
		assert.Equal(t, 12.0, r.Area())
	})

	t.Run("should return 16.0 for rectangle with height and width as 4 and 4", func(t *testing.T) {
		r := Rectangle{4, 4}
		assert.Equal(t, 16.0, r.Area())
	})

	t.Run("should return 10.0 for rectangle with height and width as 5 and 2", func(t *testing.T) {
		r := Rectangle{5, 2}
		assert.Equal(t, 10.0, r.Area())
	})

	t.Run("should return 10.0 for rectangle with height and width as 2 and 5", func(t *testing.T) {
		r := Rectangle{2, 5}
		assert.Equal(t, 10.0, r.Area())
	})
}
