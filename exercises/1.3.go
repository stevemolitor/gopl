package main

import (
    "fmt"
    "os"
    "strings"
    "time"
)

func slowEcho() {
    s, sep := "", ""
    for _, arg := range os.Args[1:] {
        s += sep + arg
        sep = " "
    }
    fmt.Println(s)
}

func fastEcho() {
    fmt.Println(strings.Join(os.Args, " "))
}

func main() {
    start := time.Now()
    slowEcho()
    elapsed := time.Since(start).Seconds()
    fmt.Printf("slowEcho: %f\n", elapsed)

    start = time.Now()
    fastEcho()
    elapsed = time.Since(start).Seconds()
    fmt.Printf("fastEcho: %f\n ", elapsed)
}
