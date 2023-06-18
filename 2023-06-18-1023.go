package main

import (
    "fmt"
)

func main() {
    // 1、打印“Hello, world!”
    fmt.Println("Hello, world!")

    // 2、声明三个变量，分别为整型、浮点型、字符串类型，并赋初值后打印输出
    var i int = 10
    var f float64 = 3.14
    var s string = "Golang"
    fmt.Printf("i=%d, f=%f, s=%s\n", i, f, s)

    // 3、编写一个函数，接收两个整型参数，返回它们的和
    func add(a, b int) int {
        return a + b
    }
    sum := add(10, 20)
    fmt.Println("10+20=", sum)

    // 4、编写一个结构体类型，并定义两个方法，分别返回结构体的长和宽
    type rectangle struct {
        length float64
        width  float64
    }
    func (r rectangle) getLength() float64 {
        return r.length
    }
    func (r rectangle) getWidth() float64 {
        return r.width
    }
    r := rectangle{length: 10, width: 5}
    fmt.Printf("Length=%f, width=%f\n", r.getLength(), r.getWidth())
}