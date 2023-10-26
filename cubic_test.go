package interp

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCubicSpline(t *testing.T) {
	x := []float64{0, 1, 2, 3.5, 7}
	y := []float64{0, 0.2, 0.1, 0.45, 0.3}
	cs, err := NewCubicSpline(x, y)
	assert.NoError(t, err)
	assert.Equal(t, 1, cs.Interval(1))
	assert.Equal(t, 0.0, cs.Eval(-1))
	assert.Equal(t, 0.3, cs.Eval(10))
	for i, val := range x {
		assert.Equal(t, y[i], cs.Eval(val))
	}
	x = []float64{0, 1, 2, 0.5, 7}
	cs, err = NewCubicSpline(x, y)
	assert.Error(t, err)
}
