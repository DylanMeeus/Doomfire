package main


import (
    "fmt"
    "syscall/js"
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

var rgbs = []uint32{0x07,0x07,0x07,
            0x1F,0x07,0x07,
            0x2F,0x0F,0x07,
            0x47,0x0F,0x07,
            0x57,0x17,0x07,
            0x67,0x1F,0x07,
            0x77,0x1F,0x07,
            0x8F,0x27,0x07,
            0x9F,0x2F,0x07,
            0xAF,0x3F,0x07,
            0xBF,0x47,0x07,
            0xC7,0x47,0x07,
            0xDF,0x4F,0x07,
            0xDF,0x57,0x07,
            0xDF,0x57,0x07,
            0xD7,0x5F,0x07,
            0xD7,0x5F,0x07,
            0xD7,0x67,0x0F,
            0xCF,0x6F,0x0F,
            0xCF,0x77,0x0F,
            0xCF,0x7F,0x0F,
            0xCF,0x87,0x17,
            0xC7,0x87,0x17,
            0xC7,0x8F,0x17,
            0xC7,0x97,0x1F,
            0xBF,0x9F,0x1F,
            0xBF,0x9F,0x1F,
            0xBF,0xA7,0x27,
            0xBF,0xA7,0x27,
            0xBF,0xAF,0x2F,
            0xB7,0xAF,0x2F,
            0xB7,0xB7,0x2F,
            0xB7,0xB7,0x37,
            0xCF,0xCF,0x6F,
            0xDF,0xDF,0x9F,
            0xEF,0xEF,0xC7,
            0xFF,0xFF,0xFF,
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

    fire := [36][36]uint32{}

    for _,fr := range fire {
        for i,_ := range fr {
            fr[i] = 0
        }
    }

    palette := make([]uint32, 37)
    for i := 0; i < len(rgbs); i+=3 {
        r := rgbs[i]
        g := rgbs[i+1]
        b := rgbs[i+2]
        clor := r
        clor = (clor << 8) + g
        clor = (clor << 8) + b
        palette = append(palette, clor)
    }

    for y,row := range fire {
        for i := 0; i < len(row); i++ {
            if y == 0 {
                fire[y][i] = 0
            } else if y == 1 {
                fire[y][i] = palette[len(palette)-1]
            } else {
                fire[y][i] = 0
            }
        }
    }


    tdiff := 0
    renderFrame = js.NewCallback(func(args []js.Value) {
        tdiff++
        if tdiff % 10 == 0 {
            tdiff = 0 
            updateFire(fire, palette)
            for row,firerow := range fire {
                for col,fireval := range firerow {
                    ctx.Set("fillStyle", fmt.Sprintf("#%06x", int64(fireval) & 0xffffff ))
                    ctx.Call("fillRect", 20 + (col * fw), maxH - float64(row*fh),fw,fh)
                }
            }
        }
        js.Global().Call("requestAnimationFrame", renderFrame)
    })

    defer renderFrame.Release()

    // start
    js.Global().Call("requestAnimationFrame", renderFrame)

    <-done
}

func updateFire(fire [36][36]uint32, palette []uint32){
    // we start from the top!
    spread := func(cury, curx int) uint32 {
        return palette[30]
        //return fire[cury-1][curx]
    }
    for y := 1; y < len(fire); y++ {
        for x,_ := range fire[y] {
            fire[y][x] = spread(y,x)
        }
    }
}


