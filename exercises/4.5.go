package main

import (
    "fmt"
)

func dedup(s []int) []int {
    j := 0
    for i := 0; i < len(s); i++ {
        if i == 0 || s[i] != s[i-1] {
            s[j] = s[i]
            j++
        }
    }
    return s[:j]
}

func main() {
    a := []int{1, 1, 2, 2, 3}
    b := dedup(a)
    fmt.Printf("%v %d\n", dedup(b), len(b))
}
