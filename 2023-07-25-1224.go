package main

import "fmt"

func main() {
    // 练习：输出Hello World
    fmt.Println("Hello World")

    // 练习：求两个数的和
    sum := add(3, 4)
    fmt.Println("3 + 4 =", sum)

    // 练习：判断一个数是否为偶数
    num := 7
    if isEven(num) {
        fmt.Println(num, "is even")
    } else {
        fmt.Println(num, "is odd")
    }
}

// 函数：求两个数的和
func add(a, b int) int {
    return a + b
}

// 函数：判断一个数是否为偶数
func isEven(num int) bool {
    if num%2 == 0 {
        return true
    }
    return false
}