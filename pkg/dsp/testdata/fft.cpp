#include <complex>
#include <vector>
#include <iostream>

using namespace std;

int prepareFFT(const int N, vector<int>* pids) {
  int n_level;
  {
    auto& i = n_level;
    for( i=0; i<64; ++i )// マジック★ナンバー！
      if( N>>i == 1) break;
  }
  vector<int>& ids = *pids;
  // ID 並び列の計算
  ids.reserve( N );
  {
    ids.push_back( 0 );
    ids.push_back( 1 );
    for( int i=0; i<n_level-1; ++i )
    {
      auto sz = ids.size();
      for_each( ids.begin(), ids.end(), [](int& x){ x*=2; } );
      ids.insert( ids.end(), ids.begin(), ids.end() );
      auto it = ids.begin();
      std::advance( it, sz );
      for_each( it, ids.end(), [](int&x){ x+=1; } );
    }// i
  }
  return n_level;
}

void FFT(const vector<complex<double>>& a, const vector<int>& ids, const int n_level, vector<complex<double>>* pout, bool is_inverse) {
  auto N = a.size();
  auto& F = *pout;
  F.resize(N);
  for (int i = 0; i < N; i++) {
    F[ i ] = a[ids[i]];
  }
  unsigned int po2 = 1;
  for (int i_level = 1; i_level <= n_level; i_level++) {
    po2 <<= 1;
    const int po2m = po2>>1;
    auto w = exp(std::complex<double>(0.0, -2 * M_PI / (double)po2));
    w = is_inverse ? conj(w): w;
    auto ws = complex<double>(1, 0);
    for (int k = 0; k < po2m; k++) {
      for (int j = 0; j < N; j += po2) {
        auto pa = &F[j+k];
        auto pb = &F[j+k+po2m];
        auto wfb = ws**pb;
        *pb = *pa - wfb;
        *pa += wfb;
      }
      ws *= w;
    }
  }
  return;
}

void IFFT(const vector<complex<double>>& a, const vector<int>& ids, const int n_level, vector<complex<double>>* pout) {
  FFT(a, ids, n_level, pout, !0 );
  auto N = a.size();
  for_each( pout->begin(), pout->end(), [N](complex<double>& val){val/=N;} );
  return;
}

int main() {
  int length = 32;

  vector< complex<double> > signal;
  vector< complex<double> > output;

  for (int i = 0; i < length; i++) {
    double d = static_cast<double>(i);
    complex<double> c(i, 0.0);

    signal.push_back(c);
  }

  int nLevel;
  vector<int> index;

  nLevel = prepareFFT(length, &index);
  FFT(signal, index, nLevel, &output, 0);

  for (int i = 0; i < length; i++) {
    printf("%.3f\t%.3f\n", output[i].real(), output[i].imag());
  }

  return 0;
}
