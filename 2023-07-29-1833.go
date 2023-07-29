package main

import "fmt"

func main() {
    // 练习1：打印Hello, World!
    fmt.Println("Hello, World!")

    // 练习2：打印1到10的数字
    for i := 1; i <= 10; i++ {
        fmt.Println(i)
    }

    // 练习3：计算1到100之间的奇数之和
    sum := 0
    for i := 1; i <= 100; i++ {
        if i%2 != 0 {
            sum += i
        }
    }
    fmt.Println("奇数之和：", sum)

    // 练习4：判断一个数是奇数还是偶数
    num := 17
    if num%2 == 0 {
        fmt.Println(num, "是偶数")
    } else {
        fmt.Println(num, "是奇数")
    }

    // 练习5：实现一个简单的计算器函数
    result := calculate(4, 2, "+")
    fmt.Println("计算结果：", result)
}

func calculate(num1, num2 int, operator string) int {
    switch operator {
    case "+":
        return num1 + num2
    case "-":
        return num1 - num2
    case "*":
        return num1 * num2
    case "/":
        if num2 != 0 {
            return num1 / num2
        }
        return 0
    default:
        return 0
    }
}
```