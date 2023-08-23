package main

import "fmt"

func main() {
    // 1. 打印 "Hello, World!"
    fmt.Println("Hello, World!")

    // 2. 定义一个整数变量 num，赋值为 10，并打印出来
    num := 10
    fmt.Println(num)

    // 3. 定义一个字符串变量 name，赋值为 "Gopher"，并打印出来
    name := "Gopher"
    fmt.Println(name)

    // 4. 定义一个布尔变量 isTrue，赋值为 true，并打印出来
    isTrue := true
    fmt.Println(isTrue)

    // 5. 定义一个浮点数变量 pi，赋值为 3.14，并打印出来
    pi := 3.14
    fmt.Println(pi)

    // 6. 定义一个整数数组 nums，包含元素 1, 2, 3, 4, 5，并打印出来
    nums := []int{1, 2, 3, 4, 5}
    fmt.Println(nums)

    // 7. 定义一个函数 add，接受两个整数参数，返回它们的和
    add := func(a, b int) int {
        return a + b
    }
    fmt.Println(add(3, 5))

    // 8. 对 1-10 的整数进行遍历，并打印出来
    for i := 1; i <= 10; i++ {
        fmt.Println(i)
    }
}