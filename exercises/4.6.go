package main

import (
    "fmt"
    "unicode"
    "unicode/utf8"
)

func trimSpaces(b []byte) []byte {
    var prev rune
    j := 0
    for i := 0; i < len(b); {
        r, size := utf8.DecodeRune(b[i:])
        if i == 0 || !unicode.IsSpace(r) || !unicode.IsSpace(prev) {
            utf8.EncodeRune(b[j:], r)
            j += size
        }
        i += size
        prev = r
    }
    return b[:j]
}

func main() {
    b := []byte("      a    b  ")
    fmt.Printf("%s%s%s\n", "*", string(trimSpaces(b)), "*")
}
