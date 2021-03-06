package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "unicode"
    "unicode/utf8"
)

func main() {
    counts := make(map[rune]int) // count of unicode characters
    categories := make(map[string]int)
    var utflen [utf8.UTFMax]int // count of lenthgs of UTF-8 encodings
    invalid := 0

    in := bufio.NewReader(os.Stdin)
    for {
        r, n, err := in.ReadRune()
        if err == io.EOF {
            break
        }
        if err != nil {
            fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
            os.Exit(1)
        }
        if r == unicode.ReplacementChar && n == 1 {
            invalid++
        }
        counts[r]++
        utflen[n]++
        if unicode.IsSpace(r) {
            categories["space"]++
        }
        if unicode.IsUpper(r) {
            categories["upper"]++
        }
        if unicode.IsLower(r) {
            categories["lower"]++
        }
        if unicode.IsDigit(r) {
            categories["digit"]++
        }
    }
    fmt.Print("rune\tcount\n")
    for c, n := range counts {
        fmt.Printf("%q\t%d\n", c, n)
    }
    fmt.Print("\nlen\tcount\n")
    for i, n := range utflen {
        if n > 0 {
            fmt.Printf("%d\t%d\n", i, n)
        }
    }
    for cat, count := range categories {
        fmt.Printf("%s\t%d\n", cat, count)
    }
    if invalid > 0 {
        fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
    }
}
