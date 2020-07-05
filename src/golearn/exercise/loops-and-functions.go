package exercise

import "math"

// Sqrt sqrt of x
func Sqrt(x float64) float64 {
	p := x / 2
	for math.Abs(x-p*p) > 0.00001 {
		p -= (p*p - x) / (2 * p)
	}
	return p
}
