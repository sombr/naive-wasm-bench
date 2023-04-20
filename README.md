Naive benchmark of several WASM languages/runtimes.
Based on numeric operations and memory access.

Results (10000000000 iterations of Pi calculation, April 2023):
1. Python (Pyodide): 450000ms
2. Golang: 5500ms
3. C: 1600ms
4. Rust: 1600ms
5. Javascript: 1300ms

Lower - better.

I was surprised (mostly with Go poor performance). Aren't you surprised? You should be.
