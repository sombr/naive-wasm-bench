package main

import "syscall/js"
import "math"
import "fmt"

func makeBench() js.Func {
    return js.FuncOf(func (this js.Value, args []js.Value) any {
        ri := args[0].Int()
        ci := args[1].Int()

        console := js.Global().Get("console")

        data := make([]float32, ri*ci, ri*ci)

        idx := 0
        var lastval float32 = 0
        console.Call("time", "benchmark Go")
        for rr := 0; rr < ri; rr += 1 {
            for cc := 0; cc < ci; cc += 1 {
                data[idx] = float32(math.Sqrt( float64(rr * cc) / float64(rr+cc+1) ))
                lastval = data[idx]
                idx += 1
            }
        }
        console.Call("timeEnd", "benchmark Go")

        return js.ValueOf(lastval)
    })
}

func makeBenchPI() js.Func {
    return js.FuncOf(func (this js.Value, args []js.Value) any {
        limit := args[0].Int()

        console := js.Global().Get("console")

        console.Call("time", "benchmark PI Go")
        var d float64 = 1
        var res float64 = 0
        for idx := 0; idx < limit; idx += 1 {
            if idx % 2 == 0 {
                res += 4.0 / d
            } else {
                res -= 4.0 / d
            }
            d += 2.0
        }

        console.Call("timeEnd", "benchmark PI Go")

        return js.ValueOf(res)
    })
}

func main() {
    fmt.Println("Go Web Assembly")
    js.Global().Set("benchmarkGO", makeBench())
    js.Global().Set("benchpiGO", makeBenchPI())

    <-make(chan bool)
}
