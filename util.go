package interp

// TDMA implements the Thomas algorithm for solving tridiagonal systems of equations
// a is the lower diagonal, b is the main diagonal, c is the upper diagonal, and d is the right hand side
func TDMA(a, b, c, d []float64) []float64 {
	n := len(d)
	x := make([]float64, n)
	// elimination step
	for i := 1; i < n; i++ {
		q := a[i] / b[i-1]
		b[i] = b[i] - q*c[i-1]
		d[i] = d[i] - q*d[i-1]
	}
	// back substitution step
	q := d[n-1] / b[n-1]
	x[n-1] = q
	// forward substitution step
	for i := n - 2; i >= 0; i-- {
		q = (d[i] - c[i]*q) / b[i]
		x[i] = q
	}
	return x
}

func Bisect(x []float64, val float64) int {
	n := len(x)
	if n == 0 {
		return 0
	}
	if val < x[0] {
		return -1
	}
	if val > x[n-1] {
		return n - 1
	}
	l := 0
	r := n - 1
	for l < r {
		if val == x[l] {
			return l
		}
		if val == x[r] {
			return r
		}
		m := (l + r) / 2
		if val == x[m] {
			return m
		} else if val < x[m] {
			r = m
		} else {
			l = m + 1
		}
	}
	return l - 1
}
