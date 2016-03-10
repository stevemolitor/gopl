package main

import (
    "fmt"
    "io"
    "net/http"
    "os"
    "strings"
)

func main() {
    for _, url := range os.Args[1:] {
        if !strings.HasPrefix(url, "http://") || !strings.HasPrefix(url, "https://") {
            url = "http://" + url
        }
        resp, err := http.Get(url)
        fmt.Fprintf(os.Stderr, "status: %s\n", resp.Status)
        if err != nil {
            fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
            os.Exit(1)
        }
        _, err2 := io.Copy(os.Stdout, resp.Body)
        if err2 != nil {
            fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err2)
        }
        resp.Body.Close()
    }
}
