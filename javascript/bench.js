
function benchmarkJS(r, c) {
    let data = new Float32Array(r*c);

    console.time("benchmark JS");
    let idx = 0;
    let lastval = 0;
    for (let rr = 0; rr < r; rr += 1) {
        for (let cc = 0; cc < c; cc += 1) {
            data[idx] = Math.sqrt(rr * cc / ( rr + cc + 1 ));
            lastval = data[idx];
            idx += 1;
        }
    }
    console.timeEnd("benchmark JS");

    return lastval;
}

function benchpiJS(limit) {
    let d = 1
    let res = 0

    console.time("benchmark PI JS");
    for (let idx = 0; idx < limit; idx += 1) {
        if (idx % 2 == 0) {
            res += 4.0 / d;
        } else {
            res -= 4.0 / d;
        }
        d += 2.0
    }
    console.timeEnd("benchmark PI JS");

    return res;
}
