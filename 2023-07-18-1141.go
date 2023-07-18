package main

import "fmt"

func main() {
    // 示例1：打印"Hello, World!"
    fmt.Println("Hello, World!")
    
    // 示例2：计算两个整数的和
    a := 5
    b := 3
    sum := a + b
    fmt.Println("The sum of", a, "and", b, "is", sum)
    
    // 示例3：判断一个数是奇数还是偶数
    num := 8
    if num%2 == 0 {
        fmt.Println(num, "is even")
    } else {
        fmt.Println(num, "is odd")
    }
    
    // 示例4：使用for循环打印出1到10的所有整数
    for i := 1; i <= 10; i++ {
        fmt.Println(i)
    }
    
    // 示例5：使用switch语句根据学生的成绩输出对应的等级
    score := 85
    switch {
    case score >= 90:
        fmt.Println("A")
    case score >= 80:
        fmt.Println("B")
    case score >= 70:
        fmt.Println("C")
    case score >= 60:
        fmt.Println("D")
    default:
        fmt.Println("F")
    }
}