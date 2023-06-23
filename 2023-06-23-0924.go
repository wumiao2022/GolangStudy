package main

import "fmt"

func main() {
    // 1. 输出 Hello World
    fmt.Println("Hello World")

    // 2. 输出字符串长度
    str := "Go is awesome!"
    fmt.Println(len(str))

    // 3. 输出整数类型
    var x int = 42
    fmt.Printf("%d\n", x)

    // 4. 输出浮点数类型
    var y float64 = 3.1415926
    fmt.Printf("%f\n", y)

    // 5. 输出布尔类型
    var z bool = true
    fmt.Println(z)

    // 6. 数组遍历
    arr := [5]int{1, 2, 3, 4, 5}
    for i := 0; i < len(arr); i++ {
        fmt.Println(arr[i])
    }

    // 7. 切片遍历
    slice := []int{1, 2, 3, 4, 5}
    for i := 0; i < len(slice); i++ {
        fmt.Println(slice[i])
    }

    // 8. 映射遍历
    m := map[string]int{"a": 1, "b": 2, "c": 3}
    for k, v := range m {
        fmt.Println(k, v)
    }

    // 9. if 判断语句
    num := 10
    if num > 5 {
        fmt.Println("num is greater than 5")
    }

    // 10. switch case 判断语句
    fruit := "apple"
    switch fruit {
    case "banana":
        fmt.Println("it's a banana")
    case "apple":
        fmt.Println("it's an apple")
    default:
        fmt.Println("it's an unknown fruit")
    }

    // 11. for 循环语句
    for i := 0; i < 10; i++ {
        fmt.Println(i)
    }

    // 12. 函数定义和调用
    add := func(x, y int) int {
        return x + y
    }
    fmt.Println(add(1, 2))

    // 13. 接口定义和实现
    type shape interface {
        area() float64
    }
    type square struct {
        side float64
    }
    func (s square) area() float64 {
        return s.side * s.side
    }
    var s shape
    s = square{5}
    fmt.Println(s.area())

    // 14. 结构体定义和使用
    type person struct {
        name string
        age  int
    }
    p := person{"Alice", 20}
    fmt.Println(p.name, p.age)

    // 15. 指针定义和使用
    x1 := 12
    var p1 *int = &x1
    fmt.Println(*p1)
}