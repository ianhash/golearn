package exercise

import "fmt"

// ErrNegativeSqrt error when sqrt input < 0
type ErrNegativeSqrt float64

// Error error implementor of ErrNegativeSqrt
func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

// Sqrt2 sqrt of x
func Sqrt2(x float64) (float64, error) {

	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	return Sqrt(x), nil
}
