#include <complex>
#include <vector>
#include <iostream>

using namespace std;

void DFT1(
  const vector<double> a,
  vector<double>& real,
  vector<double>& imag
) {
  int n = a.size();

  for (int i = 0; i < n; i++) {
    double realSum = 0.0;
    double imagSum = 0.0;
    for (int j = 0; j < n; j++) {
      double theta = 2.0 * M_PI / static_cast<double>(n) * static_cast<double>(i) * static_cast<double>(j);
      realSum += a[j] * cos(theta);
      imagSum -= a[j] * sin(theta);
    }
    real.push_back(realSum);
    imag.push_back(imagSum);
  }
  return;
}

void DFT2(
  const vector<double> a,
  vector< complex<double> >& output
) {
  int n = a.size();

  for (int i = 0; i < n; i++) {
    complex<double> sum(0.0, 0.0);

    for (int j = 0; j < n; j++) {
      double theta = 2.0 * M_PI / static_cast<double>(n) * static_cast<double>(i) * static_cast<double>(j);
      sum += a[j] * exp(complex<double>(0.0, -theta));
    }
    output.push_back(sum);
  }
  return;
}

int main() {
  vector<double> a;
  vector<double> real;
  vector<double> imag;

  for (int i = 0; i < 32; i++) {
    a.push_back(static_cast<double>(i));
  }

  DFT1(a, real, imag);
  for (int i = 0; i < 32; i++) {
    // cout << real[i] << "\t" << imag[i] << endl;
    printf("%.3f\t%.3f\n", real[i], imag[i]);
  }
  return 0;
}
