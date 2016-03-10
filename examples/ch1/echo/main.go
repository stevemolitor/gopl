package main

import (
    "fmt"
    "os"
)

func main() {
    s, sep := "", ""
    for _, arg := range os.Args[1:] {
        s += sep + arg
        sep = " "
    }
    fmt.Println(s)

    // s := strings.Join(os.Args[1:], " ")
    // fmt.Println(s)
}
