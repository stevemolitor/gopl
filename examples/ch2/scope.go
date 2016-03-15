package main

import (
    "fmt"
)

func testScope(a int) {
    // each else and else if creates a new lexical scope, nested in previous one
    if x := 1; a == 1 {
        fmt.Printf("if, x: %d\n", x)
        // fmt.Printf("if, x: %d, y: %d\n", x, y) -- undefined 'y'
    } else if y := 2; a == 2 {
        fmt.Printf("elif 1, x: %d, y: %d\n", x, y)
    } else if z := 3; a == 3 {
        fmt.Printf("elif 2, x: %d, y: %d, z: %d\n", x, y, z)
    } else {
        fmt.Printf("else, x: %d, y: %d, z: %d\n", x, y, z)
    }
}

func main() {
    testScope(1)
    testScope(2)
    testScope(3)
    testScope(4)
}
