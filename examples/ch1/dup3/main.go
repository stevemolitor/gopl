package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "strings"
)

func main() {
    counts := make(map[string]int)

    for _, fname := range os.Args[1:] {
        data, err := ioutil.ReadFile(fname)
        if err != nil {
            fmt.Fprintf(os.Stderr, "error: %v\n", err)
            continue
        }
        lines := strings.Split(string(data), "\n")
        for _, line := range lines {
            counts[line]++
        }
    }

    for line, n := range counts {
        if n > 1 {
            fmt.Printf("%d\t%s\n", n, line)
        }
    }
}
