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
