#!/bin/bash
emcc main.c -O3 -lm -o main.js -s WASM=1 -s EXPORTED_FUNCTIONS='["_benchmarkXS", "_piXS", "_main", "_malloc", "_free"]' -s EXPORTED_RUNTIME_METHODS='["cwrap"]' -s ALLOW_MEMORY_GROWTH=1 -s ABORTING_MALLOC=0
