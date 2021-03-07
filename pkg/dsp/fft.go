package dsp

import (
	"math"
)

func getPow(x int) int {
	var pow int

	// The 64 is enough because the maximum integer is 64 bit.
	for pow = 0; pow < 64; pow++ {
		if x>>pow == 1 {
			break
		}
	}

	return pow
}

func getIndex(length int) []int {
	pow := getPow(length)
	index := []int{}

	index = append(index, 0)
	index = append(index, 1)

	for i := 0; i < pow-1; i++ {
		indexLength := len(index)

		for j := range index {
			index[j] *= 2
		}
		for j := 0; j < indexLength; j++ {
			index = append(index, index[j])
		}
		for j := indexLength; j < len(index); j++ {
			index[j] += 1
		}
	}

	return index
}

// FFT returns the spectrum of the signal. Note that the signal length must be 2^N, e.g. 1024.
func FFT(signal []complex128) []complex128 {
	length := len(signal)
	index := getIndex(length)
	pow := getPow(length)
	output := make([]complex128, length)

	for i := 0; i < length; i++ {
		output[i] = signal[index[i]]
	}

	po2 := 1

	for i := 1; i <= pow; i++ {
		po2 = po2 << 1

		po2m := po2 >> 1
		theta := 2.0 * math.Pi / float64(po2)
		w := complex(math.Cos(theta), -math.Sin(theta))
		ws := complex(1.0, 0.0)

		for k := 0; k < po2m; k++ {
			for j := 0; j < length; j += po2 {
				wfb := ws * output[j+k+po2m]

				output[j+k+po2m] = output[j+k] - wfb
				output[j+k] += wfb
			}

			ws *= w
		}
	}

	return output
}
