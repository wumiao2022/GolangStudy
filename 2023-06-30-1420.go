package main

import "fmt"

func main() {
    // 示例1：打印Hello, World!
    fmt.Println("Hello, World!")

    // 示例2：计算两个数的和
    num1 := 3
    num2 := 5
    sum := num1 + num2
    fmt.Println("Sum:", sum)

    // 示例3：使用for循环打印1到10
    for i := 1; i <= 10; i++ {
        fmt.Println(i)
    }

    // 示例4：使用if-else判断一个数是奇数还是偶数
    num := 7
    if num%2 == 0 {
        fmt.Println("Even")
    } else {
        fmt.Println("Odd")
    }

    // 示例5：使用switch-case判断一个学生的成绩属于哪个等级
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

    // 示例6：定义一个结构体类型Person，并创建一个实例
    type Person struct {
        name   string
        age    int
        gender string
    }
    person := Person{name: "Alice", age: 26, gender: "female"}
    fmt.Println(person)
}