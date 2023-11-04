package interp

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBisect(t *testing.T) {
	x := []float64{0, 1, 2, 3, 4, 5}
	assert.Equal(t, -1, Bisect(x, -1))
	assert.Equal(t, 0, Bisect(x, 0))
	assert.Equal(t, 0, Bisect(x, 0.5))
	assert.Equal(t, 1, Bisect(x, 1))
	assert.Equal(t, 1, Bisect(x, 1.5))
	assert.Equal(t, 2, Bisect(x, 2))
	assert.Equal(t, 2, Bisect(x, 2.5))
	assert.Equal(t, 3, Bisect(x, 3))
	assert.Equal(t, 3, Bisect(x, 3.5))
	assert.Equal(t, 4, Bisect(x, 4))
	assert.Equal(t, 4, Bisect(x, 4.5))
	assert.Equal(t, 5, Bisect(x, 5))
	assert.Equal(t, 5, Bisect(x, 6))
}
