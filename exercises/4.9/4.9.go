package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
)

const usage = `wordfreq: counts occurences of words in a file
Usage: wordfreq [FILE_NAME]`

func main() {
    if len(os.Args) < 2 {
        fmt.Fprintln(os.Stderr, usage)
        os.Exit(1)
    }

    file, err := os.Open(os.Args[1])
    if err != nil {
        log.Fatal(err)
    }

    words := make(map[string]int)

    input := bufio.NewScanner(file)
    input.Split(bufio.ScanWords)

    for input.Scan() {
        words[input.Text()]++
    }

    fmt.Printf("word\tcount\n")
    for word, count := range words {
        fmt.Printf("%s\t%d\n", word, count)
    }
}
