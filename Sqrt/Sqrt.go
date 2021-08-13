package sqrt

import (
	// "log"
	"errors"
	"math"
)

func Sqrt(x float64 /*, error float64*/) (float64, error) {
	// Handle special cases, such as 0 and neg.numbers
	if x == 0 {
		return 1.0, nil
	}
	if x < 0 {
		// log.Fatal("Negative number does not have square root")
		return 0, errors.New("Sqrt.Sqrt: negative number does not have a square root")
	}

	// initialize a guess for Newton's numerical method
	// start with guess and adjust z so, that the diff between x and z**2 is less than a certain limit
	error := 0.001
	z := float64(1)
	for math.Abs(x-z*z) > error {
		z -= (z*z - x) / (2 * z)
	}
	return z, nil
}
