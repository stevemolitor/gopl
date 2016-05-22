package main

import (
    "crypto/sha256"
    "fmt"
    "os"
)

var pc [256]byte

func init() {
    for i := range pc {
        pc[i] = pc[i/2] + byte(i&1)
    }
}

func bitsDiff(a, b [32]byte) int {
    count := 0
    for i := 0; i < 32; i++ {
        comp := a[i] ^ b[i]
        count += int(pc[comp])
    }
    return count
}

const usage = `Exercie 4.1 computes the number of bits different between to sha256 values.

Usage:
    ./4.2 [STRING1] [STRING2]
`

func main() {
    if len(os.Args) < 3 {
        fmt.Fprintln(os.Stderr, usage)
        os.Exit(1)
    }

    a := sha256.Sum256([]byte(os.Args[1]))
    b := sha256.Sum256([]byte(os.Args[2]))

    fmt.Printf("num bits different: %d", bitsDiff(a, b))
}
