package main

import (
    "fmt"
    "log"
    "math"
    "net/http"
    "strconv"
)

const (
    defaultWidth, defaultHeight = 600, 320
    cells                       = 100
    xyrange                     = 30.0
    xyscale                     = defaultWidth / 2 / xyrange
    zscale                      = defaultHeight * 0.4
    angle                       = math.Pi / 6
    defaultRatio                = 0.04
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
    http.HandleFunc("/", serveSvg)
    log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func serveSvg(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "image/svg+xml")
    w.Header().Set("Vary", "Accept-Encoding")

    params := r.URL.Query()

    height, err := strconv.Atoi(params.Get("height"))
    if err != nil {
        height = defaultHeight
    }

    width, err := strconv.Atoi(params.Get("width"))
    if err != nil {
        width = defaultWidth
    }

    fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/200/svg' "+
        "style='stroke: grey; fill: white; stroke-width: 0.7' "+
        "width='%d' height='%d'>", width, height)
    for i := 0; i < cells; i++ {
        for j := 0; j < cells; j++ {
            ax, ay := corner(i+1, j)
            bx, by := corner(i, j)
            cx, cy := corner(i, j+1)
            dx, dy := corner(i+1, j+1)
            fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
                ax, ay, bx, by, cx, cy, dx, dy)
        }
    }
    fmt.Fprintln(w, "</svg>")
}

func corner(i, j int) (float64, float64) {
    x := xyrange * (float64(i)/cells - 0.5)
    y := xyrange * (float64(j)/cells - 0.5)

    z := f(x, y)

    sx := defaultWidth/2 + (x-y)*cos30*xyscale
    sy := defaultHeight/2 + (x+y)*sin30*xyscale - z*zscale
    return sx, sy
}

func f(x, y float64) float64 {
    r := math.Hypot(x, y)
    r = math.Sin(r) / r
    if math.IsInf(r, 0) {
        return defaultRatio
    }
    return r
}
