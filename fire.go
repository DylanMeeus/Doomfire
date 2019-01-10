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
    fw = 10
    fh = 10
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

    fire := [36][36]string{}



    for y,row := range fire {
        for i := 0; i < len(row); i++ {
            if y == 0 {
                fire[y][i] = colours[0]
            } else if y == 1 {
                fire[y][i] = colours[len(colours)-1]
            } else {
                fire[y][i] = colours[0]
            }
        }
    }


    tdiff := 0
    renderFrame = js.NewCallback(func(args []js.Value) {
        tdiff++
        if tdiff % 20 == 0 {
            fmt.Println(bodyW)
            fmt.Println(maxH)
            canvas.Set("width", bodyW)
            canvas.Set("height", bodyH)
            ctx.Call("beginPath")
            ctx.Call("clearRect",0,0,bodyW,bodyH)
            //ctx.Set("globalCompositeOperation", "destination-atop")
            tdiff = 0 
            fmt.Println("updating")
            updateFire(&fire)
            for row,firerow := range fire {
                for col,fireval := range firerow {
                    ctx.Set("fillStyle", fireval)
                    ctx.Set("strokeStyle", fireval)
                    ctx.Call("fillRect", 20 + (col * fw), maxH - float64(row*fh),fw,fh)
                }
            }
            ctx.Call("closePath")
        }
        js.Global().Call("requestAnimationFrame", renderFrame)
    })

    defer renderFrame.Release()

    // start
    js.Global().Call("requestAnimationFrame", renderFrame)

    <-done
}

func updateFire(fire *[36][36]string){
    // we start from the top!
    spread := func(cury, curx int) string {
        // the chance of being your own colour or the one from below
        if rand.Float32() > 0.9 {
            return colours[cury-1]
        }
        if colours[cury] == fire[cury][curx] {
            return colours[cury]
        }
        return fire[cury][curx]
    }
    for y := 1; y < len(fire); y++ {
        for x,_ := range fire[y] {
            fire[y][x] = spread(y,x)
        }
    }
}


