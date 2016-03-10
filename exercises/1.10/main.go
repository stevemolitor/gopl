package main

import (
    "fmt"
    "io"
    "net/http"
    "os"
    "path"
    "time"
)

func main() {
    start := time.Now()
    ch := make(chan string)
    for _, url := range os.Args[1:] {
        go fetch(url, ch) // start a go routine
    }
    for range os.Args[1:] {
        fmt.Println(<-ch) // receive from channel
    }
    fmt.Printf("%.2f elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
    start := time.Now()
    resp, err := http.Get(url)
    if err != nil {
        ch <- fmt.Sprint(err) // send to channel
        return
    }

    fname := path.Base(url)
    out, err := os.Create(fname)
    if err != nil {
        ch <- fmt.Sprintf("while opening %s: %v", fname, err)
    }

    nbytes, err := io.Copy(out, resp.Body)
    out.Close()
    resp.Body.Close()
    if err != nil {
        ch <- fmt.Sprintf("while reading %s: %v", url, err)
        return
    }

    secs := time.Since(start).Seconds()
    ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
