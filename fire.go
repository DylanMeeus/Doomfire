package main


import (
    "syscall/js"
)

func main(){
    doc := js.Global().Get("document")
    canvas := js.Global().Get("document").Call("getElementById", "mycanvas")

    bodyW := doc.Get("body").Get("clientWidth").Float()
    bodyH := doc.Get("body").Get("clientHeight").Float()

    canvas.Set("width", bodyW)
    canvas.Set("height", bodyH)

    ctx := canvas.Call("getContext", "2d")

    done := make(chan struct{}, 0) // something to keep our thread alive

    var renderFrame js.Callback

    ctx.Set("globalAlpha", 0.5)
    renderFrame = js.NewCallback(func(args []js.Value) {
        ctx.Call("beginPath")
        ctx.Set("fillStyle", "#ffff00")
        ctx.Set("strokeStyle", "#ffff00")
        ctx.Call("rect", 20,20,150,150)
        ctx.Call("fill")
        js.Global().Call("requestAnimationFrame", renderFrame)
    })

    defer renderFrame.Release()

    // start
    js.Global().Call("requestAnimationFrame", renderFrame)

    <-done
}


