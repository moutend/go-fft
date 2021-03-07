package dsp

import (
	"math"
)

// DFT returns the spectrum of the signal. Note that the signal length must be 2^N, e.g. 1024.
func DFT(input []complex128) []complex128 {
	output := make([]complex128, len(input))

	for i := 0; i < len(input); i++ {
		var realSum float64
		var imagSum float64

		for j := 0; j < len(input); j++ {
			theta := 2.0 * math.Pi / float64(len(input)) * float64(i) * float64(j)

			realSum += real(input[j]) * math.Cos(theta)
			imagSum -= real(input[j]) * math.Sin(theta)
		}

		output[i] = complex(realSum, imagSum)
	}

	return output
}
