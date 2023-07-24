package main

import "fmt"

func main() {
    // 练习1：打印 "Hello, world!"
    fmt.Println("Hello, world!")

    // 练习2：计算两个整数的和并打印结果
    a := 5
    b := 3
    fmt.Println(a + b)

    // 练习3：使用循环打印1到10的数字
    for i := 1; i <= 10; i++ {
        fmt.Println(i)
    }

    // 练习4：定义一个数组，包含5个学生的成绩，然后计算平均分
    scores := [5]int{78, 85, 92, 88, 95}
    sum := 0
    for _, score := range scores {
        sum += score
    }
    average := float64(sum) / float64(len(scores))
    fmt.Println("Average score:", average)

    // 练习5：定义一个函数，接收一个字符串作为参数并将其反转
    str := "Hello, world!"
    reversedStr := reverseString(str)
    fmt.Println("Reversed string:", reversedStr)
}

func reverseString(str string) string {
    reversed := ""
    for i := len(str) - 1; i >= 0; i-- {
        reversed += string(str[i])
    }
    return reversed
}