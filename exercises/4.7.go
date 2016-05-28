package main

import "fmt"
import "unicode/utf8"

func reverse(b []byte) []byte {
    var lrune, rrune rune
    var lsize, rsize int

    for lpos, lbeg, lend, rpos, rbeg, rend := 0, 0, 0, len(b), len(b), len(b); lbeg < rend; {
        if lrune == 0 {
            lrune, lsize = utf8.DecodeRune(b[lpos:])
            lend += lsize
        }
        if rrune == 0 {
            rrune, rsize = utf8.DecodeLastRune(b[:rpos])
            rbeg -= rsize
        }

        if lrune != 0 && utf8.ValidRune(lrune) && (rend-rbeg >= lsize) {
            utf8.EncodeRune(b[rend-lsize:], lrune)
            lpos += lsize
            rend -= lsize
            lrune = 0
        }
        if rrune != 0 && (lend-lbeg >= rsize) {
            utf8.EncodeRune(b[lbeg:], rrune)
            rpos -= rsize
            lbeg += rsize
            rrune = 0
        }
    }

    return b
}

func main() {
    // brittle depending on string length - got a little bug still
    b := []byte("hello there \u4316 \u754c")
    fmt.Printf("%s\n", b)
    fmt.Printf("%s\n", reverse(b))
}
