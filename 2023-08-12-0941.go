package main

import "fmt"

func main() {
    // 练习1：打印Hello, World!
    fmt.Println("Hello, World!")

    // 练习2：计算两个整数相加的结果
    a := 10
    b := 5
    fmt.Println("Sum:", a+b)

    // 练习3：判断一个数是否为偶数
    num := 12
    if num%2 == 0 {
        fmt.Println(num, "is even")
    } else {
        fmt.Println(num, "is odd")
    }

    // 练习4：循环打印数字1到10
    for i := 1; i <= 10; i++ {
        fmt.Println(i)
    }

    // 练习5：切片操作
    numbers := []int{1, 2, 3, 4, 5}
    fmt.Println(numbers[1:3])

    // 练习6：Map操作
    m := make(map[string]int)
    m["a"] = 1
    m["b"] = 2
    fmt.Println(m["a"])

    // 练习7：函数调用
    result := add(3, 4)
    fmt.Println("Result:", result)
}

func add(a, b int) int {
    return a + b
}