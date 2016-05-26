package main

import (
    "fmt"
    "image"
    "image/color"
    "image/gif"
    "io"
    "log"
    "math"
    "net/http"
    "net/url"
    "strconv"
)

var palette = []color.Color{color.White, color.Black}

const (
    whiteIndex = 0
    blackIndex = 1
)

type Params struct {
    Cycles, Size, Nframes, Delay int
    Res, Freq                    float64
}

func getIntParam(qparams url.Values, name string, defaultValue int) int {
    i, err := strconv.Atoi(qparams.Get(name))
    if err == nil {
        return i
    } else {
        return defaultValue
    }
}

func getFloatParam(qparams url.Values, name string, defaultValue float64) float64 {
    i, err := strconv.ParseFloat(qparams.Get(name), 64)
    if err == nil {
        return i
    } else {
        return defaultValue
    }
}

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        qparams := r.URL.Query()
        params := Params{
            Cycles:  getIntParam(qparams, "cycles", 5),
            Size:    getIntParam(qparams, "size", 100),
            Nframes: getIntParam(qparams, "nframes", 5),
            Delay:   getIntParam(qparams, "delay", 5),
            Res:     getFloatParam(qparams, "res", 0.0001),
            Freq:    getFloatParam(qparams, "freq", 1.5),
        }
        lissajous(w, params)
    })
    log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func lissajous(out io.Writer, params Params) {
    fmt.Printf("params: %v\n", params)
    cycles := float64(params.Cycles)
    size := params.Size
    anim := gif.GIF{LoopCount: params.Nframes}
    phase := 0.0

    for i := 0; i < params.Nframes; i++ {
        rect := image.Rect(0, 0, 2*size+1, 2*size+1)
        img := image.NewPaletted(rect, palette)
        for t := 0.0; t < cycles*2*math.Pi; t += params.Res {
            x := math.Sin(t)
            y := math.Sin(t*params.Freq + phase)
            img.SetColorIndex(size+int(x*float64(size)+0.5),
                size+int(y*float64(size)+0.5),
                blackIndex)
        }
        phase += 0.1
        anim.Delay = append(anim.Delay, params.Delay)
        anim.Image = append(anim.Image, img)
    }
    gif.EncodeAll(out, &anim)
}
