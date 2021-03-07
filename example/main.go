package main

import (
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"math"

	"github.com/moutend/go-fft/pkg/dsp"
)

func main() {
	data, err := ioutil.ReadFile("input.raw")

	if err != nil {
		panic(err)
	}

	signal := []complex128{}

	for i := 0; i < len(data); i += 8 {
		input := math.Float64frombits(
			binary.LittleEndian.Uint64(data[i : i+8]),
		)

		signal = append(signal, complex(input, 0.0))
	}

	spectrum := dsp.FFT(signal[0:32768])

	for _, value := range spectrum {
		fmt.Printf("%.3f\t%.3f\n", real(value), imag(value))
	}
}
