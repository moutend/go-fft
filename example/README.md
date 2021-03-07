Get the spectrum of the audio signal
====================================

This example shows hot to use `dsp.FFT()` function.

## Required Tool

- [go](https://golang.org/dl/)
- [sox](http://sox.sourceforge.net)

## Usage

Open the terminal app and run the following steps:

```console
$ sox music.wav -t f64 input.raw
$ go run main.go
```

You'll see the tab separated real and imag values.

Note that I assume that the sample rate of music.wav is 44.1 kHz and the playback time is longer than 0.1 sec.
