package main

import (
    "bytes"
    "fmt"
)

func comma(s string) string {
    var buf bytes.Buffer
    n := len(s)

    buf.WriteString(s[:n%3])
    if n%3 != 0 && n > 3 {
        buf.WriteString(",")
    }

    for i := n % 3; i < n; i += 3 {
        if n-i > 3 {
            buf.WriteString(s[i : i+3])
            buf.WriteString(",")
        } else {
            buf.WriteString(s[i:n])
        }
    }

    return buf.String()
}

func main() {
    fmt.Println(comma("1"))
    fmt.Println(comma("12"))
    fmt.Println(comma("123"))
    fmt.Println(comma("123456"))
    fmt.Println(comma("1234567"))
}
