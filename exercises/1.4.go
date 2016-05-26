package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    counts := make(map[string]int)
    files := os.Args[1:]
    fnames := make(map[string][]string) // map of line, array of fnames

    if len(files) == 0 {
        countLines(os.Stdin, "", counts, fnames)
    } else {
        for _, fname := range files {
            f, err := os.Open(fname)
            if err != nil {
                fmt.Fprintf(os.Stderr, "error: %v\n", err)
                continue
            }
            countLines(f, fname, counts, fnames)
        }
    }

    for line, n := range counts {
        files := strings.Join(fnames[line], ", ")
        fmt.Printf("line: %s, count: %d, files: %s\n", line, n, files)
    }
}

func countLines(f *os.File, fname string, counts map[string]int, fnames map[string][]string) {
    input := bufio.NewScanner(f)
    for input.Scan() {
        line := input.Text()
        counts[line]++
        if !exists(fname, fnames[line]) {
            fnames[line] = append(fnames[line], fname)
        }
    }
}

func exists(s string, list []string) bool {
    for _, value := range list {
        if value == s {
            return true
        }
    }
    return false
}
