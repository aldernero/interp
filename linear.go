package interp

// Lerp calculates the linear interpolation between two values
func Lerp(a float64, b float64, i float64) float64 {
	return a + i*(b-a)
}

// Map calculates the linear interpolation from one range to another
func Map(a float64, b float64, c float64, d float64, i float64) float64 {
	p := (i - a) / (b - a)
	return Lerp(c, d, p)
}

// Linspace creates a slice of linearly distributed values in a range
func Linspace(i float64, j float64, n int, b bool) []float64 {
	var result []float64
	N := float64(n)
	if b {
		N -= 1
	}
	d := (j - i) / N
	for k := 0; k < n; k++ {
		result = append(result, i+float64(k)*d)
	}
	return result
}
