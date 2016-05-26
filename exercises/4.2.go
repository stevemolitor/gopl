package main

import (
    "crypto/sha256"
    "crypto/sha512"
    "flag"
    "fmt"
    "io/ioutil"
    "os"
)

func main() {
    shaSize := flag.Int("shasize", 256, "sha size")
    flag.Parse()

    bytes, err := ioutil.ReadAll(os.Stdin)
    if err != nil {
        fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
    }

    switch *shaSize {
    case 512:
        sha := sha512.Sum512(bytes)
        fmt.Printf("sha512: %x\n", sha)
    case 384:
        sha := sha512.Sum384(bytes)
        fmt.Printf("sha384: %x\n", sha)
    default:
        sha := sha256.Sum256(bytes)
        fmt.Printf("sha256: %x\n", sha)
    }
}
