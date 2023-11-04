package interp

import "fmt"

type coefs [4]float64

// Cubic spline interpolation
type CubicSpline struct {
	x    []float64
	y    []float64
	s    []func(float64) float64 // spline functions
	c    []coefs                 // coefficients
	n    int
	minX float64
	maxX float64
}

// NewCubicSpline creates a new cubic spline interpolation
func NewCubicSpline(x []float64, y []float64) (CubicSpline, error) {
	var cs CubicSpline
	if len(x) != len(y) {
		return cs, fmt.Errorf("x and y must be the same length")
	}
	n := len(x)
	for i := 0; i < n-1; i++ {
		if x[i] >= x[i+1] {
			return cs, fmt.Errorf("x values must be strictly increasing")
		}
	}
	cs.x = x
	cs.y = y
	cs.s = make([]func(float64) float64, n-1)
	cs.c = make([]coefs, n-1)
	var lower, middle, upper, right []float64
	lower = make([]float64, n)
	middle = make([]float64, n)
	upper = make([]float64, n)
	right = make([]float64, n)
	middle[0] = 2
	middle[n-1] = 2
	for i := 1; i <= n-2; i++ {
		h0 := x[i] - x[i-1]
		h1 := x[i+1] - x[i]
		dy0 := y[i] - y[i-1]
		dy1 := y[i+1] - y[i]
		upper[i] = h1 / (h0 + h1)
		lower[i] = 1 - upper[i]
		middle[i] = 2
		right[i] = (6 / (h0 + h1)) * (dy1/h1 - dy0/h0)
	}
	// solve the tridiagonal system
	M := TDMA(lower, middle, upper, right)
	// calculate the spline functions
	// s[i](x) = a + b(x-xi) + c(x-xi)^2 + d(x-xi)^3
	for i := 0; i < n-1; i++ {
		h := x[i+1] - x[i]
		a := y[i]
		b := (y[i+1]-y[i])/h - h*(M[i+1]+2*M[i])/6
		c := M[i] / 2
		d := (M[i+1] - M[i]) / (6 * h)
		cs.c[i] = coefs{a, b, c, d}
	}
	cs.n = n
	cs.minX = x[0]
	cs.maxX = x[n-1]
	return cs, nil
}

// Eval evaluates the cubic spline at a point
func (cs CubicSpline) Eval(x float64) float64 {
	i := cs.Interval(x)
	if i < 0 {
		return cs.y[0]
	} else if i >= cs.n-1 {
		return cs.y[cs.n-1]
	}
	coefs := cs.c[i]
	dx := x - cs.x[i]
	return coefs[0] + dx*(coefs[1]+dx*(coefs[2]+dx*coefs[3])) // Horner's method
}

func (cs CubicSpline) Interval(x float64) int {
	return Bisect(cs.x, x)
}

func (cs CubicSpline) GetMinX() float64 {
	return cs.minX
}

func (cs CubicSpline) GetMaxX() float64 {
	return cs.maxX
}

func (cs CubicSpline) GetMinMaxX() (float64, float64) {
	return cs.minX, cs.maxX
}
