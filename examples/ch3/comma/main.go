package main

import (
    "fmt"
)

func comma(s string) string {
    n := len(s)
    if n <= 3 {
        return s
    }
    return comma(s[:n-3]) + "," + s[n-3:]
}

func main() {
    fmt.Printf("123: %s\n", comma("123"))
    fmt.Printf("123456 %s\n", comma("123456"))
    fmt.Printf("1234567 %s\n", comma("1234567"))
}
