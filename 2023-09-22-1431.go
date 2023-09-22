package main

import "fmt"

func main() {
    // 1. 使用for循环输出1到10的所有整数

    for i := 1; i <= 10; i++ {
        fmt.Printf("%d ", i)
    }
    fmt.Println()

    // 2. 声明一个字符串变量并给它赋值，然后输出该字符串的长度和第一个字符

    str := "Hello, Golang!"
    fmt.Printf("Length: %d\n", len(str))
    fmt.Printf("First character: %c\n", str[0])

    // 3. 使用switch语句判断一个数字是奇数还是偶数

    num := 5

    switch num % 2 {
    case 0:
        fmt.Println("Even")
    case 1:
        fmt.Println("Odd")
    }

    // 4. 使用数组存储一组整数，然后输出数组的所有元素

    nums := [...]int{1, 2, 3, 4, 5}

    for _, num := range nums {
        fmt.Printf("%d ", num)
    }
    fmt.Println()

    // 5. 声明一个函数，计算两个整数的和，并返回结果

    add := func(a, b int) int {
        return a + b
    }

    result := add(3, 5)
    fmt.Println("Result:", result)
}