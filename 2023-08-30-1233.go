package main

import "fmt"

func main() {
    // 示例1：打印Hello, World!
    fmt.Println("Hello, World!")

    // 示例2：计算两个数的和
    var num1, num2 int = 5, 10
    sum := num1 + num2
    fmt.Println("Sum:", sum)

    // 示例3：使用for循环打印1到10的数字
    for i := 1; i <= 10; i++ {
        fmt.Println(i)
    }

    // 示例4：判断一个数是奇数还是偶数
    number := 15
    if number % 2 == 0 {
        fmt.Println(number, "is even")
    } else {
        fmt.Println(number, "is odd")
    }

    // 示例5：数组操作，打印数组元素
    numbers := [5]int{1, 2, 3, 4, 5}
    for i := 0; i < len(numbers); i++ {
        fmt.Println(numbers[i])
    }
}
```