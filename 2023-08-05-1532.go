package main

import "fmt"

func main() {
    // 声明变量a和b，并初始化为10和20
    a := 10
    b := 20

    // 打印a和b的值
    fmt.Println("a =", a)
    fmt.Println("b =", b)

    // 交换a和b的值
    a, b = b, a

    // 再次打印a和b的值
    fmt.Println("After swapping:")
    fmt.Println("a =", a)
    fmt.Println("b =", b)
}
