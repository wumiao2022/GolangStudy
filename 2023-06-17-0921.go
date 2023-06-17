package main

import (
    "fmt"
)

func main() {
    // 1. 变量定义与赋值
    var i int = 10
    j := 20
    fmt.Println(i, j)

    // 2. 数组
    var arr [5]int
    for i := 0; i < len(arr); i++ {
        arr[i] = i * i
    }
    fmt.Println(arr)

    // 3. 切片
    slice := arr[1:4]
    fmt.Println(slice)

    // 4. map
    m := make(map[string]int)
    m["apple"] = 1
    m["banana"] = 2
    fmt.Println(m)

    // 5. for循环
    for i := 0; i < 5; i++ {
        fmt.Println(i)
    }

    // 6. if条件判断
    if j > 10 {
        fmt.Println("j is greater than 10")
    } else {
        fmt.Println("j is less than or equal to 10")
    }

    // 7. switch选择语句
    fruit := "apple"
    switch fruit {
    case "apple":
        fmt.Println("It's an apple")
    case "banana":
        fmt.Println("It's a banana")
    default:
        fmt.Println("It's something else")
    }

    // 8. 函数定义和调用
    add := func(a, b int) int {
        return a + b
    }
    sum := add(2, 3)
    fmt.Println(sum)

    // 9. 结构体定义和初始化
    type person struct {
        name string
        age  int
    }
    p := person{name: "Bob", age: 30}
    fmt.Println(p)

    // 10. 接口和类型断言
    type animal interface {
        sound() string
    }
    type dog struct{}
    func (d dog) sound() string {
        return "woof"
    }
    var a animal = dog{}
    fmt.Println(a.(dog).sound())
}