package dsp

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
)

func TestGetIndex(t *testing.T) {
	expected := []int{0, 16, 8, 24, 4, 20, 12, 28, 2, 18, 10, 26, 6, 22, 14, 30, 1, 17, 9, 25, 5, 21, 13, 29, 3, 19, 11, 27, 7, 23, 15, 31}
	actual := getIndex(32)

	for i := range expected {
		if expected[i] != actual[i] {
			t.Fatalf("expected[%d]=%v, actual[%d]=%v\n", i, expected[i], i, actual[i])
		}
	}
}

func loadTSV(path string) ([]string, error) {
	data, err := ioutil.ReadFile(path)

	if err != nil {
		return nil, err
	}

	tsv := strings.Split(string(data), "\n")

	return tsv, nil
}

func TestFFT(t *testing.T) {
	paths := []string{}

	err := filepath.Walk("testdata", func(path string, info fs.FileInfo, err error) error {
		if strings.HasSuffix(path, ".tsv") {
			paths = append(paths, path)
		}

		return nil
	})

	if err != nil {
		t.Fatal(err)
	}
	for _, path := range paths {
		tsv, err := loadTSV(path)

		if err != nil {
			t.Fatal(err)
		}

		length, err := strconv.Atoi(
			strings.TrimSuffix(filepath.Base(path), ".tsv"),
		)

		if err != nil {
			t.Fatal(err)
		}

		signal := make([]complex128, length)

		for i := 0; i < length; i++ {
			signal[i] = complex(float64(i), 0.0)
		}

		output := FFT(signal)

		for i := 0; i < length; i++ {
			expected := tsv[i]
			actual := fmt.Sprintf("%.3f\t%.3f", real(output[i]), imag(output[i]))

			if expected != actual {
				t.Fatalf("%d-th value is not match\n", i)
				t.Fatalf("\tExpected=%s\n", expected)
				t.Fatalf("\tActual=%s\n", actual)
			}
		}
	}
}
