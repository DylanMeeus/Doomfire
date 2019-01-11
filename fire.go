package main


import (
    "fmt"
    "syscall/js"
    "math/rand"
)


var colours = []string{
            "#070707",
            "#1f0707",
            "#2f0f07",
            "#470f07",
            "#571707",
            "#671f07",
            "#771f07",
            "#8f2707",
            "#9f2f07",
            "#af3f07",
            "#bf4707",
            "#c74707",
            
            "#DF4F07",
            "#DF5707",
            "#DF5707",
            
            "#D75F07",
            "#D7670F",
            
            "#cf6f0f",
            "#cf770f",
            "#cf7f0f",
            "#CF8717",
            "#C78717",
            "#C78F17",
            "#C7971F",
            "#BF9F1F",
            "#BF9F1F",
            "#BFA727",
            "#BFA727",
            "#BFAF2F",
            "#B7AF2F",
            "#B7B72F",
            "#B7B737",
            "#CFCF6F",
            "#DFDF9F",
            "#EFEFC7",
            "#FFFFFF", 
}


const (
    fw = 1
    fh = 2
)

type colour struct {
    r,g,b int
}

func main(){
    doc := js.Global().Get("document")
    canvas := js.Global().Get("document").Call("getElementById", "mycanvas")

    bodyW := doc.Get("body").Get("clientWidth").Float()
    bodyH := doc.Get("body").Get("clientHeight").Float()

    canvas.Set("width", bodyW)
    canvas.Set("height", bodyH)

    maxH := bodyH / 2

    ctx := canvas.Call("getContext", "2d")

    done := make(chan struct{}, 0) // something to keep our thread alive

    var renderFrame js.Callback

    fire := [36][400]string{}



    for y,row := range fire {
        for i := 0; i < len(row); i++ {
            fire[y][i] = colours[len(colours)-1-y]
        }
    }


    renderFrame = js.NewCallback(func(args []js.Value) {
        fmt.Println("here")
            canvas.Set("width", bodyW)
            canvas.Set("height", bodyH)
            ctx.Call("beginPath")
            ctx.Call("clearRect",0,0,bodyW,bodyH)
            //ctx.Set("globalCompositeOperation", "destination-atop")
            updateFire(&fire)
            for row,firerow := range fire {
                for col,fireval := range firerow {
                    _ = maxH
                    _ = row
                    _ = col
                    _ = fireval
                    ctx.Set("fillStyle", fireval)
                    ctx.Call("fillRect", 20 + (col * fw), maxH - float64(row*fh),fw,fh)
                }
            }
            ctx.Call("closePath")
        js.Global().Call("requestAnimationFrame", renderFrame)
    })

    defer renderFrame.Release()

    // start
    js.Global().Call("requestAnimationFrame", renderFrame)

    <-done
}

func updateFire(fire *[36][400]string){
    // we start from the top!
    return
    spread := func(cury, curx int) string {
        // the chance of being your own colour or the one from below
        if rand.Float32() > 0.2 {
            return fire[cury-1][curx]
        }
        return fire[cury][curx]
    }
    for y := 1; y < len(fire); y++ {
        for x,_ := range fire[y] {
            fire[y][x] = spread(y,x)
        }
    }
}


