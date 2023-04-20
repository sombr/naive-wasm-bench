#include <math.h>

float benchmarkXS(int r, int c, float *data) {
    int idx = 0;
    float res = 0;
    for (int rr = 0; rr < r; rr += 1) {
        for (int cc = 0; cc < c; cc += 1) {
            data[idx] = sqrt( ((float) rr*cc) / ((float) (rr+cc+1)) );
            res = data[idx];
            idx += 1;
        }
    }

    return res;
}

double piXS(long long limit) {
    // https://iq.opengenus.org/different-ways-to-calculate-pi/
    double d = 1;

    double res = 0;

    for (long long idx = 0; idx < limit; idx++) {
        if (idx % 2 == 0) { // even
            res += 4.0 / d;
        } else {
            res -= 4.0 / d;
        }

        d += 2;
    }

    return res;
}

int main(int argc, char** argv) {
}
