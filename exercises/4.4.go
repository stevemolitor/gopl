package main

import (
    "fmt"
)

func rotate(s []int, i int) []int {
    t := make([]int, len(s))
    for i, j := 0, len(s)-1; i < len(s); i, j = i+1, j-1 {
        t[i] = s[j]
    }
    return t
}

func main() {
    s := []int{1, 2, 3}
    fmt.Println(rotate(s, 0))
}
