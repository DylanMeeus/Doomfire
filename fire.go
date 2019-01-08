package main


import (
    "fmt"
    "syscall/js"
)

func main(){
    fmt.Println("Hello world!")

    doc := js.Global().Get("document")
    canvas := js.Global().Get("document").Call("getElementById", "mycanvas")

    bodyW := doc.Get("body").Get("clientWidth").Float()
    bodyH := doc.Get("body").Get("clientHeight").Float()

    canvas.Set("width", bodyW)
    canvas.Set("height", bodyH)

    ctx := canvas.Call("getContext", "2d")
    _ = ctx

    done := make(chan struct{}, 0) // something to keep our thread alive

    var renderFrame js.Callback

    renderFrame = js.NewCallback(func(args []js.Value) {
        fmt.Println("output!")
        js.Global().Call("requestAnimationFrame", renderFrame)
    })

    defer renderFrame.Release()

    // start
    js.Global().Call("requestAnimationFrame", renderFrame)

    <-done
    

}

