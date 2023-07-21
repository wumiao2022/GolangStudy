package main

import "fmt"

func main() {
    // 练习1：打印Hello, World!
    fmt.Println("Hello, World!")

    // 练习2：打印1到10的整数
    for i := 1; i <= 10; i++ {
        fmt.Println(i)
    }

    // 练习3：计算两个整数的和并打印结果
    num1 := 10
    num2 := 20
    sum := num1 + num2
    fmt.Println("Sum:", sum)

    // 练习4：判断一个整数是否为偶数，并打印结果
    num := 15
    if num%2 == 0 {
        fmt.Println(num, "is even")
    } else {
        fmt.Println(num, "is odd")
    }

    // 练习5：使用数组存储10个整数，并打印数组中的所有元素
    numbers := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    for _, num := range numbers {
   	 fmt.Println(num)
    }
}