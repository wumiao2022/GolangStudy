package main

import "fmt"

func main() {
    // 1. 输出Hello, World!
    fmt.Println("Hello, World!")

    // 2. 计算两个数的和并输出结果
    num1 := 10
    num2 := 5
    sum := num1 + num2
    fmt.Println("Sum:", sum)

    // 3. 判断一个数是否为偶数，并输出结果
    num := 6
    if num%2 == 0 {
        fmt.Println(num, "is even")
    } else {
        fmt.Println(num, "is odd")
    }

    // 4. 判断一个年份是否为闰年，并输出结果
    year := 2020
    if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
        fmt.Println(year, "is a leap year")
    } else {
        fmt.Println(year, "is not a leap year")
    }

    // 5. 求一个整数数组的和，并输出结果
    numbers := []int{2, 4, 6, 8, 10}
    sum = 0
    for _, num := range numbers {
        sum += num
    }
    fmt.Println("Sum:", sum)
}
