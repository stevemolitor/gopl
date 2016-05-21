package main

import "fmt"

func array(a [3]string) {
    a[0] = "new"
}

func arrayP(a *[3]string) {
    a[0] = "new"
}

func mapFunc(m map[string]string) {
    m["a"] = "new"
    m["b"] = "new2"
}

func slice(a []string) {
    a[0] = "new"
}

func sliceP(a *[]string) {
    (*a)[0] = "new2"
}

func main() {
    a := [3]string{"1", "2", "3"}
    array(a)
    fmt.Println("a after by value:", a)

    arrayP(&a)
    fmt.Println("a after by pointer:", a)

    m := make(map[string]string)
    m["a"] = "original"
    mapFunc(m)
    fmt.Println("m after normal call:", m)

    s := make([]string, 2)
    s[0] = "original"
    slice(s)
    fmt.Printf("slice after normal call: %v\n", s)

    sliceP(&s)
    fmt.Printf("slice after pointer call: %v\n", s)
}
