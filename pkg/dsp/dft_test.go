package dsp

import (
	"math"
	"testing"
)

func TestDFT(t *testing.T) {
	for shift := 5; shift < 11; shift++ {
		length := 2 << shift
		signal := make([]complex128, length)

		for i := 0; i < length; i++ {
			signal[i] = complex(float64(i), 0.0)
		}

		output1 := DFT(signal)
		output2 := FFT(signal)

		for i := 0; i < length; i++ {
			r1 := real(output1[i])
			r2 := real(output2[i])
			i1 := imag(output1[i])
			i2 := imag(output2[i])

			threshold := 0.000001
			isInvalidReal := math.Abs(r1-r2) > threshold
			isInvalidImag := math.Abs(i1-i2) > threshold

			if isInvalidReal {
				t.Logf("Expected Real Value: %v\n", r1)
				t.Logf("Actual Real Value: %v\n", r2)
			}
			if isInvalidImag {
				t.Logf("Expected Imag Value: %v\n", r1)
				t.Logf("Actual Imag Value: %v\n", r2)
			}
			if isInvalidReal || isInvalidImag {
				t.Fatalf("length: %d, %d-th value is not match\n", length, i)
			}
		}
	}
}
