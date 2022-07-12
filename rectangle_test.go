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
