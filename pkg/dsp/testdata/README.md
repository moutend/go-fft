# Generating spectrum TSV files

```console
$ c++ -std=c++11 fft.cpp
$ ./a.out > spectrum/32.tsv
```

The default length is 32. Edit the `fft.cpp:69` by need.

FYI, I implemented the `fft.cpp` based on the this article:

- [高速フーリエ変換(FFT)をおじさんもC++で作ってみたよ - nursの日記](https://nurs.hatenablog.com/entry/20130617/1371483633)
