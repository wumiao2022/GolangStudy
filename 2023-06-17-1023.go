package main

import (
    "fmt"
)

func main() {
    // 练习1
    a := []int{1, 2, 3, 4}
    for i := range a {
        fmt.Println(i, a[i])
    }

    // 练习2
    b := "Golang"
    for i, v := range b {
        fmt.Printf("%d %c\n", i, v) 
    }

    // 练习3
    var c uint32 = 0x12345678
    var d byte = byte(c)
    fmt.Printf("0x%x\n", d)

    // 练习4
    e := []int{1, 2, 3, 4}
    f := e[1:3]
    fmt.Println(f)

    // 练习5
    g := make(map[string]int)
    g["a"] = 1
    g["b"] = 2
    g["c"] = 3
    fmt.Println(g)

    // 练习6
    h := make(chan int)
    go func() {
        h <- 1
    }()
    fmt.Println(<-h)

    // 练习7
    i := 10
    defer fmt.Println(i)
    i = 20

    // 练习8
    j := func() int {
        return 1
    }()
    fmt.Println(j)

    // 练习9
    defer func() {
        if err := recover(); err != nil {
            fmt.Println(err)
        }
    }()
    panic("Error")

    // 练习10
    k := []int{1, 2, 3, 4}
    l := append(k[:1], k[2:]...)
    fmt.Println(l)
}